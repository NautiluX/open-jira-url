.PHONY: all
all: build

.PHONY: build
build:
	go build ./cmd/open-jira-url

install:
	go install ./cmd/open-jira-url
