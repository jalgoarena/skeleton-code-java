FROM golang:1.11 as builder

ENV GOBIN /go/bin

RUN mkdir /app
RUN mkdir /go/src/app
ADD . /go/src/app
WORKDIR /go/src/app

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN dep ensure

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /app/main .

FROM alpine

COPY --from=builder /app/main /app/main

EXPOSE 8081
ENV PROBLEMS_URL "http://localhost:8080"

ENTRYPOINT /app/main -problems-url $PROBLEMS_URL
