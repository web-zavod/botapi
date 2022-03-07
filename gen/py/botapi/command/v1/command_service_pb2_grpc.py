# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from botapi.command.v1 import command_service_pb2 as botapi_dot_command_dot_v1_dot_command__service__pb2


class CommandServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetReply = channel.unary_unary(
                '/botapi.command.v1.CommandService/GetReply',
                request_serializer=botapi_dot_command_dot_v1_dot_command__service__pb2.IncomingMessage.SerializeToString,
                response_deserializer=botapi_dot_command_dot_v1_dot_command__service__pb2.ReplyMessage.FromString,
                )


class CommandServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def GetReply(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_CommandServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'GetReply': grpc.unary_unary_rpc_method_handler(
                    servicer.GetReply,
                    request_deserializer=botapi_dot_command_dot_v1_dot_command__service__pb2.IncomingMessage.FromString,
                    response_serializer=botapi_dot_command_dot_v1_dot_command__service__pb2.ReplyMessage.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'botapi.command.v1.CommandService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class CommandService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def GetReply(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/botapi.command.v1.CommandService/GetReply',
            botapi_dot_command_dot_v1_dot_command__service__pb2.IncomingMessage.SerializeToString,
            botapi_dot_command_dot_v1_dot_command__service__pb2.ReplyMessage.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)