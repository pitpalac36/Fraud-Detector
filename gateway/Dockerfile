FROM golang:1.16

WORKDIR /gateway

ENV PRODUCTION=1

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

EXPOSE 8080

RUN go build -o ./gateway

CMD ["./gateway"]