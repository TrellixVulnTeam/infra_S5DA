# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: annotations.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='annotations.proto',
  package='milo',
  syntax='proto3',
  serialized_pb=_b('\n\x11\x61nnotations.proto\x12\x04milo\x1a\x1fgoogle/protobuf/timestamp.proto\"\xd8\x01\n\x0e\x46\x61ilureDetails\x12\'\n\x04type\x18\x01 \x01(\x0e\x32\x19.milo.FailureDetails.Type\x12\x0c\n\x04text\x18\x02 \x01(\t\x12*\n\x14\x66\x61iled_dm_dependency\x18\x03 \x03(\x0b\x32\x0c.milo.DMLink\"c\n\x04Type\x12\x0b\n\x07GENERAL\x10\x00\x12\r\n\tEXCEPTION\x10\x01\x12\t\n\x05INFRA\x10\x02\x12\x18\n\x14\x44M_DEPENDENCY_FAILED\x10\x03\x12\r\n\tCANCELLED\x10\x04\x12\x0b\n\x07\x45XPIRED\x10\x05\"\xbb\x06\n\x04Step\x12\x0c\n\x04name\x18\x01 \x01(\t\x12#\n\x07\x63ommand\x18\x02 \x01(\x0b\x32\x12.milo.Step.Command\x12\x1c\n\x06status\x18\x03 \x01(\x0e\x32\x0c.milo.Status\x12-\n\x0f\x66\x61ilure_details\x18\x04 \x01(\x0b\x32\x14.milo.FailureDetails\x12#\n\x07substep\x18\x05 \x03(\x0b\x32\x12.milo.Step.Substep\x12)\n\rstdout_stream\x18\x06 \x01(\x0b\x32\x12.milo.LogdogStream\x12)\n\rstderr_stream\x18\x07 \x01(\x0b\x32\x12.milo.LogdogStream\x12+\n\x07started\x18\x08 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12)\n\x05\x65nded\x18\t \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x0c\n\x04text\x18\x14 \x03(\t\x12%\n\x08progress\x18\x15 \x01(\x0b\x32\x13.milo.Step.Progress\x12\x18\n\x04link\x18\x16 \x01(\x0b\x32\n.milo.Link\x12\x1f\n\x0bother_links\x18\x17 \x03(\x0b\x32\n.milo.Link\x12%\n\x08property\x18\x18 \x03(\x0b\x32\x13.milo.Step.Property\x1a\x8e\x01\n\x07\x43ommand\x12\x14\n\x0c\x63ommand_line\x18\x01 \x03(\t\x12\x0b\n\x03\x63wd\x18\x02 \x01(\t\x12\x30\n\x07\x65nviron\x18\x03 \x03(\x0b\x32\x1f.milo.Step.Command.EnvironEntry\x1a.\n\x0c\x45nvironEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\x1a\x61\n\x07Substep\x12\x1a\n\x04step\x18\x01 \x01(\x0b\x32\n.milo.StepH\x00\x12/\n\x11\x61nnotation_stream\x18\x02 \x01(\x0b\x32\x12.milo.LogdogStreamH\x00\x42\t\n\x07substep\x1a,\n\x08Progress\x12\r\n\x05total\x18\x01 \x01(\x05\x12\x11\n\tcompleted\x18\x02 \x01(\x05\x1a\'\n\x08Property\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t\"\xbf\x01\n\x04Link\x12\r\n\x05label\x18\x01 \x01(\t\x12\x13\n\x0b\x61lias_label\x18\x02 \x01(\t\x12\r\n\x03url\x18\x03 \x01(\tH\x00\x12+\n\rlogdog_stream\x18\x04 \x01(\x0b\x32\x12.milo.LogdogStreamH\x00\x12-\n\x0eisolate_object\x18\x05 \x01(\x0b\x32\x13.milo.IsolateObjectH\x00\x12\x1f\n\x07\x64m_link\x18\x06 \x01(\x0b\x32\x0c.milo.DMLinkH\x00\x42\x07\n\x05value\"<\n\x0cLogdogStream\x12\x0e\n\x06server\x18\x01 \x01(\t\x12\x0e\n\x06prefix\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\"-\n\rIsolateObject\x12\x0e\n\x06server\x18\x01 \x01(\t\x12\x0c\n\x04hash\x18\x02 \x01(\t\"K\n\x06\x44MLink\x12\x0e\n\x06server\x18\x01 \x01(\t\x12\r\n\x05quest\x18\x02 \x01(\t\x12\x0f\n\x07\x61ttempt\x18\x03 \x01(\x03\x12\x11\n\texecution\x18\x04 \x01(\x03*<\n\x06Status\x12\x0b\n\x07RUNNING\x10\x00\x12\x0b\n\x07SUCCESS\x10\x01\x12\x0b\n\x07\x46\x41ILURE\x10\x02\x12\x0b\n\x07PENDING\x10\x03\x62\x06proto3')
  ,
  dependencies=[google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,])
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

