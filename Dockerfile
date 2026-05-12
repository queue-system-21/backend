FROM golang:1.26.2-alpine3.23

ARG port

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o queue_bin

EXPOSE $port

CMD [ "/app/queue_bin" ]