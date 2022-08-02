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

VLINTER            = "v1.47.3"

install-linter:
	@echo -e $(BLUE_COLOR)[install-linter]$(DEFAULT_COLOR)
ifeq (,$(wildcard $(LOCALBIN)/$(LINTER)))
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
