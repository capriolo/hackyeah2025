WEBAPP_DIR := ./apps/webapp
CORE_DIR := ./apps/core

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
