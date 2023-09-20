FROM golang:1.19 as builder

WORKDIR /workspace
COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY main.go main.go

RUN CGO_ENABLED=0 GOOS=linux GOARCH=$TARGETARCH go build -a -o kvm-device-plugin main.go

FROM alpine:3.18
WORKDIR /
COPY --from=builder /workspace/kvm-device-plugin .
RUN mkdir -p /var/lib/kubelet/device-plugins/

ENTRYPOINT ["/kvm-device-plugin"]