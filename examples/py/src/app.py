import logging
from typing import Callable, NewType, Type
from importlib import import_module

import grpc
from grpc_reflection.v1alpha import reflection

Service = NewType("Service", object)

logger = logging.getLogger(__name__)

class Application:
    def __init__(self, grpc_port="[::]:5000"):
        logger.info('Preparing a gRPC server...')

        self.grpc_server = grpc.aio.server()
        self.grpc_port = grpc_port

        self._service_names: list[Service] = []

    async def add_service(
            self,
            service_class: Type[Service],
            contract_name: str = "example_name",
            version: str = "1",
            ):

        class_name = service_class.__name__
        method_name = f"add_{class_name}Servicer_to_server"
        grpc_module_name = f"{camel_to_snake(class_name)}_pb2_grpc"
        meta_module_name = f"{camel_to_snake(class_name)}_pb2"

        logger.debug(
            'Calling method `%s` for service `%s` from `%s` module...',
            method_name,
            class_name,
            grpc_module_name,
        )

        grpc_module = import_module(
            f'botapi.{contract_name}.v{version}.{grpc_module_name}')
        call: Callable = getattr(grpc_module, method_name)
        call(service_class(), self.grpc_server)

        # Add service to reflection API

        meta_module = import_module(
            f'botapi.{contract_name}.v{version}.{meta_module_name}')

        service_name = \
            meta_module.DESCRIPTOR.services_by_name[class_name].full_name


        self._service_names.append(service_name)

        logger.info('Service `%s` has been added.', class_name)


    async def run(self):
        """Run the application."""

        logger.info('Preparing reflection API...')
        self._service_names.append(reflection.SERVICE_NAME)
        reflection.enable_server_reflection(
            self._service_names,
            self.grpc_server
        )

        logger.info('Opening a port for RPCs...')
        self.grpc_server.add_insecure_port(self.grpc_port)

        logger.info('Starting the gRPC server...')
        await self.grpc_server.start()

        logger.info('Listening to %s.', self.grpc_port)
        return await self.grpc_server.wait_for_termination()

def camel_to_snake(s):
    return ''.join(['_'+c.lower() if c.isupper() else c for c in s]).lstrip('_')
