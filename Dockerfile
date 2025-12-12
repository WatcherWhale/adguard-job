FROM golang:1.25.4 AS build

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
