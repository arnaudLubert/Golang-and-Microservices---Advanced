FROM golang:1.17

WORKDIR /go/src/transactions

ENV GO111MODULE=off

COPY . .
RUN ls -a

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["transactions"]

 