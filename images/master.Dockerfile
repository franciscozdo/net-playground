FROM golang:1.18

WORKDIR /usr/src/master

COPY master/go.mod master/go.sum ./
RUN go mod download && go mod verify

COPY master .
RUN go build -v -o /usr/local/bin/master

CMD ["master", "-c", "/var/app/data/config.yaml"]
