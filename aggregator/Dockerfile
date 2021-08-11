FROM golang:1.16

WORKDIR /aggregator

ENV PRODUCTION=1

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

EXPOSE 8084

RUN go build -o ./aggregator

CMD ["./aggregator"]