_STATUS = _descriptor.EnumDescriptor(
  name='Status',
  full_name='milo.Status',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='RUNNING', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUCCESS', index=1, number=1,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='FAILURE', index=2, number=2,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PENDING', index=3, number=3,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=1489,
  serialized_end=1549,
)
_sym_db.RegisterEnumDescriptor(_STATUS)

Status = enum_type_wrapper.EnumTypeWrapper(_STATUS)
RUNNING = 0
SUCCESS = 1
FAILURE = 2
PENDING = 3


_FAILUREDETAILS_TYPE = _descriptor.EnumDescriptor(
  name='Type',
  full_name='milo.FailureDetails.Type',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='GENERAL', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='EXCEPTION', index=1, number=1,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='INFRA', index=2, number=2,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='DM_DEPENDENCY_FAILED', index=3, number=3,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CANCELLED', index=4, number=4,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='EXPIRED', index=5, number=5,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=178,
  serialized_end=277,
)
_sym_db.RegisterEnumDescriptor(_FAILUREDETAILS_TYPE)


_FAILUREDETAILS = _descriptor.Descriptor(
  name='FailureDetails',
  full_name='milo.FailureDetails',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='type', full_name='milo.FailureDetails.type', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='text', full_name='milo.FailureDetails.text', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='failed_dm_dependency', full_name='milo.FailureDetails.failed_dm_dependency', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _FAILUREDETAILS_TYPE,
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=61,
  serialized_end=277,
)


_STEP_COMMAND_ENVIRONENTRY = _descriptor.Descriptor(
  name='EnvironEntry',
  full_name='milo.Step.Command.EnvironEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='milo.Step.Command.EnvironEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='value', full_name='milo.Step.Command.EnvironEntry.value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=_descriptor._ParseOptions(descriptor_pb2.MessageOptions(), _b('8\001')),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=875,
  serialized_end=921,
)

_STEP_COMMAND = _descriptor.Descriptor(
  name='Command',
  full_name='milo.Step.Command',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='command_line', full_name='milo.Step.Command.command_line', index=0,
      number=1, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='cwd', full_name='milo.Step.Command.cwd', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='environ', full_name='milo.Step.Command.environ', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[_STEP_COMMAND_ENVIRONENTRY, ],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=779,
  serialized_end=921,
)

_STEP_SUBSTEP = _descriptor.Descriptor(
  name='Substep',
  full_name='milo.Step.Substep',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='step', full_name='milo.Step.Substep.step', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='annotation_stream', full_name='milo.Step.Substep.annotation_stream', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
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
    _descriptor.OneofDescriptor(
      name='substep', full_name='milo.Step.Substep.substep',
      index=0, containing_type=None, fields=[]),
  ],
  serialized_start=923,
  serialized_end=1020,
)

