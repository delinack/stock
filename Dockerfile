FROM golang:latest AS builder

### Build stage ###
WORKDIR /gopath_dir/src/stock
ENV GOPATH /gopath_dir
ENV GOBIN $GOPATH/bin

COPY go.sum $GOPATH/src/stock
COPY go.mod $GOPATH/src/stock
RUN go mod download

COPY . $GOPATH/src/stock

RUN CGO_ENABLED=0 go build -o /app ./cmd/app


### Deploy stage ###
FROM redhat/ubi8:8.7-1112

COPY --from=builder /app /app

CMD ["/app"]
