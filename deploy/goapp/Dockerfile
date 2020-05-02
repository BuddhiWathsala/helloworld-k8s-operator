FROM golang:1.12-alpine

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/BuddhiWathsala/helloworld/

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the hello-world app
CMD ["go", "run", "hello-world.go"]