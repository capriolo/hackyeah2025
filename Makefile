WEBAPP_DIR := ./apps/webapp
CORE_DIR := ./apps/core
OUTPUT_DIR := dist
CORE_MAIN_FILE := cmd/main.go

.PHONY: clean

all: core_build

webapp_install: $(WEBAPP_DIR)/package.json
	cd $(WEBAPP_DIR) && \
		npm install

webapp_run:
	cd $(WEBAPP_DIR) && \
		npm run dev

webapp_build: $(WEBAPP_DIR)/**/* webapp_install
	cd $(WEBAPP_DIR) && \
		WEBAPP_DIST_DIR=apps/core/internal/static/webapp \
		npm run build

core_run:
	cd $(CORE_DIR) && \
		ENV_PATH=../../.env go run -race cmd/main.go

core_build: $(CORE_DIR)/**/*.go webapp_build
	mkdir -p $(OUTPUT_DIR)
	cd $(CORE_DIR) && \
		go build -o ../../$(OUTPUT_DIR)/core $(CORE_MAIN_FILE)

clean:
	rm -rf $(OUTPUT_DIR) && \
	rm -rf $(WEBAPP_DIR)/dist \
