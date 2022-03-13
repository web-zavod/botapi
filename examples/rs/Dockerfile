FROM rust:1.59 as build

RUN cargo new --bin healthcheck

WORKDIR healthcheck

COPY Cargo.toml Cargo.toml

RUN rustup component add rustfmt

RUN mkdir target
RUN cargo build --release --target-dir target/

RUN rm src/*.rs

COPY src src

RUN rm ./target/release/deps/healthcheck*
RUN cargo build --release --target-dir target/

FROM debian:buster-slim as release

COPY --from=build /healthcheck/target/release/healthcheck .

CMD ["./healthcheck"]
