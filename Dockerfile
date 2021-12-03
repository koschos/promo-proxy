FROM golang:1.17 as builder
RUN update-ca-certificates
WORKDIR /app
COPY ./ ./
ARG version=dev
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags "-X main.version=$version" -o app ./main.go

FROM alpine:3.11
USER nobody
WORKDIR /app
COPY --from=builder /app .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
CMD ["./app"]
