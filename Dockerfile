FROM golang:latest

ENV GOPATH=/go
RUN go get github.com/go-task/task/cmd/task
ENV PATH=$PATH:/go/bin

WORKDIR /go/src/github.com/daregod/amazon-co-uk-scraper
COPY ./ /go/src/github.com/daregod/amazon-co-uk-scraper
RUN task update-deps test build
CMD scrape-server
EXPOSE 8007
