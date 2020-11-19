FROM golang:alpine as build-env
WORKDIR /app
ADD . /app
RUN cd /app && go build -o /app/cpfdv

FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
WORKDIR /app
COPY --from=build-env /app/cpfdv /app

EXPOSE 8080
ENTRYPOINT ./cpfdv
