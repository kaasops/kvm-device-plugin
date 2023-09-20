FROM golang:1.19 as builder

WORKDIR /workspace
COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY main.go main.go

RUN CGO_ENABLED=0 GOOS=linux GOARCH=$TARGETARCH go build -a -o kvm-device-plugin main.go

FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/kvm-device-plugin .
USER 65532:65532

ENTRYPOINT ["/kvm-device-plugin"]