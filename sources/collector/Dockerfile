#### BUILD STAGE ####################################
FROM golang:1.24.5-bookworm AS builder

WORKDIR /app

# Setup dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy sources and build application
COPY . .
RUN CGO_ENABLED=0 go build -o main .

#### FINAL STAGE ####################################
FROM gcr.io/distroless/static-debian12

COPY --from=builder /app/main /

# Setup non-root user
USER nonroot:nonroot

CMD ["/main"]
