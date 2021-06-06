FROM golang:alpine

COPY .env_example .env

# Copy dir to workdir
COPY . /go/src/github.com/FelipeAz/golibcontrol

# Install git & bash
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /go/src/github.com/FelipeAz/golibcontrol/cmd/management/

# Copy go mod and sum files
COPY go.mod go.sum /go/src/github.com/FelipeAz/golibcontrol/

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

## Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

#RUN go run main.go

EXPOSE 8081

CMD ["/go/src/github.com/FelipeAz/golibcontrol/cmd/management/main"]