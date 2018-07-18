FROM golang

ADD . /go/src/github.com/tushar9989/github-search

RUN go get github.com/PuerkitoBio/goquery
RUN go test github.com/tushar9989/github-search/internal/repository
RUN go test github.com/tushar9989/github-search/internal/service
RUN go install github.com/tushar9989/github-search/cmd

ENTRYPOINT /go/bin/cmd