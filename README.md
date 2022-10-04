# Grab It Yourself Backend

## Prerequisites

1. Go 1.19

## Setup

1. Clone the repository
2. Run `go mod download` to download the dependencies (unless you are using GoLand, then just click sync dependencies in the imports)
3. Copy `config.yml.example` to `config.yml` in `./<service>/config`and fill in the values.
4. Run `go run <service>/cmd/main.go` to run the service.

## Project Structure

Each directory in the root folder is an individual microservice (api, user, shop, order, auth, payment, notification).
Each microservice has the following structure:

- cmd
  - main.go (service entrypoint)
- config
  - config.go (for accessing config variables)
  - config.yml (for setting config variables)
- pkg
  - repository (for accessing the databases)
  - other_pkgs (types, errors, logic, handlers, util, etc)
- Dockerfile

## Creating a new service

1. Create a new directory in the root folder
2. Follow the structure and code conventions of the other services (mostly config stuffs)

## Migration

Create a new migration:

```bash
migrate create -ext sql -dir ./migrations -seq <migration_name>
```

Up migrations:

```bash
migrate -path ./migrations -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" up
```

down migrations:

```bash
migrate -path ./migrations -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" down
```

## Protobuf

1. Install protoc compiler binary. See [here](https://grpc.io/docs/protoc-installation/) for more details.
2. Install Go plugin for protoc. See [here](https://grpc.io/docs/languages/go/quickstart/) for more details.

   ```bash
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
   ```

3. Generate Protobuf file in each service

   ```bash
   sh script/create_proto.sh <service-name>
   ```

4. Once you finish editing the proto file, run the following command to generate the Go code

   ```bash
   sh script/generate_proto.sh <service-name>
   ```
