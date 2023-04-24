FROM rust:1.67.0 as builder
RUN apt-get update && apt-get install -y musl-tools
WORKDIR /usr/src/hyper_server
COPY . .

RUN rustup target add x86_64-unknown-linux-musl
RUN cargo build --release --target x86_64-unknown-linux-musl

RUN ls -lh
RUN ls -lh target/release/
RUN ldd target/release/hyper_server

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /usr/local/bin/
COPY --from=builder /usr/src/hyper_server/target/x86_64-unknown-linux-musl/release/hyper_server .
CMD ["./hyper_server"]
