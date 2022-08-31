FROM alpine:3.16.2
ENV ENVOY_DASHBOARD_LISTENADDR=0.0.0.0:8080
COPY envoyproxy-dashboard /usr/local/bin/
ENTRYPOINT [ "envoyproxy-dashboard" ]
