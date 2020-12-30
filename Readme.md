# Local setup
* Install version `go1.15.x `
* Run `local-setup` to setup hooks and other tools
* Install golangci-lint threw brew `brew install golangci/tap/golangci-lint` with version v1.21..0

# Commands for local development
* Run app :`make run`
* All tests : `make test`
* Build binary : `make build`
* Coverage : `make coverage`
* Lint : `make lint`

# Folder Structure
- **internal:** contains application code
  
# Run Application
* go run main.go

# Curl Request

Find best match

curl --location --request POST 'localhost:8080/api/v1/rhymes/match-word' \
--header 'Content-Type: application/json' \
--data-raw '{
"words": ["Shooting","Disputing"]
}'

To Get reference list:-

curl --location --request POST 'localhost:8080/api/v1/rhymes/reference-list' \
--header 'Content-Type: application/json' \
--data-raw '{}  '

#Logging
* Logrus for logging
* Context for timeouts
