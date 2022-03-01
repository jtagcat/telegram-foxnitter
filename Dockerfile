
# 1.17 hardcoded; stuck on https://github.com/moby/moby/issues/29110
# moved to https://gist.github.com/jtagcat/189b2fd239687ab700f54faa46907df4

FROM golang:1.17 AS builder
WORKDIR /wd

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds

FROM alpine
LABEL org.opencontainers.image.source="https://github.com/jtagcat/telegram-foxnitter"
WORKDIR /wd
#RUN apk --no-cache add ca-certificates
COPY --from=builder /wd/app ./
CMD ["./app"]
