# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: api/v3/api_proto/issues.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import field_mask_pb2 as google_dot_protobuf_dot_field__mask__pb2
from google_proto.google.api import annotations_pb2 as google__proto_dot_google_dot_api_dot_annotations__pb2
from google_proto.google.api import field_behavior_pb2 as google__proto_dot_google_dot_api_dot_field__behavior__pb2
from google_proto.google.api import resource_pb2 as google__proto_dot_google_dot_api_dot_resource__pb2
from api.v3.api_proto import issue_objects_pb2 as api_dot_v3_dot_api__proto_dot_issue__objects__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='api/v3/api_proto/issues.proto',
  package='monorail.v3',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=_b('\n\x1d\x61pi/v3/api_proto/issues.proto\x12\x0bmonorail.v3\x1a google/protobuf/field_mask.proto\x1a)google_proto/google/api/annotations.proto\x1a,google_proto/google/api/field_behavior.proto\x1a&google_proto/google/api/resource.proto\x1a$api/v3/api_proto/issue_objects.proto\"<\n\x0fGetIssueRequest\x12)\n\x04name\x18\x01 \x01(\tB\x1b\xfa\x41\x15\n\x13\x61pi.crbug.com/Issue\xe0\x41\x02\"\x8e\x01\n\x13SearchIssuesRequest\x12/\n\x08projects\x18\x01 \x03(\tB\x1d\xfa\x41\x17\n\x15\x61pi.crbug.com/Project\xe0\x41\x02\x12\r\n\x05query\x18\x02 \x01(\t\x12\x11\n\tpage_size\x18\x03 \x01(\x05\x12\x12\n\npage_token\x18\x04 \x01(\t\x12\x10\n\x08order_by\x18\x05 \x01(\t\"S\n\x14SearchIssuesResponse\x12\"\n\x06issues\x18\x01 \x03(\x0b\x32\x12.monorail.v3.Issue\x12\x17\n\x0fnext_page_token\x18\x02 \x01(\t\"i\n\x13ListCommentsRequest\x12+\n\x06parent\x18\x01 \x01(\tB\x1b\xfa\x41\x15\n\x13\x61pi.crbug.com/Issue\xe0\x41\x02\x12\x11\n\tpage_size\x18\x02 \x01(\x05\x12\x12\n\npage_token\x18\x03 \x01(\t\"W\n\x14ListCommentsResponse\x12&\n\x08\x63omments\x18\x01 \x03(\x0b\x32\x14.monorail.v3.Comment\x12\x17\n\x0fnext_page_token\x18\x02 \x01(\t\"\xa1\x03\n\nIssueDelta\x12>\n\x05issue\x18\x01 \x01(\x0b\x32\x12.monorail.v3.IssueB\x1b\xfa\x41\x15\n\x13\x61pi.crbug.com/Issue\xe0\x41\x02\x12/\n\x0bupdate_mask\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.FieldMask\x12+\n\nccs_remove\x18\x03 \x03(\tB\x17\xfa\x41\x14\n\x12\x61pi.crbug.com/User\x12\x37\n\x18\x62locked_on_issues_remove\x18\x04 \x03(\x0b\x32\x15.monorail.v3.IssueRef\x12\x35\n\x16\x62locking_issues_remove\x18\x05 \x03(\x0b\x32\x15.monorail.v3.IssueRef\x12:\n\x11\x63omponents_remove\x18\x06 \x03(\tB\x1f\xfa\x41\x1c\n\x1a\x61pi.crbug.com/ComponentDef\x12\x15\n\rlabels_remove\x18\x07 \x03(\t\x12\x32\n\x11\x66ield_vals_remove\x18\x08 \x03(\x0b\x32\x17.monorail.v3.FieldValue\"\x8c\x02\n\rApprovalDelta\x12\x32\n\x0e\x61pproval_value\x18\x01 \x01(\x0b\x32\x1a.monorail.v3.ApprovalValue\x12/\n\x0bupdate_mask\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.FieldMask\x12\x31\n\x10\x61pprovers_remove\x18\x03 \x03(\tB\x17\xfa\x41\x14\n\x12\x61pi.crbug.com/User\x12/\n\x0e\x66ield_vals_add\x18\x04 \x03(\x0b\x32\x17.monorail.v3.FieldValue\x12\x32\n\x11\x66ield_vals_remove\x18\x05 \x03(\x0b\x32\x17.monorail.v3.FieldValue\"\x8a\x01\n\x13ModifyIssuesRequest\x12,\n\x06\x64\x65ltas\x18\x01 \x03(\x0b\x32\x17.monorail.v3.IssueDeltaB\x03\xe0\x41\x02\x12,\n\x0bnotify_type\x18\x02 \x01(\x0e\x32\x17.monorail.v3.NotifyType\x12\x17\n\x0f\x63omment_content\x18\x03 \x01(\t\":\n\x14ModifyIssuesResponse\x12\"\n\x06issues\x18\x01 \x03(\x0b\x32\x12.monorail.v3.Issue\"\xd7\x01\n\x1cMakeIssueFromTemplateRequest\x12-\n\x08template\x18\x01 \x01(\tB\x1b\xfa\x41\x18\n\x16\x61pi.crbug.com/Template\x12\x35\n\x14template_issue_delta\x18\x02 \x01(\x0b\x32\x17.monorail.v3.IssueDelta\x12<\n\x18template_approval_deltas\x18\x03 \x03(\x0b\x32\x1a.monorail.v3.ApprovalDelta\x12\x13\n\x0b\x64\x65scription\x18\x04 \x01(\t*I\n\nNotifyType\x12\x1b\n\x17NOTIFY_TYPE_UNSPECIFIED\x10\x00\x12\t\n\x05\x45MAIL\x10\x01\x12\x13\n\x0fNO_NOTIFICATION\x10\x02\x32\xa7\x03\n\x06Issues\x12>\n\x08GetIssue\x12\x1c.monorail.v3.GetIssueRequest\x1a\x12.monorail.v3.Issue\"\x00\x12U\n\x0cSearchIssues\x12 .monorail.v3.SearchIssuesRequest\x1a!.monorail.v3.SearchIssuesResponse\"\x00\x12U\n\x0cListComments\x12 .monorail.v3.ListCommentsRequest\x1a!.monorail.v3.ListCommentsResponse\"\x00\x12U\n\x0cModifyIssues\x12 .monorail.v3.ModifyIssuesRequest\x1a!.monorail.v3.ModifyIssuesResponse\"\x00\x12X\n\x15MakeIssueFromTemplate\x12).monorail.v3.MakeIssueFromTemplateRequest\x1a\x12.monorail.v3.Issue\"\x00\x62\x06proto3')
  ,
  dependencies=[google_dot_protobuf_dot_field__mask__pb2.DESCRIPTOR,google__proto_dot_google_dot_api_dot_annotations__pb2.DESCRIPTOR,google__proto_dot_google_dot_api_dot_field__behavior__pb2.DESCRIPTOR,google__proto_dot_google_dot_api_dot_resource__pb2.DESCRIPTOR,api_dot_v3_dot_api__proto_dot_issue__objects__pb2.DESCRIPTOR,])

