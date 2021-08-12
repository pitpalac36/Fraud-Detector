FROM golang:1.16

WORKDIR /processor

ENV PRODUCTION 1
ENV NORM_URL ws://normalizer:8082
ENV AI_URL ws://ai:8083

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

EXPOSE 8081

RUN go build -o ./processor

CMD ["./processor"]