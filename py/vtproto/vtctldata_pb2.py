# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: vtctldata.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


import logutil_pb2 as logutil__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='vtctldata.proto',
  package='vtctldata',
  syntax='proto3',
  serialized_pb=_b('\n\x0fvtctldata.proto\x12\tvtctldata\x1a\rlogutil.proto\"B\n\x1a\x45xecuteVtctlCommandRequest\x12\x0c\n\x04\x61rgs\x18\x01 \x03(\t\x12\x16\n\x0e\x61\x63tion_timeout\x18\x02 \x01(\x03\"<\n\x1b\x45xecuteVtctlCommandResponse\x12\x1d\n\x05\x65vent\x18\x01 \x01(\x0b\x32\x0e.logutil.EventB(Z&github.com/xsec-lab/vitess/go/vt/proto/vtctldatab\x06proto3')
  ,
  dependencies=[logutil__pb2.DESCRIPTOR,])




_EXECUTEVTCTLCOMMANDREQUEST = _descriptor.Descriptor(
  name='ExecuteVtctlCommandRequest',
  full_name='vtctldata.ExecuteVtctlCommandRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='args', full_name='vtctldata.ExecuteVtctlCommandRequest.args', index=0,
      number=1, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='action_timeout', full_name='vtctldata.ExecuteVtctlCommandRequest.action_timeout', index=1,
      number=2, type=3, cpp_type=2, label=1,
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
  serialized_start=45,
  serialized_end=111,
)


_EXECUTEVTCTLCOMMANDRESPONSE = _descriptor.Descriptor(
  name='ExecuteVtctlCommandResponse',
  full_name='vtctldata.ExecuteVtctlCommandResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='event', full_name='vtctldata.ExecuteVtctlCommandResponse.event', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
  serialized_start=113,
  serialized_end=173,
)

_EXECUTEVTCTLCOMMANDRESPONSE.fields_by_name['event'].message_type = logutil__pb2._EVENT
DESCRIPTOR.message_types_by_name['ExecuteVtctlCommandRequest'] = _EXECUTEVTCTLCOMMANDREQUEST
DESCRIPTOR.message_types_by_name['ExecuteVtctlCommandResponse'] = _EXECUTEVTCTLCOMMANDRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ExecuteVtctlCommandRequest = _reflection.GeneratedProtocolMessageType('ExecuteVtctlCommandRequest', (_message.Message,), dict(
  DESCRIPTOR = _EXECUTEVTCTLCOMMANDREQUEST,
  __module__ = 'vtctldata_pb2'
  # @@protoc_insertion_point(class_scope:vtctldata.ExecuteVtctlCommandRequest)
  ))
_sym_db.RegisterMessage(ExecuteVtctlCommandRequest)

ExecuteVtctlCommandResponse = _reflection.GeneratedProtocolMessageType('ExecuteVtctlCommandResponse', (_message.Message,), dict(
  DESCRIPTOR = _EXECUTEVTCTLCOMMANDRESPONSE,
  __module__ = 'vtctldata_pb2'
  # @@protoc_insertion_point(class_scope:vtctldata.ExecuteVtctlCommandResponse)
  ))
_sym_db.RegisterMessage(ExecuteVtctlCommandResponse)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('Z&github.com/xsec-lab/vitess/go/vt/proto/vtctldata'))
# @@protoc_insertion_point(module_scope)
