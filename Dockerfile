FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o fake-twitter ./cmd/api

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/fake-twitter .
EXPOSE 8080

CMD ["./fake-twitter"]