# This version-strategy uses git tags to set the version string
VERSION := $(shell git describe --tags --always --dirty)
GIT_COMMIT := $(shell git rev-list -1 HEAD)

# HELP
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help

help:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

version: ## Show version
	@echo $(VERSION) \(git commit: $(GIT_COMMIT)\)

workflow-generate: ## Generate github workflows from templates
	cd .github && sh workflows.sh

# HTTPS TASK
key: ## [HTTP] Generate key
	openssl genrsa -out server.key 2048
	openssl ecparam -genkey -name secp384r1 -out server.key

cert: ## [HTTP] Generate self signed certificate
	openssl req -new -x509 -sha256 -key server.key -out server.pem -days 3650

dev: ## [HTTP] Run server in dev mode
	docker-compose up --build

scale-dev: ## [HTTP] Scale server
	docker-compose up --build --scale app=2 --scale server=2

scale-prod: ## [HTTP] Scale server
	docker-compose -f docker-compose.stage.yml up --build --scale placio-backend=3

test: ## [HTTP] Run tests
	go test -v -coverprofile cover.out ./cmd/app/...
	go test -coverprofile=coverage.out -covermode=atomic ./cmd/app/...
	go tool cover -html=coverage.out -o coverage.html
	#open coverage.html
	go tool cover -html=cover.out -o cover.htmltest:

test-pkg:## [HTTP] Run tests
	#go test -v -coverprofile test-cover.out ./pkg/...
	go test -coverprofile=coverage.out -covermode=atomic ./cmd/app/...
	go tool cover -html=coverage.out -o coverage.html
	#open coverage.html
	#go tool cover -html=cover.out -o cover.html

open: ## [HTTP] Run linters
	open cover.html

integration-test: ## [HTTP] Run integration tests
	chmod +x wait-for-server.sh
	chmod +x wait-for-it.sh
	chmod +x test.sh
	./test.sh

swagger: ## Generate docs
	swag init -g cmd/app/main.go --output docs/app --parseDependency --parseInternal
redoc: ##  Generate redoc
	redocly bundle docs/app/swagger.yaml --output docs/redoc.yml

swagger-build: swagger redoc ## Generate docs
	export NODE_OPTIONS="--max-old-space-size=16384"
	redocly build-docs docs/app/swagger.yaml --output docs/redoc-static.html

split-swagger: ## Split swagger
	redocly split docs/openapi.yml --outDir=placio

html-doc: ## Build docs
	redocly build-docs docs/openapi.yml --output redoc.html

docs: ## Generate documentation
	swag init -g cmd/app/main.go --output docs/app --parseDependency --parseInternal
	redocly bundle docs/app/swagger.yaml --output docs/redoc.yml

generate: ## Run Go generate
	export GOWORK=off
	go run entgo.io/ent/cmd/ent generate  --idtype string ./app/ent/schema

grpc: ## Run Go generate
	 protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./grpc/proto/*/*.proto


migrate: ## Run database migration. Pass migration name as a parameter, eg: make migrate migration_name=init_schema
	atlas migrate diff $(migration_name) \
  	--dir "file://ent/migrate/migrations" \
  	--to "file://./cmd/app/ent/schema" \
  	--dev-url "docker://postgres/15/test?search_path=public"


docs-preview: ## Generate documentation
	redocly preview-docs docs/redoc.yml

# DOCKER TASKS
docker-build: ## [DOCKER] Build given container. Example: `make docker-build BIN=user`
	docker build -f cmd/$(BIN)/Dockerfile --no-cache --build-arg BIN=$(BIN) --build-arg VERSION=$(VERSION) --build-arg GIT_COMMIT=$(GIT_COMMIT) -t tutis-api-$(BIN) .

docker-run: ## [DOCKER] Run container on given port. Example: `make docker-run BIN=user PORT=3000`
	docker run -i -t --rm -p=$(PORT):$(PORT) --name="tutis-api-$(BIN)" tutis-api-$(BIN)

docker-stop: ## [DOCKER] Stop docker container. Example: `make docker-stop BIN=user`
	docker stop tutis-api-$(BIN)

docker-rm: docker-stop ## [DOCKER] Stop and then remove docker container. Example: `make docker-rm BIN=user`
	docker rm tutis-api-$(BIN)

docker-publish: docker-tag-latest docker-tag-version docker-publish-latest docker-publish-version ## [DOCKER] Docker publish. Example: `make docker-publish BIN=user REGISTRY=https://your-registry.com`

docker-publish-latest:
	@echo 'publish latest to $(REGISTRY)'
	docker push $(REGISTRY)/tutis-api-$(BIN):latest

docker-publish-version:
	@echo 'publish $(VERSION) to $(REGISTRY)'
	docker push $(REGISTRY)/tutis-api-$(BIN):$(VERSION)

docker-tag: docker-tag-latest docker-tag-version ## [DOCKER] Tag current container. Example: `make docker-tag BIN=user REGISTRY=https://your-registry.com`

docker-tag-latest:
	@echo 'create tag latest'
	docker tag tutis-api-$(BIN) $(REGISTRY)/tutis-api-$(BIN):latest

docker-tag-version:
	@echo 'create tag $(VERSION)'
	docker tag tutis-api-$(BIN) $(REGISTRY)/tutis-api-$(BIN):$(VERSION)

docker-release: docker-build docker-publish ## [DOCKER] Docker release - build, tag and push the container. Example: `make docker-release BIN=user REGISTRY=https://your-registry.com`

# TERRAFORM TASKS
terraform-install: ## [TERRAFORM] Install terraform deployment to your kubernetes cluster. Example: `make terraform-install`
	cd k8s \
  kubectl apply --validate=false -f https://github.com/jetstack/cert-manager/releases/download/v1.3.1/cert-manager.crds.yaml \
	terraform init \
	terraform apply \
  terraform output -raw templates | kubectl apply -f -

terraform-destroy: ## [TERRAFORM] Remove deployment kubernetes cluster. Example: `make terraform-destroy`
	cd k8s \
  terraform destroy

# TELEPRESENCE TASKS
telepresence-swap-local: ## [TELEPRESENCE] Replace the existing deployment with the Telepresence proxy for local process. Example: `make telepresence-swap-local BIN=user PORT=3000 DEPLOYMENT=tutis-api-user`
	go build -o cmd/$(BIN)/$(BIN) cmd/$(BIN)/main.go
	telepresence \
	--swap-deployment $(DEPLOYMENT) \
	--expose 3000 \
	--run ./cmd/$(BIN)/$(BIN) \
	--port=$(PORT) \
	--method vpn-tcp

telepresence-swap-docker: ## [TELEPRESENCE] Replace the existing deployment with the Telepresence proxy for local docker image. Example: `make telepresence-swap-docker BIN=user PORT=3000 DEPLOYMENT=tutis-api-user`
	telepresence \
	--swap-deployment $(DEPLOYMENT) \
	--docker-run -i -t --rm -p=$(PORT):$(PORT) --name="$(BIN)" $(BIN):latest

# HELPERS
# generate script to login to aws docker repo
CMD_REPOLOGIN := "eval $$\( aws ecr"
ifdef AWS_CLI_PROFILE
CMD_REPOLOGIN += " --profile $(AWS_CLI_PROFILE)"
endif
ifdef AWS_CLI_REGION
CMD_REPOLOGIN += " --region $(AWS_CLI_REGION)"
endif
CMD_REPOLOGIN += " get-login \)"

aws-repo-login: ## [HELPER] login to AWS-ECR
	@eval $(CMD_REPOLOGIN)
