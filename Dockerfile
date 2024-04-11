FROM golang:alpine AS build-env

WORKDIR /go/src/lizisky.com/lizisky

COPY . .

ENV GOPROXY=https://goproxy.cn
RUN apk update
RUN apk add --no-cache git make bash libc-dev linux-headers gcc
RUN make build

FROM alpine

WORKDIR /root
WORKDIR /data

COPY --from=build-env /go/src/lizisky.com/lizisky/bin/liziskyd /usr/bin/liziskyd

EXPOSE 8081

CMD ["liziskyd"]
