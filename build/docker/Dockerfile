FROM golang:1.22 

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    PATH=/go/bin:/usr/local/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin \
    GOPATH=/go \
    PWD=/hello-privilee
    
RUN apt-get install -y

WORKDIR /hello-privilee

COPY . /hello-privilee

RUN chmod +x zone.sh

RUN go mod init hello-privilee

RUN go mod tidy && \
    go build

EXPOSE 1323

# Run
# CMD ["go", "run", "hello-privilee"]