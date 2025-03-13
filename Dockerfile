FROM golang:1.21

WORKDIR /app

COPY . .

RUN apt-get update && apt-get install -y gcc libc6-dev

ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GOARCH=amd64

RUN go build -o main .

CMD ["/app/main"]
