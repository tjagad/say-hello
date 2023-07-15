FROM golang:1.20.6 as builder
WORKDIR /say-hello/
COPY . .
RUN CGO_ENABLED=0 go build -o say-hello /say-hello/*.go

FROM alpine:latest
WORKDIR /say-hello
COPY --from=builder /say-hello/say-hello .
EXPOSE 9080

CMD ./say-hello
