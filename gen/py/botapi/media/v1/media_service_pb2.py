# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: botapi/media/v1/media_service.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='botapi/media/v1/media_service.proto',
  package='botapi.media.v1',
  syntax='proto3',
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n#botapi/media/v1/media_service.proto\x12\x0f\x62otapi.media.v1\"-\n\x0cVoiceMessage\x12\x0f\n\x07user_id\x18\x01 \x01(\x03\x12\x0c\n\x04\x64\x61ta\x18\x02 \x01(\x0c\",\n\x0bTextMessage\x12\x0f\n\x07user_id\x18\x01 \x01(\x03\x12\x0c\n\x04text\x18\x02 \x01(\t2a\n\x0cMediaService\x12Q\n\x10GetTextFromVoice\x12\x1d.botapi.media.v1.VoiceMessage\x1a\x1c.botapi.media.v1.TextMessage\"\x00\x62\x06proto3'
)




_VOICEMESSAGE = _descriptor.Descriptor(
  name='VoiceMessage',
  full_name='botapi.media.v1.VoiceMessage',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='user_id', full_name='botapi.media.v1.VoiceMessage.user_id', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='data', full_name='botapi.media.v1.VoiceMessage.data', index=1,
      number=2, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=b"",
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=56,
  serialized_end=101,
)


_TEXTMESSAGE = _descriptor.Descriptor(
  name='TextMessage',
  full_name='botapi.media.v1.TextMessage',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='user_id', full_name='botapi.media.v1.TextMessage.user_id', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='text', full_name='botapi.media.v1.TextMessage.text', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=103,
  serialized_end=147,
)

DESCRIPTOR.message_types_by_name['VoiceMessage'] = _VOICEMESSAGE
DESCRIPTOR.message_types_by_name['TextMessage'] = _TEXTMESSAGE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

VoiceMessage = _reflection.GeneratedProtocolMessageType('VoiceMessage', (_message.Message,), {
  'DESCRIPTOR' : _VOICEMESSAGE,
  '__module__' : 'botapi.media.v1.media_service_pb2'
  # @@protoc_insertion_point(class_scope:botapi.media.v1.VoiceMessage)
  })
_sym_db.RegisterMessage(VoiceMessage)

TextMessage = _reflection.GeneratedProtocolMessageType('TextMessage', (_message.Message,), {
  'DESCRIPTOR' : _TEXTMESSAGE,
  '__module__' : 'botapi.media.v1.media_service_pb2'
  # @@protoc_insertion_point(class_scope:botapi.media.v1.TextMessage)
  })
_sym_db.RegisterMessage(TextMessage)



_MEDIASERVICE = _descriptor.ServiceDescriptor(
  name='MediaService',
  full_name='botapi.media.v1.MediaService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=149,
  serialized_end=246,
  methods=[
  _descriptor.MethodDescriptor(
    name='GetTextFromVoice',
    full_name='botapi.media.v1.MediaService.GetTextFromVoice',
    index=0,
    containing_service=None,
    input_type=_VOICEMESSAGE,
    output_type=_TEXTMESSAGE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_MEDIASERVICE)

DESCRIPTOR.services_by_name['MediaService'] = _MEDIASERVICE

# @@protoc_insertion_point(module_scope)
