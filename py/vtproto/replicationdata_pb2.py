# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: replicationdata.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='replicationdata.proto',
  package='replicationdata',
  syntax='proto3',
  serialized_pb=_b('\n\x15replicationdata.proto\x12\x0freplicationdata\"\xb6\x01\n\x06Status\x12\x10\n\x08position\x18\x01 \x01(\t\x12\x18\n\x10slave_io_running\x18\x02 \x01(\x08\x12\x19\n\x11slave_sql_running\x18\x03 \x01(\x08\x12\x1d\n\x15seconds_behind_master\x18\x04 \x01(\r\x12\x13\n\x0bmaster_host\x18\x05 \x01(\t\x12\x13\n\x0bmaster_port\x18\x06 \x01(\x05\x12\x1c\n\x14master_connect_retry\x18\x07 \x01(\x05\x42.Z,github.com/xsec-lab/go/vt/proto/replicationdatab\x06proto3')
)




_STATUS = _descriptor.Descriptor(
  name='Status',
  full_name='replicationdata.Status',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='position', full_name='replicationdata.Status.position', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='slave_io_running', full_name='replicationdata.Status.slave_io_running', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='slave_sql_running', full_name='replicationdata.Status.slave_sql_running', index=2,
      number=3, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='seconds_behind_master', full_name='replicationdata.Status.seconds_behind_master', index=3,
      number=4, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='master_host', full_name='replicationdata.Status.master_host', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='master_port', full_name='replicationdata.Status.master_port', index=5,
      number=6, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='master_connect_retry', full_name='replicationdata.Status.master_connect_retry', index=6,
      number=7, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=43,
  serialized_end=225,
)

DESCRIPTOR.message_types_by_name['Status'] = _STATUS
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Status = _reflection.GeneratedProtocolMessageType('Status', (_message.Message,), dict(
  DESCRIPTOR = _STATUS,
  __module__ = 'replicationdata_pb2'
  # @@protoc_insertion_point(class_scope:replicationdata.Status)
  ))
_sym_db.RegisterMessage(Status)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('Z,github.com/xsec-lab/go/vt/proto/replicationdata'))
# @@protoc_insertion_point(module_scope)
