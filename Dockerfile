FROM golang:1.12 as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -v -o proxy

FROM alpine
RUN apk add --no-cache ca-certificates
COPY --from=builder /app/proxy /app/proxy
CMD ["/app/proxy"]
