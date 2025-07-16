#### BUILD STAGE ####################################
FROM node:24.4-bookworm-slim AS builder

WORKDIR /app

# Setup dependencies
COPY package.json package-lock.json ./
RUN npm ci

# Copy sources and build application
COPY . .
RUN npm run build

#### DOWNLOAD STAGE #################################
FROM node:24.4-bookworm-slim AS downloader

WORKDIR /app

COPY package.json package-lock.json ./
RUN npm ci --omit=dev

#### FINAL STAGE ####################################
FROM gcr.io/distroless/nodejs24-debian12

WORKDIR /app
COPY --from=builder /app/dist ./dist
COPY --from=downloader /app/node_modules ./node_modules

# Setup non-root user
USER nonroot:nonroot

CMD ["./dist/index.js"]
