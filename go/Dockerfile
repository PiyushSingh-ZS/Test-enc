FROM golang:1.20
ARG PAT

WORKDIR /usr/src/app

RUN go env -w GOPRIVATE=developer.zopsmart.com/go,github.com/zopsmart
RUN apt-get -y update
RUN apt-get -y install git
RUN git config --global url."https://${PAT}:x-oauth-basic@github.com/".insteadOf "https://github.com/"

RUN git config --global http.sslVerify false

COPY go.mod ./
RUN go get developer.zopsmart.com/go/gofr

RUN rm go.mod