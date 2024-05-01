# Use the official Golang image to create a build artifact.
FROM golang:1.18 as builder

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY . ./

# Build the binary.
# -o specifies the output file
# Using CGO_ENABLED=0 to ensure a fully static binary
RUN CGO_ENABLED=0 GOOS=linux go build -v -o server ./main.go

# Use the official Alpine image for a lean production container.
FROM alpine:latest
WORKDIR /root/

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/server .
COPY config.toml ./

# Expose port 8080 to the Docker host, so we can access it
# from the outside.
EXPOSE 8080

# Run the web service on container startup.
CMD ["./server"]
