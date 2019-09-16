# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Ekta Garg <gargekta65@gmail.com>"

#Create a directory
RUN mkdir /app 

#Add the entire code to newly created directory
ADD . /app/ 

# Set the Current Working Directory inside the container
WORKDIR /app

# Build the Go app
RUN go build -o main .

# Command to run the executable
CMD ["/app/main"]

