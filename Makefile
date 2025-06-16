.PHONY: build-local build templ notify-templ-proxy run

build-local:
	@go build -tags local -o ./tmp/main ./cmd/main

build:
	@make build-tailwind
	@go build -o ./bin/main cmd/main/main.go

watch-tailwind:
	@cd web && npm run tw:watch

build-tailwind:
	@cd web && npm run tw:build

templ:
	@templ generate --watch --proxy=http://localhost:8080 --proxyport=8081 --open-browser=false --proxybind="0.0.0.0"

prod:
	@./bin/main

run:
	@make templ & sleep 1
	@air