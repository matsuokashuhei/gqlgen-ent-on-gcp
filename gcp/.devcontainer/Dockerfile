FROM golang:1.17.3 AS builder

ENV terraform_version=v1.1.7

ENV GO111MODULE="on"
ENV GOOS="linux"
ENV CGO_ENABLED=0
RUN git clone https://github.com/hashicorp/terraform.git -b $terraform_version /tmp/terraform && \
    cd /tmp/terraform && \
    go build -o /usr/local/bin/terraform

FROM alpine:latest as terraform
COPY --from=builder /usr/local/bin/terraform /usr/local/bin/terraform
