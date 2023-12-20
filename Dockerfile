FROM golang:latest

EXPOSE 8081

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o main src/main.go

CMD [ "./main" ]