_NOTIFYTYPE = _descriptor.EnumDescriptor(
  name='NotifyType',
  full_name='monorail.v3.NotifyType',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='NOTIFY_TYPE_UNSPECIFIED', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='EMAIL', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='NO_NOTIFICATION', index=2, number=2,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=1845,
  serialized_end=1918,
)
_sym_db.RegisterEnumDescriptor(_NOTIFYTYPE)

NotifyType = enum_type_wrapper.EnumTypeWrapper(_NOTIFYTYPE)
NOTIFY_TYPE_UNSPECIFIED = 0
EMAIL = 1
NO_NOTIFICATION = 2



_GETISSUEREQUEST = _descriptor.Descriptor(
  name='GetIssueRequest',
  full_name='monorail.v3.GetIssueRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='monorail.v3.GetIssueRequest.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\372A\025\n\023api.crbug.com/Issue\340A\002'), file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=247,
  serialized_end=307,
)


_SEARCHISSUESREQUEST = _descriptor.Descriptor(
  name='SearchIssuesRequest',
  full_name='monorail.v3.SearchIssuesRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='projects', full_name='monorail.v3.SearchIssuesRequest.projects', index=0,
      number=1, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\372A\027\n\025api.crbug.com/Project\340A\002'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='query', full_name='monorail.v3.SearchIssuesRequest.query', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='page_size', full_name='monorail.v3.SearchIssuesRequest.page_size', index=2,
      number=3, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='page_token', full_name='monorail.v3.SearchIssuesRequest.page_token', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='order_by', full_name='monorail.v3.SearchIssuesRequest.order_by', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=310,
  serialized_end=452,
)


