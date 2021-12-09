OUT_DIR ?= ./bin/kcli

.PHONY: compile
compile:
	go build -o $(OUT_DIR) .