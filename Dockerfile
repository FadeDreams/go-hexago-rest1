# Use the official Go image as the base image.
FROM golang:1.21.3

# Set the working directory inside the container.
WORKDIR /usr/src/goapp/

# Copy your application code and go.mod file into the container.
ADD ./go.mod /usr/src/goapp/
ADD . /usr/src/goapp/

# Set the necessary environment variables.
ENV GO111MODULE="on" \
    CGO_ENABLED="1" \
    GO_GC="off"

# Update, upgrade, and install required packages.
RUN apt-get update --fix-missing && apt-get upgrade -y && apt-get install -y curl build-essential

# Download dependencies, tidy, verify, and build the application.
RUN go mod download && go mod tidy && go mod verify && go build -o main .

# Expose the port your application will listen on (replace 8080 with the actual port).
EXPOSE 8080

# Define the command to run your application.
CMD ["./main"]

