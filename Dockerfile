# Build stage - Security hardened Alpine with latest Go
FROM golang:1.25.5-alpine3.22 AS builder

# Security: Install only essential dependencies and clean up
RUN apk add --no-cache --update \
    bash \
    curl \
    jq \
    git \
    ca-certificates && \
    apk upgrade --no-cache && \
    rm -rf /var/cache/apk/* /tmp/*

# Set working directory and create secure cache directories
WORKDIR /build
RUN mkdir -p /.cache && \
    chmod 700 /.cache && \
    chown root:root /.cache

# Arguments for build configuration
ARG SERVICE=my-finance-app
ARG PORT=8080
ARG SWAGGER_HOST=0.0.0.0:8080
ENV SERVICE=${SERVICE}
ENV PORT=${PORT}
ENV SWAGGER_HOST=${SWAGGER_HOST}

# Copy the entire project
COPY . .

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o my-finance-app

## Generate Swagger documentation
#RUN go run tools/main.go swagger-generate staging && \
#    chmod -R 555 ./docs && \
#    echo "Swagger generated for staging environment" && \
#    ls -la docs/

# Try to fix the module issue by recreating it
RUN go mod tidy && go mod verify

# Production stage - Latest Alpine with security updates
FROM alpine:3.22

# Security: Install only essential packages with specific versions and clean up
RUN apk add --no-cache --update \
    ca-certificates \
    tzdata && \
    apk upgrade --no-cache && \
    rm -rf /var/cache/apk/* /tmp/* /var/tmp/* && \
    addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup -s /sbin/nologin -h /app

# Create app directory with secure permissions
WORKDIR /app
RUN chown appuser:appgroup /app && \
    chmod 750 /app && \
    mkdir -p /tmp/app-cache && \
    chown appuser:appgroup /tmp/app-cache && \
    chmod 700 /tmp/app-cache && \
    ln -s /tmp/app-cache /app/.cache

# Copy the binary from builder stage with secure permissions
COPY --from=builder --chown=root:root --chmod=555 /build/my-finance-app .

# Copy generated docs with secure permissions
#COPY --from=builder --chown=root:root --chmod=555 /build/docs ./docs

# Copy translation files with secure permissions
#COPY --from=builder --chown=root:root --chmod=555 /build/translations ./translations

# Switch to non-root user
USER appuser

# Set secure environment variables
ENV GIN_MODE=release \
    HOME=/app \
    GOCACHE=/tmp/app-cache/gocache \
    GOMODCACHE=/tmp/app-cache/gomodcache \
    XDG_CACHE_HOME=/tmp/app-cache \
    CGO_ENABLED=0 \
    GODEBUG=netdns=go

# Set PORT environment variable
ARG PORT=8080
ENV PORT=${PORT}

# Expose port
EXPOSE ${PORT}

# Command to run
CMD ["./my-finance-app"]
