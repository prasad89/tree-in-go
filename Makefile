APP_NAME := tree
BUILD_DIR := bin

.PHONY: all
all: build

.PHONY: build
build:
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME)

.PHONY: run
run: build
	./$(BUILD_DIR)/$(APP_NAME)

.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)
