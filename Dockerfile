FROM golang:1.17 AS builder

WORKDIR /src
COPY . .
RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux go build -o /gh-api-demo -a -ldflags '-linkmode external -extldflags "-static"' ./cmd/cli

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /gh-api-demo /gh-api-demo

ENTRYPOINT ["/gh-api-demo"]
