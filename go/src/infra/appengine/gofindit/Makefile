GAE_PY = ../../../../../luci/appengine/components/tools/gae.py

default: help

help:
	@echo "Available commands:"
	@sed -n '/^[a-zA-Z0-9_]*:/s/:.*//p' <Makefile

format:
	gofmt -s -w .

test:
	go test ./...

deploy:
	$(GAE_PY) upload -A gofindit-dev

deploy-prod:
	$(GAE_PY) upload -A gofindit
