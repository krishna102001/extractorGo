FROM golang:1.24.5-alpine3.22

WORKDIR /app

COPY . /app/

RUN go mod download

RUN  go build -o /myapp

EXPOSE 8080

CMD [ "/myapp" ]