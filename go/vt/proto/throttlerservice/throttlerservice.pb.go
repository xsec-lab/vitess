// Code generated by protoc-gen-go. DO NOT EDIT.
// source: throttlerservice.proto

package throttlerservice // import "github.com/xsec-lab/go/vt/proto/throttlerservice"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import throttlerdata "github.com/xsec-lab/go/vt/proto/throttlerdata"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ThrottlerClient is the client API for Throttler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ThrottlerClient interface {
	// MaxRates returns the current max rate for each throttler of the process.
	MaxRates(ctx context.Context, in *throttlerdata.MaxRatesRequest, opts ...grpc.CallOption) (*throttlerdata.MaxRatesResponse, error)
	// SetMaxRate allows to change the current max rate for all throttlers
	// of the process.
	SetMaxRate(ctx context.Context, in *throttlerdata.SetMaxRateRequest, opts ...grpc.CallOption) (*throttlerdata.SetMaxRateResponse, error)
	// GetConfiguration returns the configuration of the MaxReplicationlag module
	// for the given throttler or all throttlers if "throttler_name" is empty.
	GetConfiguration(ctx context.Context, in *throttlerdata.GetConfigurationRequest, opts ...grpc.CallOption) (*throttlerdata.GetConfigurationResponse, error)
	// UpdateConfiguration (partially) updates the configuration of the
	// MaxReplicationlag module for the given throttler or all throttlers if
	// "throttler_name" is empty.
	// If "copy_zero_values" is true, fields with zero values will be copied
	// as well.
	UpdateConfiguration(ctx context.Context, in *throttlerdata.UpdateConfigurationRequest, opts ...grpc.CallOption) (*throttlerdata.UpdateConfigurationResponse, error)
	// ResetConfiguration resets the configuration of the MaxReplicationlag module
	// to the initial configuration for the given throttler or all throttlers if
	// "throttler_name" is empty.
	ResetConfiguration(ctx context.Context, in *throttlerdata.ResetConfigurationRequest, opts ...grpc.CallOption) (*throttlerdata.ResetConfigurationResponse, error)
}

type throttlerClient struct {
	cc *grpc.ClientConn
}

func NewThrottlerClient(cc *grpc.ClientConn) ThrottlerClient {
	return &throttlerClient{cc}
}

