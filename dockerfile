# Start from golang base image
FROM golang:alpine

# Add Maintainer info
LABEL maintainer="Matheus Vasconcelos <matheus.vasconcelos.dev@gmail.com>"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git && apk add --no-cache bash && apk add build-base

# Setup folders
RUN mkdir /app
WORKDIR /app

# Copy the source from the current directory to the working Directory inside the container
COPY . .
COPY .env .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

#Setup hot-reload for dev stage
RUN go install -mod=mod github.com/githubnemo/CompileDaemon
RUN go get -v golang.org/x/tools/gopls
RUN go mod tidy

ENTRYPOINT CompileDaemon --build="go build -a -installsuffix cgo -o main ." --command=./main