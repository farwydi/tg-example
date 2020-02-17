FROM golang:alpine as builder
RUN apk --update --no-cache add bash
WORKDIR /app
ADD . .
RUN go build -o app

FROM alpine as prod
WORKDIR /app
COPY --from=builder /app/app /app/app
COPY --from=builder /app/config.toml /app/config.toml
EXPOSE 8080
CMD ["./app"]