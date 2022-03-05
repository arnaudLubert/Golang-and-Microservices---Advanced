FROM golang:1.17

WORKDIR /go/src/authentication

ENV GO111MODULE=auto
RUN go env -w GOPROXY=https://goproxy.io,direct

COPY . .
RUN ls -a

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["authentication"]