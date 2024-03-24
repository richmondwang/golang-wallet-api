#
#
# build
FROM golang:1.22-alpine AS build

RUN adduser -u 1000 -D godev
USER 1000
ENV GOOS=linux GOARCH=amd64 CGO_ENABLED=0
ENV PATH="$PATH:/go/bin/${GOOS}_${GOARCH}"

RUN mkdir /home/godev/build &&\
    chown godev:godev -R /home/godev/build
WORKDIR /home/godev/build

COPY --chown=godev:godev . /home/godev/build/
RUN go mod download
RUN go test ./pkg/...

#
#
# Migrate Build
FROM build as dbmigrate
RUN go build -o main cmd/dbmigrate/main.go
ENTRYPOINT ["./main"]


#
#
# Dev Build
FROM build as dev
RUN go build -o main cmd/api/main.go

RUN go install github.com/cosmtrek/air@latest
CMD ["air", "--build.cmd", "go build -o main cmd/api/main.go", "--build.bin", "./main", "--build.exclude_dir", "mocks", "--build.exclude_regex", "_test.go"]

#
#
# Dist build
FROM alpine:3

RUN apk --no-cache add ca-certificates

RUN adduser -u 1000 -D app
USER 1000

COPY --chown=app:app --from=dev /home/godev/build/main /app/

WORKDIR /app
ENTRYPOINT ["./main"]
