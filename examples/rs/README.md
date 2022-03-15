## Rust healthcheck service example

To add Healthcheck feature to your Rust service, you have to implement the `HealthcheckService` trait for some of your structure:

```rust
use botapi::healthcheck::healthcheck_service_server::{
    HealthcheckService, HealthcheckServiceServer,
};
use botapi::healthcheck::CheckResponse;

pub struct Service {}

#[tonic::async_trait]
impl HealthcheckService for Service {
    async fn check(&self, _request: Request<()>) -> Result<Response<CheckResponse>, Status> {
        log::debug!("processing check call");
        Ok(tonic::Response::new(CheckResponse {
            message: "ok".to_string(),
        }))
    }
}

```
