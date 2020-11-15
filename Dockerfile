# latest golang base image
FROM golang:latest

# Maintainer Info
LABEL maintainer="Utsav <utsav.r.vaghani@gmail.com>"

# Current Working Directory inside the container
WORKDIR /app

# Copy go mod file
COPY go.mod .

# Copy go sum file
COPY go.sum .

# Download all dependencies
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 8000 to the outside world
EXPOSE 8000

# Command to run the executable
CMD ["./main"]

