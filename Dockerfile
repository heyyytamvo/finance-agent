# ---------- Build stage ----------
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy go.mod first (better caching)
COPY go.mod ./
RUN go mod download

# Copy source
COPY . .

# Build static binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o app

# ---------- Runtime stage ----------
FROM gcr.io/distroless/base-debian12

WORKDIR /app

COPY --from=builder /app/app .

CMD ["./app"]
