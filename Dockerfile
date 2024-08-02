FROM golang:1.22.5@sha256:829eff99a4b2abffe68f6a3847337bf6455d69d17e49ec1a97dac78834754bd6 AS build

WORKDIR /build

COPY go.mod ./
RUN go mod download

COPY cmd ./cmd
COPY pkg ./pkg

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w" -v -o ./bin/adguard-job ./cmd

FROM scratch
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build /build/bin/adguard-job /adguard-job

ENTRYPOINT ["/adguard-job"]
