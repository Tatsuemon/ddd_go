FROM golang:1.15.0
ENV GO111MODULE=on
WORKDIR /ddd_go
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

RUN go get github.com/pilu/fresh && \
    go get -v github.com/rubenv/sql-migrate/...