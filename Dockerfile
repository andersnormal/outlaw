# BUILD
#FROM golang:latest as build
FROM golang:latest

LABEL maintainer="sejamich@googlemail.com"

WORKDIR /go/github.com/andersnormal/outlaw
COPY . .

RUN echo $GOPATH
RUN go get -d -v ./...

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o outlaw -v -ldflags "-extldflags '-static'" -a -installsuffix cgo main.go

# RUNTIME
#FROM scratch

#MAINTAINER Jan Michalowsky <sejamich@googlemail.com>

#COPY --from=build /go/github.com/andersnormal/outlaw/outlaw /outlaw
RUN cp /go/github.com/andersnormal/outlaw/outlaw /outlaw
CMD [ "/outlaw" ]