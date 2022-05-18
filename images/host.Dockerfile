FROM golang:1.18

WORKDIR /usr/src/host

COPY host/go.mod host/go.sum ./
RUN go mod download && go mod verify

COPY host .
RUN go build -v -o /usr/local/bin/host

WORKDIR /usr/src/http
COPY services/http .
RUN go build -v -o /usr/local/bin/http

WORKDIR /usr/src/echo
COPY services/echo .
RUN go build -v -o /usr/local/bin/echo

CMD ["sh", "/var/app/data/start.sh"]
