FROM golang:1.21.1

WORKDIR /usr/src/app

COPY . .

RUN rm .env
RUN mv .env.production .env

RUN go build

CMD ["./orderService"]