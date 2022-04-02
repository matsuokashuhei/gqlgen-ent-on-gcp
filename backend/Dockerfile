FROM golang:1.17-buster AS builder

ENV GO111MODULE="on"
ENV GOOS="linux"
ENV CGO_ENABLED=0

WORKDIR /app
COPY . /app
RUN go mod download && go mod tidy && cd cmd/app && go build -o landin -a .

FROM alpine:latest
COPY --from=builder /app/cmd/app/landin /usr/local/bin/landin
EXPOSE 8080
ENTRYPOINT [ "/usr/local/bin/landin" ]
