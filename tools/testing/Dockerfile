# server build stage
FROM golang:1.13.0 AS app-build-env
ARG SERVICE_NAME
ADD . /src
WORKDIR /src
RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/$SERVICE_NAME/main.go

# final stage
FROM alpine:3.10.1
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
ARG SERVICE_NAME
WORKDIR /app
COPY --from=app-build-env /src/app .
ENTRYPOINT ["./app"]

EXPOSE 8081