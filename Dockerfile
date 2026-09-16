FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o webapp .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/webapp .
COPY --from=builder /app/public ./public
EXPOSE 3000
CMD ["./webapp"]