from golang:1.25.3-trixie as build

workdir /go/src/
add . /go/src/
run go clean
run go get
run go install
run go install github.com/swaggo/swag/cmd/swag@v1.16.4
run swag init

# Build
run CGO_ENABLED=0 GOOS=linux go build -o /go/bin/ktfs

from alpine:latest as prod
run apk --no-cache add tzdata
workdir /go/src
copy --from=build go/bin/ktfs /go/bin/ktfs
copy --from=build /go/src/.env /go/src/.env

expose 8080

# Run
entrypoint [ "/go/bin/ktfs" ]