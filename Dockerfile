FROM golang:1.21

WORKDIR /app

COPY . .

RUN go build -o myapi

EXPOSE 8083

CMD ["./myapi"]