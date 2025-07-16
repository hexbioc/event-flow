#### BUILD STAGE ####################################
FROM golang:1.24.5-bookworm AS builder

WORKDIR /app

# Setup dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy sources
COPY . .

# Run tests and build application
RUN go test ./... && \
    CGO_ENABLED=0 go build -o server .

#### FINAL STAGE ####################################
FROM gcr.io/distroless/static-debian12

COPY --from=builder /app/server /

# Setup non-root user
USER nonroot:nonroot

CMD ["/server"]
