FROM golang:1.24-alpine

WORKDIR /usr/src/app


RUN go mod download

COPY . .
RUN go build -v -o /usr/local/bin/app ./...

CMD ["app"]
