FROM golang

ENV GO111MODULE=on

WORKDIR /app

COPY . .

RUN cd service && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
RUN ls
EXPOSE 8111

CMD ["./app/service/service"]