FROM golang:1.21.4-alpine
RUN apk add git
ENV CGO_ENABLED=0
WORKDIR /
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go install -ldflags "-X main.version=$(git describe --tags)"

FROM alpine:3.18.4
ENV ENVOY_DASHBOARD_LISTENADDR=0.0.0.0:8080
COPY --from=0 /go/bin/envoyproxy-dashboard /usr/local/bin/
ENTRYPOINT [ "envoyproxy-dashboard" ]

# for goreleaser
# FROM alpine:3.16.2
# ENV ENVOY_DASHBOARD_LISTENADDR=0.0.0.0:8080
# COPY envoyproxy-dashboard /usr/local/bin/
# ENTRYPOINT [ "envoyproxy-dashboard" ]
