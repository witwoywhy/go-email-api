FROM golang:1.21.6 as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build \
    -o goapp

################################################
FROM alpine:latest

RUN apk add tzdata
ENV TZ=Asia/Bangkok

WORKDIR /app

COPY ./configs ./configs
COPY --from=builder /app/goapp ./goapp

EXPOSE 8080

CMD ["./goapp"]