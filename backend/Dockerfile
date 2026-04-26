# ========= BUILD STAGE =========
FROM golang:1.24-alpine AS builder

WORKDIR /build

RUN apk add --no-cache gcc musl-dev

ENV CGO_ENABLED=1

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app main.go

FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache libc6-compat

COPY --from=builder /build/app .

RUN chmod +x /app/app

EXPOSE 8080

CMD ["/app/app"]