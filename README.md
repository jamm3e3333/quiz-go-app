# Quiz Application
This is a CLI-based quiz application that communicates with a gRPC server. The application allows users to answer quiz questions and receive feedback on their performance compared to others.

## Features
- CLI tool to take quizzes
- gRPC-based server to handle quiz submissions
- Questions and answers stored in memory
- Performance comparison with other users
- Built using Go, gRPC, and Protocol Buffers
- Project Structure


## Prerequisites
- Go 1.22+ installed
- Docker and Docker Compose installed
- Make for managing commands

## Setup and Run Locally
You can run the application using Docker or manually using Go.

### 1. Run the Application with Docker
   The application is dockerized and run using docker-compose defined: [docker-compose.yml](./docker-compose.yaml).

Build the application and start the gRPC server listening by default on port 8088:

```makefile
make up
```

This will:
- Build the cli app and build the server and start it inside a Docker container.

Stop the application:
```makefile
make down
```

### 2. Running the gRPC Server Manually
   You can run the gRPC server manually using Go executable.

Start the gRPC server:
```bash
go run cmd/start_server/.
```

This will start the gRPC server, which will by default listen on port 8088.

### 3. Build the binaries (CLI Tool & gRPC Server)
   The CLI tool allows you to take quizzes by communicating with the gRPC server. Both CLI tool and gRPC server can be built using a provided script and are located in the [target/cli](./target/cli), [target/server](./target/server) folders after building.
```bash
script/build.sh 
```

### 4. Run the CLI
After building the CLI, you can run it as follows:

```bash
./target/cli quiz
```

The CLI will interact with the running gRPC server, allowing you to take a quiz.

### gRPC Protocol Buffers
The communication between the CLI and the server is done using gRPC and Protocol Buffers. The .proto file defining the communication is located in [quiz-proto](./grpc/protobuff/quiz.proto).

If you need to regenerate the gRPC code (e.g., after modifying the .proto file):

Run the protobuf generation script:


```makefile
make generate-proto
```

This will regenerate the Go code for the gRPC service.

### 4. Run the Tests
   You can run the tests using the provided script:

```makefile 
make test
```

This will run the tests within in the docker container. For running the tests locally, you can run the following commands with e.g.:

```bash
go test -race -v -count=1 -timeout 50s -coverpkg=./... -coverprofile=./tmp/coverage ./...
```