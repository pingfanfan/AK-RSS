FROM golang:1.22-alpine AS builder

WORKDIR /src
COPY go.mod ./
RUN go mod download
COPY . .
RUN go mod tidy
RUN go build -o /out/opmlwatch ./cmd/opmlwatch

FROM alpine:3.20
RUN adduser -D -u 10001 appuser
WORKDIR /app
COPY --from=builder /out/opmlwatch /usr/local/bin/opmlwatch
USER appuser
ENTRYPOINT ["opmlwatch"]
CMD ["run", "--config", "examples/config.digest.yaml", "--dry-run"]
