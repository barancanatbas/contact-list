FROM golang:1.16.5 as development
# Add a work directory inside container
WORKDIR /app
# Cache and install dependencies
COPY go.mod go.sum ./
RUN go mod download
# Copy the source from the current directory to the working Directory inside the container
COPY . .
# Install Reflex for development
#RUN go install github.com/cespare/reflex@latest
# Expose port
EXPOSE 8080
# Start app
CMD  go run main.go