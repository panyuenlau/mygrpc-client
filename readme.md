# Prerequisite
- Install docker for your operating system: https://docs.docker.com/get-docker/
-  Install Golang and protobuf:
```
brew install go
brew install protobuf
```

# Setup environment variables
```
export GOROOT=/usr/local/opt/go/libexec
export GOPATH=$HOME/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOROOT:$GOPATH:$GOBIN
```

# gRPC
## Install the protocol compiler plugin
```
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.0
```

## Generate grpc client code from .proto service definition.
At the project directory, run:
```
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/service.proto
```


# Docker
## Build docker image
At the project directory, run:
```
docker build -t [image-name] .
```

Check to see if the image was built successfully
```
docker images
```

## Run docker image
Run the docker image in interactive mode:
```
docker run --net-host -it [image-name]
```

The simple grpc client is now sending request to the server at `localhost:50051` and waiting for response!