_STEP_PROGRESS = _descriptor.Descriptor(
  name='Progress',
  full_name='milo.Step.Progress',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='total', full_name='milo.Step.Progress.total', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='completed', full_name='milo.Step.Progress.completed', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
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
  serialized_start=1022,
  serialized_end=1066,
)

_STEP_PROPERTY = _descriptor.Descriptor(
  name='Property',
  full_name='milo.Step.Property',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='milo.Step.Property.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='value', full_name='milo.Step.Property.value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
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
  serialized_start=1068,
  serialized_end=1107,
)

_STEP = _descriptor.Descriptor(
  name='Step',
  full_name='milo.Step',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='milo.Step.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='command', full_name='milo.Step.command', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='status', full_name='milo.Step.status', index=2,
      number=3, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='failure_details', full_name='milo.Step.failure_details', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='substep', full_name='milo.Step.substep', index=4,
      number=5, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='stdout_stream', full_name='milo.Step.stdout_stream', index=5,
      number=6, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='stderr_stream', full_name='milo.Step.stderr_stream', index=6,
      number=7, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='started', full_name='milo.Step.started', index=7,
      number=8, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='ended', full_name='milo.Step.ended', index=8,
      number=9, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='text', full_name='milo.Step.text', index=9,
      number=20, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='progress', full_name='milo.Step.progress', index=10,
      number=21, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='link', full_name='milo.Step.link', index=11,
      number=22, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='other_links', full_name='milo.Step.other_links', index=12,
      number=23, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='property', full_name='milo.Step.property', index=13,
      number=24, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[_STEP_COMMAND, _STEP_SUBSTEP, _STEP_PROGRESS, _STEP_PROPERTY, ],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=280,
  serialized_end=1107,
)


_LINK = _descriptor.Descriptor(
  name='Link',
  full_name='milo.Link',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='label', full_name='milo.Link.label', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='alias_label', full_name='milo.Link.alias_label', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='url', full_name='milo.Link.url', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='logdog_stream', full_name='milo.Link.logdog_stream', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='isolate_object', full_name='milo.Link.isolate_object', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='dm_link', full_name='milo.Link.dm_link', index=5,
      number=6, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
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
    _descriptor.OneofDescriptor(
      name='value', full_name='milo.Link.value',
      index=0, containing_type=None, fields=[]),
  ],
  serialized_start=1110,
  serialized_end=1301,
)


_LOGDOGSTREAM = _descriptor.Descriptor(
  name='LogdogStream',
  full_name='milo.LogdogStream',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='server', full_name='milo.LogdogStream.server', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='prefix', full_name='milo.LogdogStream.prefix', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='name', full_name='milo.LogdogStream.name', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
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
  serialized_start=1303,
  serialized_end=1363,
)


_ISOLATEOBJECT = _descriptor.Descriptor(
  name='IsolateObject',
  full_name='milo.IsolateObject',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='server', full_name='milo.IsolateObject.server', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='hash', full_name='milo.IsolateObject.hash', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
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
  serialized_start=1365,
  serialized_end=1410,
)


_DMLINK = _descriptor.Descriptor(
  name='DMLink',
  full_name='milo.DMLink',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='server', full_name='milo.DMLink.server', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='quest', full_name='milo.DMLink.quest', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='attempt', full_name='milo.DMLink.attempt', index=2,
      number=3, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='execution', full_name='milo.DMLink.execution', index=3,
      number=4, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
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
  serialized_start=1412,
  serialized_end=1487,
)

