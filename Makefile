BIN_DIR := $(GOPATH)/bin
GOMETALINTER := $(BIN_DIR)/gometalinter
PKGS := $(shell go list ./... | grep -v /vendor)
run-all-local:
	docker-compose up --build

# runs both unit and integration tests
.PHONY: test
test:
	docker-compose -f docker-compose.test.yaml up --build --abort-on-container-exit --remove-orphans --force-recreate
	docker-compose -f docker-compose.test.yaml down --volumes

run-unit-test:
	go test ./... --tags=unit

coverage:
	go test $(PKGS) -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage_report.html

$(GOMETALINTER):
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install &> /dev/null

.PHONY: lint
lint: $(GOMETALINTER)
	gometalinter ./... --vendor --deadline=60s