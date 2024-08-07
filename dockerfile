FROM --platform=linux/amd64 ignitehq/cli:v28.5.1 AS ignite
FROM --platform=linux/amd64 golang:1.22.5-bookworm AS builder
COPY --from=ignite /usr/bin/ignite /usr/bin/ignite
USER root
WORKDIR /build/
COPY ./ /build/
RUN ignite chain build \
    --release.targets linux:amd64 \
    --output ./release \
    --release
RUN tar -zxvf release/concord_linux_amd64.tar.gz

FROM --platform=linux/amd64 ubuntu
ENV LD_LIBRARY_PATH=/usr/local/lib
RUN apt-get update \
 && apt-get install -y ca-certificates vim curl jq
WORKDIR /root/
RUN curl -fL 'https://github.com/CosmWasm/wasmvm/releases/download/v2.1.0/libwasmvm.x86_64.so' > /usr/local/lib/libwasmvm.x86_64.so
COPY --from=builder /build/concordd /usr/local/bin
CMD ["concordd", "start"]
