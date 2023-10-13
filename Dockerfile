# alpine:3.18.4
FROM alpine@sha256:48d9183eb12a05c99bcc0bf44a003607b8e941e1d4f41f9ad12bdcc4b5672f86

RUN mkdir app

RUN ls -la

# Copy binary & static assets from build to main folder.
COPY hub /app
COPY web/views /app/web/views
COPY web/public /app/web/public

# Export necessary port.
WORKDIR /app

USER nonroot

# Command to run when starting the container.
ENTRYPOINT ["./hub"]
