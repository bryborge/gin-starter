FROM golang:1.21.3

WORKDIR /app

COPY . .

RUN go build -o main main.go

CMD ["./main"]
