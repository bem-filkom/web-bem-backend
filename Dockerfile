FROM golang:1.23.2 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/app

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/main .

ENTRYPOINT [ "/app/main" ]