_FAILUREDETAILS.fields_by_name['type'].enum_type = _FAILUREDETAILS_TYPE
_FAILUREDETAILS.fields_by_name['failed_dm_dependency'].message_type = _DMLINK
_FAILUREDETAILS_TYPE.containing_type = _FAILUREDETAILS
_STEP_COMMAND_ENVIRONENTRY.containing_type = _STEP_COMMAND
_STEP_COMMAND.fields_by_name['environ'].message_type = _STEP_COMMAND_ENVIRONENTRY
_STEP_COMMAND.containing_type = _STEP
_STEP_SUBSTEP.fields_by_name['step'].message_type = _STEP
_STEP_SUBSTEP.fields_by_name['annotation_stream'].message_type = _LOGDOGSTREAM
_STEP_SUBSTEP.containing_type = _STEP
_STEP_SUBSTEP.oneofs_by_name['substep'].fields.append(
  _STEP_SUBSTEP.fields_by_name['step'])
_STEP_SUBSTEP.fields_by_name['step'].containing_oneof = _STEP_SUBSTEP.oneofs_by_name['substep']
_STEP_SUBSTEP.oneofs_by_name['substep'].fields.append(
  _STEP_SUBSTEP.fields_by_name['annotation_stream'])
_STEP_SUBSTEP.fields_by_name['annotation_stream'].containing_oneof = _STEP_SUBSTEP.oneofs_by_name['substep']
_STEP_PROGRESS.containing_type = _STEP
_STEP_PROPERTY.containing_type = _STEP
_STEP.fields_by_name['command'].message_type = _STEP_COMMAND
_STEP.fields_by_name['status'].enum_type = _STATUS
_STEP.fields_by_name['failure_details'].message_type = _FAILUREDETAILS
_STEP.fields_by_name['substep'].message_type = _STEP_SUBSTEP
_STEP.fields_by_name['stdout_stream'].message_type = _LOGDOGSTREAM
_STEP.fields_by_name['stderr_stream'].message_type = _LOGDOGSTREAM
_STEP.fields_by_name['started'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_STEP.fields_by_name['ended'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_STEP.fields_by_name['progress'].message_type = _STEP_PROGRESS
_STEP.fields_by_name['link'].message_type = _LINK
_STEP.fields_by_name['other_links'].message_type = _LINK
_STEP.fields_by_name['property'].message_type = _STEP_PROPERTY
_LINK.fields_by_name['logdog_stream'].message_type = _LOGDOGSTREAM
_LINK.fields_by_name['isolate_object'].message_type = _ISOLATEOBJECT
_LINK.fields_by_name['dm_link'].message_type = _DMLINK
_LINK.oneofs_by_name['value'].fields.append(
  _LINK.fields_by_name['url'])
_LINK.fields_by_name['url'].containing_oneof = _LINK.oneofs_by_name['value']
_LINK.oneofs_by_name['value'].fields.append(
  _LINK.fields_by_name['logdog_stream'])
_LINK.fields_by_name['logdog_stream'].containing_oneof = _LINK.oneofs_by_name['value']
_LINK.oneofs_by_name['value'].fields.append(
  _LINK.fields_by_name['isolate_object'])
_LINK.fields_by_name['isolate_object'].containing_oneof = _LINK.oneofs_by_name['value']
_LINK.oneofs_by_name['value'].fields.append(
  _LINK.fields_by_name['dm_link'])
_LINK.fields_by_name['dm_link'].containing_oneof = _LINK.oneofs_by_name['value']
DESCRIPTOR.message_types_by_name['FailureDetails'] = _FAILUREDETAILS
DESCRIPTOR.message_types_by_name['Step'] = _STEP
DESCRIPTOR.message_types_by_name['Link'] = _LINK
DESCRIPTOR.message_types_by_name['LogdogStream'] = _LOGDOGSTREAM
DESCRIPTOR.message_types_by_name['IsolateObject'] = _ISOLATEOBJECT
DESCRIPTOR.message_types_by_name['DMLink'] = _DMLINK
DESCRIPTOR.enum_types_by_name['Status'] = _STATUS

FailureDetails = _reflection.GeneratedProtocolMessageType('FailureDetails', (_message.Message,), dict(
  DESCRIPTOR = _FAILUREDETAILS,
  __module__ = 'annotations_pb2'
  # @@protoc_insertion_point(class_scope:milo.FailureDetails)
  ))
_sym_db.RegisterMessage(FailureDetails)

Step = _reflection.GeneratedProtocolMessageType('Step', (_message.Message,), dict(

  Command = _reflection.GeneratedProtocolMessageType('Command', (_message.Message,), dict(

    EnvironEntry = _reflection.GeneratedProtocolMessageType('EnvironEntry', (_message.Message,), dict(
      DESCRIPTOR = _STEP_COMMAND_ENVIRONENTRY,
      __module__ = 'annotations_pb2'
      # @@protoc_insertion_point(class_scope:milo.Step.Command.EnvironEntry)
      ))
    ,
    DESCRIPTOR = _STEP_COMMAND,
    __module__ = 'annotations_pb2'
    # @@protoc_insertion_point(class_scope:milo.Step.Command)
    ))
  ,

  Substep = _reflection.GeneratedProtocolMessageType('Substep', (_message.Message,), dict(
    DESCRIPTOR = _STEP_SUBSTEP,
    __module__ = 'annotations_pb2'
    # @@protoc_insertion_point(class_scope:milo.Step.Substep)
    ))
  ,

  Progress = _reflection.GeneratedProtocolMessageType('Progress', (_message.Message,), dict(
    DESCRIPTOR = _STEP_PROGRESS,
    __module__ = 'annotations_pb2'
    # @@protoc_insertion_point(class_scope:milo.Step.Progress)
    ))
  ,

  Property = _reflection.GeneratedProtocolMessageType('Property', (_message.Message,), dict(
    DESCRIPTOR = _STEP_PROPERTY,
    __module__ = 'annotations_pb2'
    # @@protoc_insertion_point(class_scope:milo.Step.Property)
    ))
  ,
  DESCRIPTOR = _STEP,
  __module__ = 'annotations_pb2'
  # @@protoc_insertion_point(class_scope:milo.Step)
  ))
_sym_db.RegisterMessage(Step)
_sym_db.RegisterMessage(Step.Command)
_sym_db.RegisterMessage(Step.Command.EnvironEntry)
_sym_db.RegisterMessage(Step.Substep)
_sym_db.RegisterMessage(Step.Progress)
_sym_db.RegisterMessage(Step.Property)

Link = _reflection.GeneratedProtocolMessageType('Link', (_message.Message,), dict(
  DESCRIPTOR = _LINK,
  __module__ = 'annotations_pb2'
  # @@protoc_insertion_point(class_scope:milo.Link)
  ))
_sym_db.RegisterMessage(Link)

LogdogStream = _reflection.GeneratedProtocolMessageType('LogdogStream', (_message.Message,), dict(
  DESCRIPTOR = _LOGDOGSTREAM,
  __module__ = 'annotations_pb2'
  # @@protoc_insertion_point(class_scope:milo.LogdogStream)
  ))
_sym_db.RegisterMessage(LogdogStream)

IsolateObject = _reflection.GeneratedProtocolMessageType('IsolateObject', (_message.Message,), dict(
  DESCRIPTOR = _ISOLATEOBJECT,
  __module__ = 'annotations_pb2'
  # @@protoc_insertion_point(class_scope:milo.IsolateObject)
  ))
_sym_db.RegisterMessage(IsolateObject)

DMLink = _reflection.GeneratedProtocolMessageType('DMLink', (_message.Message,), dict(
  DESCRIPTOR = _DMLINK,
  __module__ = 'annotations_pb2'
  # @@protoc_insertion_point(class_scope:milo.DMLink)
  ))
_sym_db.RegisterMessage(DMLink)


_STEP_COMMAND_ENVIRONENTRY.has_options = True
_STEP_COMMAND_ENVIRONENTRY._options = _descriptor._ParseOptions(descriptor_pb2.MessageOptions(), _b('8\001'))
# @@protoc_insertion_point(module_scope)
