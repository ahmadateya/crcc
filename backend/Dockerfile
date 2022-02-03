FROM golang:1.17-alpine

# Setup the Go app
RUN mkdir /backend
COPY . /backend
WORKDIR /backend

RUN apk add --update --no-cache openssl bash alpine-sdk autoconf vim gd libzip-dev zip unzip
RUN go get github.com/githubnemo/CompileDaemon

# Expose the port the app will run on
EXPOSE 8080

# Add VOLUMEs to allow backup of config, logs and databases
#HEALTHCHECK CMD ["wget", "-q", "0.0.0.0:8080"]

# Set the default command to run when starting the container
