FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o webapp .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/webapp .
COPY --from=builder /app/public ./public
EXPOSE 3000
CMD ["./webapp"]