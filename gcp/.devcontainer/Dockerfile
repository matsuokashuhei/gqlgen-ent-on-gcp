FROM golang:1.17.3 AS builder

ENV terraform_version=v1.1.9

ENV GO111MODULE="on"
ENV GOOS="linux"
ENV GOARCH="arm64"
ENV CGO_ENABLED=0
RUN git clone https://github.com/hashicorp/terraform.git -b $terraform_version /tmp/terraform && \
    cd /tmp/terraform && \
    go build -o /usr/local/bin/terraform

ENV TERRAGRUNT_VERSION=v0.36.7
RUN curl -L https://github.com/gruntwork-io/terragrunt/releases/download/${TERRAGRUNT_VERSION}/terragrunt_linux_arm64 > terragrunt && \
    mv terragrunt /usr/local/bin/terragrunt && \
    chmod +x /usr/local/bin/terragrunt

FROM alpine:latest as terraform
COPY --from=builder /usr/local/bin/terraform /usr/local/bin/terraform
COPY --from=builder /usr/local/bin/terragrunt /usr/local/bin/terragrunt

RUN apk --update add curl git && rm -rf /var/cache/apk/*
CMD ["git","--version"]
