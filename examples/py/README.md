## Python healthcheck service example
To add Healthcheck feature to your Python service, you have to inherit some class from [HealthcheckServiceServicer](https://github.com/web-zavod/botapi/blob/master/gen/py/botapi/healthcheck/v1/healthcheck_service_pb2_grpc.py) that generated from [Hejalthcheck_service.proto](https://github.com/web-zavod/botapi/blob/master/proto/botapi/healthcheck/v1/healthcheck_service.proto) and implement `Check` method:


##### Interface
```python
from botapi.healthcheck.v1.healthcheck_service_pb2_grpc import HealthcheckServiceServicer


class HealthcheckServiceServicer(object):
    """Service for health check request.
    """

    def Check(self, request, context):
        """check method returns any string to validate that service is running and
        grpc interface operates
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')
```


##### Implementation
```python
from botapi.healthcheck.v1.healthcheck_service_pb2 import CheckResponse


class HealthcheckService(HealthcheckServiceServicer):
    async def Check(
            self,
            request: Empty,
            context: grpc.ServicerContext,
            ):
        """Method to check health of service"""
        logger.info("Healthcheck call")
        return CheckResponse(message="Ok!")
```

After this, you can check if method is working by calling it via [gRPC OX](http://localhost:6969) after executing:

```bash
docker-compose up
```
