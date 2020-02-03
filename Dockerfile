FROM golang:alpine

WORKDIR /opt/excel-to-markdown
COPY . /opt/excel-to-markdown

RUN go get && go build

FROM alpine:latest

WORKDIR /root

RUN mkdir -p /usr/local/bin

COPY --from=0 /opt/excel-to-markdown/excel-to-markdown /usr/local/bin

ENTRYPOINT ["/usr/local/bin/excel-to-markdown"] 