_SEARCHISSUESRESPONSE = _descriptor.Descriptor(
  name='SearchIssuesResponse',
  full_name='monorail.v3.SearchIssuesResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='issues', full_name='monorail.v3.SearchIssuesResponse.issues', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='next_page_token', full_name='monorail.v3.SearchIssuesResponse.next_page_token', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=454,
  serialized_end=537,
)


_LISTCOMMENTSREQUEST = _descriptor.Descriptor(
  name='ListCommentsRequest',
  full_name='monorail.v3.ListCommentsRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='parent', full_name='monorail.v3.ListCommentsRequest.parent', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\372A\025\n\023api.crbug.com/Issue\340A\002'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='page_size', full_name='monorail.v3.ListCommentsRequest.page_size', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='page_token', full_name='monorail.v3.ListCommentsRequest.page_token', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=539,
  serialized_end=644,
)


_LISTCOMMENTSRESPONSE = _descriptor.Descriptor(
  name='ListCommentsResponse',
  full_name='monorail.v3.ListCommentsResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='comments', full_name='monorail.v3.ListCommentsResponse.comments', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='next_page_token', full_name='monorail.v3.ListCommentsResponse.next_page_token', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=646,
  serialized_end=733,
)


_ISSUEDELTA = _descriptor.Descriptor(
  name='IssueDelta',
  full_name='monorail.v3.IssueDelta',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='issue', full_name='monorail.v3.IssueDelta.issue', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\372A\025\n\023api.crbug.com/Issue\340A\002'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='update_mask', full_name='monorail.v3.IssueDelta.update_mask', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='ccs_remove', full_name='monorail.v3.IssueDelta.ccs_remove', index=2,
      number=3, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\372A\024\n\022api.crbug.com/User'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='blocked_on_issues_remove', full_name='monorail.v3.IssueDelta.blocked_on_issues_remove', index=3,
      number=4, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='blocking_issues_remove', full_name='monorail.v3.IssueDelta.blocking_issues_remove', index=4,
      number=5, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='components_remove', full_name='monorail.v3.IssueDelta.components_remove', index=5,
      number=6, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\372A\034\n\032api.crbug.com/ComponentDef'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='labels_remove', full_name='monorail.v3.IssueDelta.labels_remove', index=6,
      number=7, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='field_vals_remove', full_name='monorail.v3.IssueDelta.field_vals_remove', index=7,
      number=8, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=736,
  serialized_end=1153,
)


_APPROVALDELTA = _descriptor.Descriptor(
  name='ApprovalDelta',
  full_name='monorail.v3.ApprovalDelta',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='approval_value', full_name='monorail.v3.ApprovalDelta.approval_value', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='update_mask', full_name='monorail.v3.ApprovalDelta.update_mask', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='approvers_remove', full_name='monorail.v3.ApprovalDelta.approvers_remove', index=2,
      number=3, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\372A\024\n\022api.crbug.com/User'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='field_vals_add', full_name='monorail.v3.ApprovalDelta.field_vals_add', index=3,
      number=4, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='field_vals_remove', full_name='monorail.v3.ApprovalDelta.field_vals_remove', index=4,
      number=5, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1156,
  serialized_end=1424,
)


