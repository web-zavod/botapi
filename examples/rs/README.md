## Rust healthcheck server example

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

## Rust healthcheck client example

```rust
use std::env;
use std::error::Error;

use lazy_static::lazy_static;
use tonic::transport::Channel;

use botapi::healthcheck::healthcheck_service_client::HealthcheckServiceClient;

lazy_static! {
    static ref GRPC_SERVER_ADDRESS: String =
        env::var("GRPC_SERVER_ADDRESS").expect("GRPC_SERVER_ADDRESS");
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn Error + 'static>> {
    log::info!("Starting client");

    let healthcheck_chan = Channel::from_static(&GRPC_SERVER_ADDRESS)
        .connect()
        .await
        .expect("can't create a channel");

    let mut healthcheck_svc = HealthcheckServiceClient::new(healthcheck_chan);

    let test = healthcheck_svc.check(tonic::Request::new(())).await?;
    log::info!("{}", test.into_inner().message);

    Ok(())
}

```
