FROM golang:1.20 AS build

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o chat-app .

CMD ["./chat-app"]
