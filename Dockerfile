FROM golang:1.16-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /main

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /tmp

ENV TZ="Europe/Berlin"
ENV USERNAME="YOUR_USERNAME_GOES_HERE"
ENV PASSWORD="YOUR_PASSWORD_GOES_HERE"

COPY --from=build /main /webdav

EXPOSE 80

USER nonroot:nonroot

ENTRYPOINT ["/webdav"]