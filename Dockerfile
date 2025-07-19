FROM golang:1.24.5-alpine3.22 AS build-stage

WORKDIR /app

COPY . /app/

RUN go mod tidy

RUN  CGO_ENABLED=0 GOOS=linux go build -o /myapp


#second stage
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /myapp /myapp

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT [ "./myapp" ]