# Use a container with Go pre-installed
FROM quay.io/projectquay/golang:1.17

# Copy our source file into the container
COPY src/hello-world.go /go/hello-world.go

# Set the default environment variables
ENV MESSAGE "Welcome! You can change this message by editing the MESSAGE environment variable."
ENV HOME /go

# Set permissions to the /go folder (for OpenShift)
RUN chgrp -R 0 /go && chmod -R g+rwX /go

# Just documentation.
# This container needs Docker or OpenShift to help with networking
EXPOSE 8080

# OpenShift picks up this label and creates a service
LABEL io.openshift.expose-services 8080/http

# OpenShift uses root group instead of root user
USER 1001

# Command to run when container starts up
CMD go run hello-world.go
