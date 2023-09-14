# Use an official Golang runtime as the base image
FROM golang:1.18-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to the container
COPY go.mod .
COPY go.sum .

# Download and cache Go modules
RUN go mod download

# Copy the core and taskqueue directories to the container
COPY ./core ./core
COPY ./services/taskqueue ./services/taskqueue

# Set environment variables
ENV PORT_TASKQUEUE_SERVICE=:8080
ENV DB_HOST=yourhost
ENV DB_PORT=yourport
ENV DB_USER=youruser
ENV DB_PASSWORD=yourpassowrd
ENV DB_DBNAME_SMSQUEUE=taskqueue
ENV JWT_SECRET_KEY=gokitIsAmazingTech

# Build the taskqueue service
RUN go build -o taskqueue ./services/taskqueue/main.go

# Expose the port that the service will run on
EXPOSE 8080:8080

# Command to run the taskqueue service
CMD ["./taskqueue"]
