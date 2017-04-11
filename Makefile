default: docker_build
deps: ## Install dependencies
	@echo "$(INFO_COLOR)==> $(RESET)$(BOLD)Installing Dependencies$(RESET)"
	go get -u github.com/golang/dep/...
	dep ensure -v

build: deps  ## Build as linux binary
	@echo "$(INFO_COLOR)==> $(RESET)$(BOLD)Building$(RESET)"
	go build

docker_build: ## Build as linux binary
	@echo "$(INFO_COLOR)==> $(RESET)$(BOLD)Building$(RESET)"
	docker build -t prd .
	docker run --dns 8.8.8.8 --rm -v `pwd`:/go/src/github.com/pyama86/prd -w /go/src/github.com/pyama86/prd -t prd:latest
