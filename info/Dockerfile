### Builder
FROM golang:1.17-alpine3.13 as builder
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build
COPY . .

RUN apk update && apk add git
RUN go build -o main .

FROM alpine:3.13
COPY --from=builder /build /

ENTRYPOINT [ "/main" ]