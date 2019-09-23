FROM golang:1.12 AS builder

COPY . /build/
WORKDIR /build/service
RUN go build

FROM alpine:3.10
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
WORKDIR /emqxample/
COPY --from=builder /build/service ./

EXPOSE 8111
ENTRYPOINT sh -c "/emqxample/service"