_MODIFYISSUESREQUEST = _descriptor.Descriptor(
  name='ModifyIssuesRequest',
  full_name='monorail.v3.ModifyIssuesRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='deltas', full_name='monorail.v3.ModifyIssuesRequest.deltas', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\340A\002'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='notify_type', full_name='monorail.v3.ModifyIssuesRequest.notify_type', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='comment_content', full_name='monorail.v3.ModifyIssuesRequest.comment_content', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1427,
  serialized_end=1565,
)


_MODIFYISSUESRESPONSE = _descriptor.Descriptor(
  name='ModifyIssuesResponse',
  full_name='monorail.v3.ModifyIssuesResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='issues', full_name='monorail.v3.ModifyIssuesResponse.issues', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1567,
  serialized_end=1625,
)


_MAKEISSUEFROMTEMPLATEREQUEST = _descriptor.Descriptor(
  name='MakeIssueFromTemplateRequest',
  full_name='monorail.v3.MakeIssueFromTemplateRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='template', full_name='monorail.v3.MakeIssueFromTemplateRequest.template', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\372A\030\n\026api.crbug.com/Template'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='template_issue_delta', full_name='monorail.v3.MakeIssueFromTemplateRequest.template_issue_delta', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='template_approval_deltas', full_name='monorail.v3.MakeIssueFromTemplateRequest.template_approval_deltas', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='description', full_name='monorail.v3.MakeIssueFromTemplateRequest.description', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1628,
  serialized_end=1843,
)

_SEARCHISSUESRESPONSE.fields_by_name['issues'].message_type = api_dot_v3_dot_api__proto_dot_issue__objects__pb2._ISSUE
_LISTCOMMENTSRESPONSE.fields_by_name['comments'].message_type = api_dot_v3_dot_api__proto_dot_issue__objects__pb2._COMMENT
_ISSUEDELTA.fields_by_name['issue'].message_type = api_dot_v3_dot_api__proto_dot_issue__objects__pb2._ISSUE
_ISSUEDELTA.fields_by_name['update_mask'].message_type = google_dot_protobuf_dot_field__mask__pb2._FIELDMASK
_ISSUEDELTA.fields_by_name['blocked_on_issues_remove'].message_type = api_dot_v3_dot_api__proto_dot_issue__objects__pb2._ISSUEREF
_ISSUEDELTA.fields_by_name['blocking_issues_remove'].message_type = api_dot_v3_dot_api__proto_dot_issue__objects__pb2._ISSUEREF
_ISSUEDELTA.fields_by_name['field_vals_remove'].message_type = api_dot_v3_dot_api__proto_dot_issue__objects__pb2._FIELDVALUE
_APPROVALDELTA.fields_by_name['approval_value'].message_type = api_dot_v3_dot_api__proto_dot_issue__objects__pb2._APPROVALVALUE
_APPROVALDELTA.fields_by_name['update_mask'].message_type = google_dot_protobuf_dot_field__mask__pb2._FIELDMASK
_APPROVALDELTA.fields_by_name['field_vals_add'].message_type = api_dot_v3_dot_api__proto_dot_issue__objects__pb2._FIELDVALUE
_APPROVALDELTA.fields_by_name['field_vals_remove'].message_type = api_dot_v3_dot_api__proto_dot_issue__objects__pb2._FIELDVALUE
_MODIFYISSUESREQUEST.fields_by_name['deltas'].message_type = _ISSUEDELTA
_MODIFYISSUESREQUEST.fields_by_name['notify_type'].enum_type = _NOTIFYTYPE
_MODIFYISSUESRESPONSE.fields_by_name['issues'].message_type = api_dot_v3_dot_api__proto_dot_issue__objects__pb2._ISSUE
_MAKEISSUEFROMTEMPLATEREQUEST.fields_by_name['template_issue_delta'].message_type = _ISSUEDELTA
_MAKEISSUEFROMTEMPLATEREQUEST.fields_by_name['template_approval_deltas'].message_type = _APPROVALDELTA
DESCRIPTOR.message_types_by_name['GetIssueRequest'] = _GETISSUEREQUEST
DESCRIPTOR.message_types_by_name['SearchIssuesRequest'] = _SEARCHISSUESREQUEST
DESCRIPTOR.message_types_by_name['SearchIssuesResponse'] = _SEARCHISSUESRESPONSE
DESCRIPTOR.message_types_by_name['ListCommentsRequest'] = _LISTCOMMENTSREQUEST
DESCRIPTOR.message_types_by_name['ListCommentsResponse'] = _LISTCOMMENTSRESPONSE
DESCRIPTOR.message_types_by_name['IssueDelta'] = _ISSUEDELTA
DESCRIPTOR.message_types_by_name['ApprovalDelta'] = _APPROVALDELTA
DESCRIPTOR.message_types_by_name['ModifyIssuesRequest'] = _MODIFYISSUESREQUEST
DESCRIPTOR.message_types_by_name['ModifyIssuesResponse'] = _MODIFYISSUESRESPONSE
DESCRIPTOR.message_types_by_name['MakeIssueFromTemplateRequest'] = _MAKEISSUEFROMTEMPLATEREQUEST
DESCRIPTOR.enum_types_by_name['NotifyType'] = _NOTIFYTYPE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

