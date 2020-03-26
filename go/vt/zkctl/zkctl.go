/*
Copyright 2017 Google Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

/*
Commands for controlling an external zookeeper process.
*/

package zkctl

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"os/exec"
	"path"
	"strconv"
	"syscall"
	"time"

	zookeeper "github.com/samuel/go-zookeeper/zk"

	"github.com/xsec-lab/go/vt/env"
	"github.com/xsec-lab/go/vt/log"
)

const (
	// startWaitTime is the number of seconds to wait at Start.
	startWaitTime = 20
	// shutdownWaitTime is the number of seconds to wait at Shutdown.
	shutdownWaitTime = 20
)

// Zkd manages the running of ZooKeeper servers.
type Zkd struct {
	config *ZkConfig
	done   chan struct{}
}

// NewZkd creates a Zkd.
func NewZkd(config *ZkConfig) *Zkd {
	return &Zkd{config: config}
}

// Done returns a channel that is closed when the underlying process started
// by this Zkd has terminated. If the process was started by someone else, this
// channel will never be closed.
func (zkd *Zkd) Done() <-chan struct{} {
	return zkd.done
}

/*
 ZOO_LOG_DIR=""
 ZOO_CFG="/.../zoo.cfg"
 ZOOMAIN="org.apache.zookeeper.server.quorum.QuorumPeerMain"
 java -DZOO_LOG_DIR=${ZOO_LOG_DIR} -cp $CLASSPATH $ZOOMAIN $YT_ZK_CFG
*/

// Start runs an already initialized ZooKeeper server.
func (zkd *Zkd) Start() error {
	log.Infof("zkctl.Start")
	// NOTE(msolomon) use a script here so we can detach and continue to run
	// if the wrangler process dies. this pretty much the same as mysqld_safe.
	args := []string{
		zkd.config.LogDir(),
		zkd.config.ConfigFile(),
		zkd.config.PidFile(),
	}
	root, err := env.VtRoot()
	if err != nil {
		return err
	}
	dir := path.Join(root, "bin")
	cmd := exec.Command(path.Join(dir, "zksrv.sh"), args...)
	cmd.Env = os.Environ()
	cmd.Dir = dir

	if err = cmd.Start(); err != nil {
		return err
	}

	// give it some time to succeed - usually by the time the socket emerges
	// we are in good shape
	for i := 0; i < startWaitTime; i++ {
		zkAddr := fmt.Sprintf(":%v", zkd.config.ClientPort)
		conn, connErr := net.Dial("tcp", zkAddr)
		if connErr != nil {
			err = connErr
			time.Sleep(time.Second)
			continue
		} else {
			err = nil
			conn.Write([]byte("ruok"))
			reply := make([]byte, 4)
			conn.Read(reply)
			if string(reply) != "imok" {
				err = fmt.Errorf("local zk unhealthy: %v %v", zkAddr, reply)
			}
			conn.Close()
			break
		}
	}
	zkd.done = make(chan struct{})
	go func(done chan<- struct{}) {
		// wait so we don't get a bunch of defunct processes
		cmd.Wait()
		close(done)
	}(zkd.done)
	return err
}

// Shutdown kills a ZooKeeper server, but keeps its data dir intact.
func (zkd *Zkd) Shutdown() error {
	log.Infof("zkctl.Shutdown")
	pidData, err := ioutil.ReadFile(zkd.config.PidFile())
	if err != nil {
		return err
	}
	pid, err := strconv.Atoi(string(bytes.TrimSpace(pidData)))
	if err != nil {
		return err
	}
	err = syscall.Kill(pid, syscall.SIGKILL)
	if err != nil && err != syscall.ESRCH {
		return err
	}
	for i := 0; i < shutdownWaitTime; i++ {
		if syscall.Kill(pid, syscall.SIGKILL) == syscall.ESRCH {
			return nil
		}
		time.Sleep(time.Second)
	}
	return fmt.Errorf("Shutdown didn't kill process %v", pid)
}

func (zkd *Zkd) makeCfg() (string, error) {
	root, err := env.VtRoot()
	if err != nil {
		return "", err
	}
	cnfTemplatePaths := []string{path.Join(root, "config/zkcfg/zoo.cfg")}
	return MakeZooCfg(cnfTemplatePaths, zkd.config, "# generated by vt")
}

// Init generates a new config and then starts ZooKeeper.
func (zkd *Zkd) Init() error {
	if zkd.Inited() {
		return fmt.Errorf("zk already inited")
	}

	log.Infof("zkd.Init")
	for _, path := range zkd.config.DirectoryList() {
		if err := os.MkdirAll(path, 0775); err != nil {
			log.Errorf("%v", err)
			return err
		}
		// FIXME(msolomon) validate permissions?
	}

	configData, err := zkd.makeCfg()
	if err == nil {
		err = ioutil.WriteFile(zkd.config.ConfigFile(), []byte(configData), 0664)
	}
	if err != nil {
		log.Errorf("failed creating %v: %v", zkd.config.ConfigFile(), err)
		return err
	}

	err = zkd.config.WriteMyid()
	if err != nil {
		log.Errorf("failed creating %v: %v", zkd.config.MyidFile(), err)
		return err
	}

	if err = zkd.Start(); err != nil {
		log.Errorf("failed starting, check %v", zkd.config.LogDir())
		return err
	}

	zkAddr := fmt.Sprintf("localhost:%v", zkd.config.ClientPort)
	zk, session, err := zookeeper.Connect([]string{zkAddr}, startWaitTime*time.Second)
	if err != nil {
		return err
	}
	event := <-session
	if event.State != zookeeper.StateConnecting {
		return event.Err
	}
	event = <-session
	if event.State != zookeeper.StateConnected {
		return event.Err
	}
	defer zk.Close()

	return nil
}

// Teardown shuts down the server and removes its data dir.
func (zkd *Zkd) Teardown() error {
	log.Infof("zkctl.Teardown")
	if err := zkd.Shutdown(); err != nil {
		log.Warningf("failed zookeeper shutdown: %v", err.Error())
	}
	var removalErr error
	for _, dir := range zkd.config.DirectoryList() {
		log.V(6).Infof("remove data dir %v", dir)
		if err := os.RemoveAll(dir); err != nil {
			log.Errorf("failed removing %v: %v", dir, err.Error())
			removalErr = err
		}
	}
	return removalErr
}

// Inited returns true if the server config has been initialized.
func (zkd *Zkd) Inited() bool {
	myidFile := zkd.config.MyidFile()
	_, statErr := os.Stat(myidFile)
	if statErr == nil {
		return true
	} else if statErr.(*os.PathError).Err != syscall.ENOENT {
		panic("can't access file " + myidFile + ": " + statErr.Error())
	}
	return false
}
