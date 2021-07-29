fmt: ## gofmt and goimports all go files
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done

dep: ## Get all dependencies
	go env -w GOPROXY=direct
	go env -w GOSUMDB=off
	go mod download
	go mod tidy
	go mod vendor

build: dep ## Build a beta version
	go build -mod vendor -race -o golang-web-server .

start-service: ##Run the Service
	go run main.go

stop-service: ##Terminate the Service
	kill $(lsof -ti:8080,8082)

# Absolutely awesome: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help