GetIssueRequest = _reflection.GeneratedProtocolMessageType('GetIssueRequest', (_message.Message,), dict(
  DESCRIPTOR = _GETISSUEREQUEST,
  __module__ = 'api.v3.api_proto.issues_pb2'
  # @@protoc_insertion_point(class_scope:monorail.v3.GetIssueRequest)
  ))
_sym_db.RegisterMessage(GetIssueRequest)

SearchIssuesRequest = _reflection.GeneratedProtocolMessageType('SearchIssuesRequest', (_message.Message,), dict(
  DESCRIPTOR = _SEARCHISSUESREQUEST,
  __module__ = 'api.v3.api_proto.issues_pb2'
  # @@protoc_insertion_point(class_scope:monorail.v3.SearchIssuesRequest)
  ))
_sym_db.RegisterMessage(SearchIssuesRequest)

SearchIssuesResponse = _reflection.GeneratedProtocolMessageType('SearchIssuesResponse', (_message.Message,), dict(
  DESCRIPTOR = _SEARCHISSUESRESPONSE,
  __module__ = 'api.v3.api_proto.issues_pb2'
  # @@protoc_insertion_point(class_scope:monorail.v3.SearchIssuesResponse)
  ))
_sym_db.RegisterMessage(SearchIssuesResponse)

ListCommentsRequest = _reflection.GeneratedProtocolMessageType('ListCommentsRequest', (_message.Message,), dict(
  DESCRIPTOR = _LISTCOMMENTSREQUEST,
  __module__ = 'api.v3.api_proto.issues_pb2'
  # @@protoc_insertion_point(class_scope:monorail.v3.ListCommentsRequest)
  ))
_sym_db.RegisterMessage(ListCommentsRequest)

ListCommentsResponse = _reflection.GeneratedProtocolMessageType('ListCommentsResponse', (_message.Message,), dict(
  DESCRIPTOR = _LISTCOMMENTSRESPONSE,
  __module__ = 'api.v3.api_proto.issues_pb2'
  # @@protoc_insertion_point(class_scope:monorail.v3.ListCommentsResponse)
  ))
_sym_db.RegisterMessage(ListCommentsResponse)

IssueDelta = _reflection.GeneratedProtocolMessageType('IssueDelta', (_message.Message,), dict(
  DESCRIPTOR = _ISSUEDELTA,
  __module__ = 'api.v3.api_proto.issues_pb2'
  # @@protoc_insertion_point(class_scope:monorail.v3.IssueDelta)
  ))
_sym_db.RegisterMessage(IssueDelta)