func (c *throttlerClient) MaxRates(ctx context.Context, in *throttlerdata.MaxRatesRequest, opts ...grpc.CallOption) (*throttlerdata.MaxRatesResponse, error) {
	out := new(throttlerdata.MaxRatesResponse)
	err := c.cc.Invoke(ctx, "/throttlerservice.Throttler/MaxRates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *throttlerClient) SetMaxRate(ctx context.Context, in *throttlerdata.SetMaxRateRequest, opts ...grpc.CallOption) (*throttlerdata.SetMaxRateResponse, error) {
	out := new(throttlerdata.SetMaxRateResponse)
	err := c.cc.Invoke(ctx, "/throttlerservice.Throttler/SetMaxRate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *throttlerClient) GetConfiguration(ctx context.Context, in *throttlerdata.GetConfigurationRequest, opts ...grpc.CallOption) (*throttlerdata.GetConfigurationResponse, error) {
	out := new(throttlerdata.GetConfigurationResponse)
	err := c.cc.Invoke(ctx, "/throttlerservice.Throttler/GetConfiguration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *throttlerClient) UpdateConfiguration(ctx context.Context, in *throttlerdata.UpdateConfigurationRequest, opts ...grpc.CallOption) (*throttlerdata.UpdateConfigurationResponse, error) {
	out := new(throttlerdata.UpdateConfigurationResponse)
	err := c.cc.Invoke(ctx, "/throttlerservice.Throttler/UpdateConfiguration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *throttlerClient) ResetConfiguration(ctx context.Context, in *throttlerdata.ResetConfigurationRequest, opts ...grpc.CallOption) (*throttlerdata.ResetConfigurationResponse, error) {
	out := new(throttlerdata.ResetConfigurationResponse)
	err := c.cc.Invoke(ctx, "/throttlerservice.Throttler/ResetConfiguration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ThrottlerServer is the server API for Throttler service.
type ThrottlerServer interface {
	// MaxRates returns the current max rate for each throttler of the process.
	MaxRates(context.Context, *throttlerdata.MaxRatesRequest) (*throttlerdata.MaxRatesResponse, error)
	// SetMaxRate allows to change the current max rate for all throttlers
	// of the process.
	SetMaxRate(context.Context, *throttlerdata.SetMaxRateRequest) (*throttlerdata.SetMaxRateResponse, error)
	// GetConfiguration returns the configuration of the MaxReplicationlag module
	// for the given throttler or all throttlers if "throttler_name" is empty.
	GetConfiguration(context.Context, *throttlerdata.GetConfigurationRequest) (*throttlerdata.GetConfigurationResponse, error)
	// UpdateConfiguration (partially) updates the configuration of the
	// MaxReplicationlag module for the given throttler or all throttlers if
	// "throttler_name" is empty.
	// If "copy_zero_values" is true, fields with zero values will be copied
	// as well.
	UpdateConfiguration(context.Context, *throttlerdata.UpdateConfigurationRequest) (*throttlerdata.UpdateConfigurationResponse, error)
	// ResetConfiguration resets the configuration of the MaxReplicationlag module
	// to the initial configuration for the given throttler or all throttlers if
	// "throttler_name" is empty.
	ResetConfiguration(context.Context, *throttlerdata.ResetConfigurationRequest) (*throttlerdata.ResetConfigurationResponse, error)
}

func RegisterThrottlerServer(s *grpc.Server, srv ThrottlerServer) {
	s.RegisterService(&_Throttler_serviceDesc, srv)
}

func _Throttler_MaxRates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(throttlerdata.MaxRatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThrottlerServer).MaxRates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/throttlerservice.Throttler/MaxRates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThrottlerServer).MaxRates(ctx, req.(*throttlerdata.MaxRatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Throttler_SetMaxRate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(throttlerdata.SetMaxRateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThrottlerServer).SetMaxRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/throttlerservice.Throttler/SetMaxRate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThrottlerServer).SetMaxRate(ctx, req.(*throttlerdata.SetMaxRateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Throttler_GetConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(throttlerdata.GetConfigurationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThrottlerServer).GetConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/throttlerservice.Throttler/GetConfiguration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThrottlerServer).GetConfiguration(ctx, req.(*throttlerdata.GetConfigurationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Throttler_UpdateConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(throttlerdata.UpdateConfigurationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThrottlerServer).UpdateConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/throttlerservice.Throttler/UpdateConfiguration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThrottlerServer).UpdateConfiguration(ctx, req.(*throttlerdata.UpdateConfigurationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Throttler_ResetConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(throttlerdata.ResetConfigurationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThrottlerServer).ResetConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/throttlerservice.Throttler/ResetConfiguration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThrottlerServer).ResetConfiguration(ctx, req.(*throttlerdata.ResetConfigurationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Throttler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "throttlerservice.Throttler",
	HandlerType: (*ThrottlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MaxRates",
			Handler:    _Throttler_MaxRates_Handler,
		},
		{
			MethodName: "SetMaxRate",
			Handler:    _Throttler_SetMaxRate_Handler,
		},
		{
			MethodName: "GetConfiguration",
			Handler:    _Throttler_GetConfiguration_Handler,
		},
		{
			MethodName: "UpdateConfiguration",
			Handler:    _Throttler_UpdateConfiguration_Handler,
		},
		{
			MethodName: "ResetConfiguration",
			Handler:    _Throttler_ResetConfiguration_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "throttlerservice.proto",
}

func init() {
	proto.RegisterFile("throttlerservice.proto", fileDescriptor_throttlerservice_151ce3faa7ac0b15)
}

var fileDescriptor_throttlerservice_151ce3faa7ac0b15 = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x3d, 0x4b, 0xc4, 0x40,
	0x10, 0x86, 0x05, 0x41, 0x74, 0xaa, 0x63, 0x0f, 0x2c, 0xae, 0xf0, 0xab, 0x50, 0x4f, 0x30, 0x0b,
	0xfa, 0x0f, 0xb4, 0xb0, 0xba, 0x26, 0xa7, 0x8d, 0xdd, 0xea, 0x8d, 0x71, 0x51, 0x76, 0xe2, 0xce,
	0x24, 0xf8, 0xbf, 0xfd, 0x03, 0x42, 0xe2, 0xae, 0x64, 0xfc, 0xb8, 0x74, 0xe1, 0x7d, 0x9f, 0x7d,
	0x1f, 0x02, 0x03, 0xbb, 0xf2, 0x1c, 0x49, 0xe4, 0x15, 0x23, 0x63, 0x6c, 0xfd, 0x23, 0x16, 0x75,
	0x24, 0x21, 0x33, 0xd1, 0xf9, 0x6c, 0x9a, 0x93, 0x95, 0x13, 0xd7, 0x63, 0x17, 0x1f, 0x9b, 0xb0,
	0x73, 0x9b, 0x72, 0xb3, 0x80, 0xed, 0x85, 0x7b, 0x2f, 0x9d, 0x20, 0x9b, 0xbd, 0x62, 0xc8, 0xa7,
	0xa2, 0xc4, 0xb7, 0x06, 0x59, 0x66, 0xfb, 0x7f, 0xf6, 0x5c, 0x53, 0x60, 0x3c, 0xda, 0x30, 0x4b,
	0x80, 0x25, 0xca, 0x57, 0x61, 0x0e, 0xd4, 0x83, 0xef, 0x2a, 0x4d, 0x1e, 0xfe, 0x43, 0xe4, 0x51,
	0x84, 0xc9, 0x0d, 0xca, 0x35, 0x85, 0x27, 0x5f, 0x35, 0xd1, 0x89, 0xa7, 0x60, 0x8e, 0xd5, 0x43,
	0x0d, 0x24, 0xc1, 0xc9, 0x5a, 0x2e, 0x6b, 0x02, 0x4c, 0xef, 0xea, 0x95, 0x13, 0x1c, 0x9a, 0xe6,
	0x6a, 0xe1, 0x17, 0x26, 0xc9, 0xce, 0xc6, 0xa0, 0xd9, 0xf7, 0x02, 0xa6, 0x44, 0xd6, 0x3f, 0x76,
	0xaa, 0x36, 0x7e, 0x22, 0xc9, 0x36, 0x1f, 0x41, 0x26, 0xd9, 0x95, 0xbd, 0x3f, 0x6f, 0xbd, 0x20,
	0x73, 0xe1, 0xc9, 0xf6, 0x5f, 0xb6, 0x22, 0xdb, 0x8a, 0xed, 0xae, 0xc2, 0xea, 0xdb, 0x79, 0xd8,
	0xea, 0xf2, 0xcb, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x49, 0x64, 0xc0, 0xd9, 0x6e, 0x02, 0x00,
	0x00,
}
