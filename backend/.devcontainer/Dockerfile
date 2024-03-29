# FROM golang:1.17-alpine AS base
# WORKDIR /app

# ENV GO111MODULE="on"
# ENV GOOS="linux"
# ENV CGO_ENABLED=0

# RUN apk update \
#     && apk add --no-cache \
#     ca-certificates \
#     curl \
#     tzdata \
#     git \
#     && update-ca-certificates

# FROM base AS dev
# WORKDIR /app

# RUN go get -u github.com/cosmtrek/air && go install github.com/go-delve/delve/cmd/dlv@latest
# EXPOSE 5000
# EXPOSE 2345

# ENTRYPOINT ["air"]

# FROM base AS builder
# WORKDIR /app

# COPY . /app
# RUN go mod download \
#     && go mod verify

# RUN go build -o todo -a .

# FROM alpine:latest as prod

# COPY --from=builder /app/todo /usr/local/bin/todo
# EXPOSE 5000

# ENTRYPOINT ["/usr/local/bin/todo"]


FROM golang:1.17-buster AS base
# ENV workdir=/go/src/app
# WORKDIR $workdir

# ENV GO111MODULE="on"
# ENV GOOS="linux"
# ENV CGO_ENABLED=0
# RUN go get golang.org/x/tools/gopls@latest
# RUN go get -u github.com/BurntSushi/toml
# RUN go get -u github.com/cosmtrek/air


# COPY go.mod go.sum $workdir
# RUN go mod download
# COPY . .
ENV sourcedir=backend
ENV workdir=/app/backend
WORKDIR $workdir

ENV GO111MODULE="on"
ENV GOOS="linux"
ENV CGO_ENABLED=0

RUN go get golang.org/x/tools/gopls@latest
RUN go get -u github.com/BurntSushi/toml
RUN go get -u github.com/cosmtrek/air@latest
RUN go get -u github.com/go-delve/delve/cmd/dlv@latest


COPY $sourcedir/go.mod $sourcedir/go.sum $workdir
RUN go mod download
COPY . .
