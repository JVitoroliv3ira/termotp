VERSION := $(shell git describe --tags --always || echo "dev")

BINARY_NAME = termotp

BUILD_DIR = build

PLATFORMS = linux/amd64 darwin/amd64 windows/amd64

LDFLAGS = -X 'github.com/JVitoroliv3ira/termotp/internal/version.Version=$(VERSION)'

.DEFAULT_GOAL := help

help:
	@echo "Comandos disponíveis:"
	@echo "  make build        - Compila o projeto para a plataforma atual"
	@echo "  make release      - Compila binários para Linux, macOS e Windows"
	@echo "  make clean        - Remove os binários compilados"
	@echo "  make version      - Exibe a versão atual do projeto"

build:
	@echo "🚀 Compilando TermOTP versão $(VERSION)..."
	mkdir -p $(BUILD_DIR)
	go build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(BINARY_NAME)
	@echo "✅ Build concluído: $(BUILD_DIR)/$(BINARY_NAME)"

release: clean
	@echo "🚀 Criando release para versão $(VERSION)..."
	mkdir -p $(BUILD_DIR)
	$(foreach platform, $(PLATFORMS), \
		$(eval GOOS=$(word 1,$(subst /, ,$(platform)))) \
		$(eval GOARCH=$(word 2,$(subst /, ,$(platform)))) \
		echo "🔹 Compilando para $(GOOS)/$(GOARCH)..." && \
		GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(BINARY_NAME)-$(GOOS)-$(GOARCH); \
	)
	@echo "✅ Todos os binários foram gerados em $(BUILD_DIR)/"

clean:
	@echo "🧹 Removendo arquivos de build..."
	rm -rf $(BUILD_DIR)
	@echo "✅ Build limpo!"

version:
	@echo "🔹 TermOTP - Versão $(VERSION)"
