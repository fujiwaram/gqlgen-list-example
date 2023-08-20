up: ## start
	docker-compose up

down: ## stop
	docker-compose down

clean_data: ## remove data
	rm -rf ./data

gqlgen: ## generate gqlgen
	go install github.com/99designs/gqlgen@v0.17.36
	go run -mod=mod github.com/99designs/gqlgen

help: ## Show help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help
