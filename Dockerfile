FROM golang:1.23.4-alpine

WORKDIR /

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd

EXPOSE 8081

CMD ["./main"]
