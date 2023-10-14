# alpine:3.18.4
FROM alpine@sha256:48d9183eb12a05c99bcc0bf44a003607b8e941e1d4f41f9ad12bdcc4b5672f86

RUN mkdir app

RUN ls -la

# Copy binary
COPY hub /app

WORKDIR /app

ENTRYPOINT ["./hub"]
