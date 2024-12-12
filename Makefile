SHELL=/bin/bash
PROJECT_DIRECTORY := $(shell pwd)
CFLAGS=-g
export CFLAGS
target:
	$(MAKE) -C target

local-start:
	go run ebiten/main.go

local-build:
	docker build -t game-builder -f cicd/Dockerfile .
	docker create --name game-builder-container game-builder ls -lah /app/
	docker cp game-builder-container:/app/tower-defense .
	docker rm -f game-builder-container

local-build-wasm:
	docker build -t game-builder -f cicd/Wasm.Dockerfile .
	docker create --name game-builder-container game-builder ls -lah /app/
	docker cp game-builder-container:/app/canon-defense.wasm html/
	docker cp game-builder-container:/usr/local/go/misc/wasm/wasm_exec.js html/
	docker rm -f game-builder-container