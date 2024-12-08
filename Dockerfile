FROM golang:1.23.1-alpine AS builder

WORKDIR /app

USER root

COPY ["go.mod", "go.sum", "./"]
RUN --mount=type=cache,target=/go/pkg/mod/ \
    go mod download -x
COPY . .

RUN go build -o /migrate ./cmd/migration/main.go
RUN go build -o /booking ./cmd/api/main.go

FROM alpine AS runner

COPY --from=builder /migrate /
COPY --from=builder /booking /
COPY web /web

EXPOSE 4000

CMD ["ash", "-c", "/migrate;/booking"]
