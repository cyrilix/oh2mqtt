FROM docker.io/golang:1.21-alpine AS builder

WORKDIR /go/src
ADD . .

RUN CGO_ENABLED=0 go build -mod vendor -tags netgo ./cmd/oh2mqtt/

#####################################
FROM gcr.io/distroless/static

LABEL org.opencontainers.artifact.description="Mqtt gateway for OpenHome"
LABEL org.opencontainers.image.source="https://github.com/cyrilix/oh2mqtt"
LABEL org.opencontainers.image.licenses=Apache-2.0

USER 1234
COPY --from=builder /go/src/oh2mqtt /go/bin/oh2mqtt
ENTRYPOINT ["/go/bin/oh2mqtt"]
