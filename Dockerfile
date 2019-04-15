FROM golang:1.12.1-stretch

# Lable for image
LABEL maintainer="dinhluong"

RUN mkdir -p  $GOPATH/src/github.com/dinhluong/postgresql-golang
# Set work dir.
WORKDIR $GOPATH/src/github.com/dinhluong/postgresql-golang

COPY . .
# Go get
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

RUN ["bin/sh", "env.sh"]

ENTRYPOINT [ "entrypoint.sh" ]

EXPOSE 8000