ApprovalDelta = _reflection.GeneratedProtocolMessageType('ApprovalDelta', (_message.Message,), dict(
  DESCRIPTOR = _APPROVALDELTA,
  __module__ = 'api.v3.api_proto.issues_pb2'
  # @@protoc_insertion_point(class_scope:monorail.v3.ApprovalDelta)
  ))
_sym_db.RegisterMessage(ApprovalDelta)

ModifyIssuesRequest = _reflection.GeneratedProtocolMessageType('ModifyIssuesRequest', (_message.Message,), dict(
  DESCRIPTOR = _MODIFYISSUESREQUEST,
  __module__ = 'api.v3.api_proto.issues_pb2'
  # @@protoc_insertion_point(class_scope:monorail.v3.ModifyIssuesRequest)
  ))
_sym_db.RegisterMessage(ModifyIssuesRequest)

ModifyIssuesResponse = _reflection.GeneratedProtocolMessageType('ModifyIssuesResponse', (_message.Message,), dict(
  DESCRIPTOR = _MODIFYISSUESRESPONSE,
  __module__ = 'api.v3.api_proto.issues_pb2'
  # @@protoc_insertion_point(class_scope:monorail.v3.ModifyIssuesResponse)
  ))
_sym_db.RegisterMessage(ModifyIssuesResponse)

MakeIssueFromTemplateRequest = _reflection.GeneratedProtocolMessageType('MakeIssueFromTemplateRequest', (_message.Message,), dict(
  DESCRIPTOR = _MAKEISSUEFROMTEMPLATEREQUEST,
  __module__ = 'api.v3.api_proto.issues_pb2'
  # @@protoc_insertion_point(class_scope:monorail.v3.MakeIssueFromTemplateRequest)
  ))
_sym_db.RegisterMessage(MakeIssueFromTemplateRequest)


_GETISSUEREQUEST.fields_by_name['name']._options = None
_SEARCHISSUESREQUEST.fields_by_name['projects']._options = None
_LISTCOMMENTSREQUEST.fields_by_name['parent']._options = None
_ISSUEDELTA.fields_by_name['issue']._options = None
_ISSUEDELTA.fields_by_name['ccs_remove']._options = None
_ISSUEDELTA.fields_by_name['components_remove']._options = None
_APPROVALDELTA.fields_by_name['approvers_remove']._options = None
_MODIFYISSUESREQUEST.fields_by_name['deltas']._options = None
_MAKEISSUEFROMTEMPLATEREQUEST.fields_by_name['template']._options = None

_ISSUES = _descriptor.ServiceDescriptor(
  name='Issues',
  full_name='monorail.v3.Issues',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=1921,
  serialized_end=2344,
  methods=[
  _descriptor.MethodDescriptor(
    name='GetIssue',
    full_name='monorail.v3.Issues.GetIssue',
    index=0,
    containing_service=None,
    input_type=_GETISSUEREQUEST,
    output_type=api_dot_v3_dot_api__proto_dot_issue__objects__pb2._ISSUE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='SearchIssues',
    full_name='monorail.v3.Issues.SearchIssues',
    index=1,
    containing_service=None,
    input_type=_SEARCHISSUESREQUEST,
    output_type=_SEARCHISSUESRESPONSE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='ListComments',
    full_name='monorail.v3.Issues.ListComments',
    index=2,
    containing_service=None,
    input_type=_LISTCOMMENTSREQUEST,
    output_type=_LISTCOMMENTSRESPONSE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='ModifyIssues',
    full_name='monorail.v3.Issues.ModifyIssues',
    index=3,
    containing_service=None,
    input_type=_MODIFYISSUESREQUEST,
    output_type=_MODIFYISSUESRESPONSE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='MakeIssueFromTemplate',
    full_name='monorail.v3.Issues.MakeIssueFromTemplate',
    index=4,
    containing_service=None,
    input_type=_MAKEISSUEFROMTEMPLATEREQUEST,
    output_type=api_dot_v3_dot_api__proto_dot_issue__objects__pb2._ISSUE,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_ISSUES)

DESCRIPTOR.services_by_name['Issues'] = _ISSUES

# @@protoc_insertion_point(module_scope)
