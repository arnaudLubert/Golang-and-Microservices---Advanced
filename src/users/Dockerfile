FROM golang:1.17

WORKDIR /go/src/users

ENV GO111MODULE=auto
RUN go env -w GOPROXY=https://goproxy.io,direct
COPY . .
RUN ls -a

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["users"]
