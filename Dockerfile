# Fetch
FROM golang:1.22 AS fetch-stage
COPY go.mod go.sum /app/
WORKDIR /app
RUN go mod download

# Generate
FROM ghcr.io/a-h/templ:v0.2.793 AS generate-stage
COPY --chown=65532:65532 . /app
WORKDIR /app
RUN ["templ", "generate"]

# Build
FROM golang:1.22 AS build-stage
# Add NodeJS repository and install Node.js
RUN curl -fsSL https://deb.nodesource.com/setup_20.x | bash - \
    && apt-get install -y nodejs

# Make sure we're using the same templ version in Go code
COPY go.mod go.sum /app/
WORKDIR /app
RUN go mod download
RUN go get github.com/a-h/templ@v0.2.793

# Copy generated code and other files
COPY --from=generate-stage /app /app
WORKDIR /app

# Create static directory before build
RUN mkdir -p /app/static

# Run the build
RUN make build

# Test
FROM build-stage AS test-stage
RUN go test -v ./...

# Deploy
FROM gcr.io/distroless/base-debian12 AS deploy-stage
WORKDIR /app
COPY --from=build-stage /app/bin/main /app/main

# Instead of trying to copy static, just create an empty directory
# This ensures the application has a static directory to work with
WORKDIR /app
# Create an empty static directory in the distroless image
# This is a hack to make sure the directory exists
COPY --from=build-stage /app/static/. /app/static/
# Ensure .env file exists to avoid runtime errors
RUN touch .env 

EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/app/main"]