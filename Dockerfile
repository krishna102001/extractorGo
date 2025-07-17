FROM golang:1.24.5-alpine3.22

WORKDIR /app

COPY go* /app/

RUN go mod download

COPY *.go /app/

RUN CGO_ENABLED=0 GOOS=linux go build -o /myapp

EXPOSE 8080

CMD [ "/myapp" ]