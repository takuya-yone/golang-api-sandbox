FROM golang:1.19.0-alpine AS builder
RUN mkdir /work
WORKDIR /work
COPY . /work
RUN go build main.go

FROM alpine:latest
RUN apk update
RUN apk add curl --no-cache
COPY --from=builder /work/main .
EXPOSE 8080
ENTRYPOINT ./main