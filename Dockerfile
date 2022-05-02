# builder image
FROM golang:1.18-alpine as builder
RUN apk add git
RUN mkdir /build
ADD . /build
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -o app ./cmd/go-getir

# final image
FROM alpine:3.15
COPY --from=builder /build/app .
COPY --from=builder /build/config.yml .

ENTRYPOINT [ "./app" ]