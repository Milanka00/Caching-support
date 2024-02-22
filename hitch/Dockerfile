# Use a base image with Hitch pre-installed (e.g., Alpine Linux with Hitch package)
FROM alpine:latest

# Install Hitch package
RUN apk --no-cache add hitch

# Copy Hitch configuration file
COPY hitch.conf /etc/hitch/hitch.conf

# Expose the default Hitch port
EXPOSE 443

# Run Hitch with the provided configuration
CMD ["hitch", "--config", "/etc/hitch/hitch.conf"]
