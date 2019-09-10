FROM golang:1.12.0-alpine3.9

RUN mkdir /app
WORKDIR /app
RUN apk add --update --no-cache git
RUN git clone https://github.com/vortgo/ma-api ./
RUN go mod download
RUN go build -o main .
COPY .env ./

CMD ["/app/main"]

EXPOSE 1323