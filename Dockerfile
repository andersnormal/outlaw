# BUILD
FROM golang:latest as build

WORKDIR /go/github.com/andersnormal/outlaw
COPY . .

RUN echo $GOPATH

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o outlaw -v -ldflags "-extldflags '-static'" -a -installsuffix cgo main.go

# RUNTIME
FROM scratch

LABEL maintainer="sebastian@andersnormal.us"

COPY --from=build /go/github.com/andersnormal/outlaw/outlaw /outlaw
ENTRYPOINT [  "/outlaw" ]