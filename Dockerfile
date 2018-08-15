FROM golang:1.10.3

WORKDIR /go/src/github.com/Legenddigital/lddld
COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure
RUN go install . ./cmd/...

EXPOSE 9108

CMD lddld
