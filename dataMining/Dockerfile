# -------- Build stage --------
FROM golang:1.25-alpine AS builder

WORKDIR /app

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

RUN apk add --no-cache \
    openjdk21 \
    maven \
    git \
    curl \
    bash \
    ca-certificates \
    github-cli \
    docker-cli

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o prequal

# -------- Runtime stage --------
FROM alpine:3.23.3

WORKDIR /app

RUN apk add --no-cache \
    openjdk21 \
    maven \
    git \
    curl \
    bash \
    ca-certificates \
    github-cli \
    docker-cli

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/prequal /usr/local/bin/prequal

VOLUME ["/app/workspace"]

ENTRYPOINT ["prequal"]
