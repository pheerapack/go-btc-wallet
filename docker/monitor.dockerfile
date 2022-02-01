FROM golang:1.17-alpine AS builder
WORKDIR /app

COPY . .

RUN go build cmd/monitor/main.go

RUN ls -lrt

FROM alpine:latest  

WORKDIR /application
COPY --from=builder /app/docker/script/running-monitor.sh /application/running.sh
COPY --from=builder /app/main /application/btcwallet-monitor

ENV PATH "$PATH:/application"

RUN chmod +x -R /application
RUN ls -lrt /application
ENTRYPOINT ["sh","/application/running.sh"]