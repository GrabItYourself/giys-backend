# Grab It Yourself Backend

## Prerequisites
1. Go 1.19

## Setup
1. Clone the repository
2. Run `go mod download` to download the dependencies (unless you are using GoLand, then just click sync dependencies in the imports)
3. Copy `config.yml.example` to `config.yml` in `./<service>/config`and fill in the values.
4. Run `go run <service>/cmd/main.go` to run the service.