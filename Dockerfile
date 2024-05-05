# alpine:3.19.1
FROM --platform=$BUILDPLATFORM alpine@sha256:c5b1261d6d3e43071626931fc004f70149baeba2c8ec672bd4f27761f8e1ad6b
RUN mkdir app

# Copy binary
COPY hub /app

WORKDIR /app

ENTRYPOINT ["./hub"]
