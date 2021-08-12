FROM golang:1.16

WORKDIR /aggregator

ENV PRODUCTION 1
ENV REDIS_ADDR redis:6379
ENV DENORM_URL ws://denormalizer:8085
ENV DASHBOARD_URL ws://dashboard:8086

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

EXPOSE 8084

RUN go build -o ./aggregator

CMD ["./aggregator"]