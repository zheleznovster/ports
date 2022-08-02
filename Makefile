# Colors
BLUE_COLOR    = "\033[0;34m"
DEFAULT_COLOR = "\033[m"
LINTER = golangci-lint

BASEPATH = $(shell pwd)
LOCALBIN = $(BASEPATH)/bin
PATH := $(LOCALBIN):$(PATH)
SHELL := env PATH=$(PATH) /bin/bash

INSTALLLINTER		=   $(GOINSTALL) github.com/golangci/golangci-lint/cmd/golangci-lint@$(VLINTER)
INSTALLLINTERBIN	=   curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(LOCALBIN) $(VLINTER)

VCOUBERTURA        = "v1.2.0"
VLINTER            = "v1.44.2"
VGOTESTSUM         = "v1.7.0"

GOOSARCH    = GOOS=linux GOARCH=amd64
GOCMD       = go
GOBUILD     = $(GOOSARCH) $(GOCMD) build
GOINSTALL   = GOPATH=$(GOPATH) GOBIN=$(LOCALBIN) $(GOCMD) install
GOTEST      = gotestsum

ISTALLTEST      	=   $(GOINSTALL) github.com/boumenot/gocover-cobertura@$(VCOUBERTURA) && \
						$(GOINSTALL) gotest.tools/gotestsum@$(VGOTESTSUM)

install-linter:
	@echo -e $(BLUE_COLOR)[install-linter]$(DEFAULT_COLOR)
ifeq (,$(wildcard $(LOCALBIN)/$(LINTER)))
	@echo 'installing linter'
	@$(INSTALLLINTERBIN)
	@echo 'linter installed'
else
	@echo 'linter already installed'
endif

lint: install-linter golangci-lint

golangci-lint:
	@echo -e $(BLUE_COLOR)[golangci-lint]$(DEFAULT_COLOR)
	@mkdir -p report
	@$(LINTER) version
	@$(LINTER) run --issues-exit-code 1 --out-format code-climate | tee report/gl-code-quality-report.json | jq -r '.[] | "\(.location.path):\(.location.lines.begin) \(.description)"'; exit "$${PIPESTATUS[0]}"


build:
	@echo -e $(BLUE_COLOR)[build]$(DEFAULT_COLOR)
	$(GOBUILD) -o . ./...



install-test:
	@echo -e $(BLUE_COLOR)[install-test]$(DEFAULT_COLOR)
	@$(ISTALLTEST)

test: install-test
	@echo -e $(BLUE_COLOR)[test]$(DEFAULT_COLOR)
	@mkdir -p report
	@$(GOTEST) --junitfile report/test-junit.xml --format testname --jsonfile report/test.json -- $(GOBUILDFLAG) -race -coverprofile=report/coverage.out ./...
	@gocover-cobertura < report/coverage.out > report/coverage.xml
