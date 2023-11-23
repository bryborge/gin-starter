FROM golang:1.21.3 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main

# ------------------------------------------------------------------------------

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app .

EXPOSE 8000

CMD ["./main"]
