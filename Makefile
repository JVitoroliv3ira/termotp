VERSION := $(shell git describe --tags --always || echo "dev")

BINARY_NAME = totp

BUILD_DIR = build

PLATFORMS = linux/amd64 darwin/amd64 windows/amd64

LDFLAGS = -X 'github.com/JVitoroliv3ira/termotp/internal/version.Version=$(VERSION)'

.DEFAULT_GOAL := help

help:
	@echo "Comandos disponíveis:"
	@echo "  make test         - Executa os testes da aplicação"
	@echo "  make build        - Compila o projeto para a plataforma atual"
	@echo "  make release      - Compila binários para Linux, macOS e Windows"
	@echo "  make build-linux  - Compila o binário para Linux"
	@echo "  make build-mac    - Compila o binário para macOS"
	@echo "  make build-win    - Compila o binário para Windows"
	@echo "  make clean        - Remove os binários compilados"
	@echo "  make version      - Exibe a versão atual do projeto"

test:
	go test -v -cover ./...

build:
	@echo "🚀 Compilando TermOTP versão $(VERSION)..."
	mkdir -p $(BUILD_DIR)
	go build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(BINARY_NAME)
	@echo "✅ Build concluído: $(BUILD_DIR)/$(BINARY_NAME)"

build-linux:
	@echo "🐧 Compilando para Linux..."
	mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64
	@echo "✅ Build para Linux concluído: $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64"

build-mac:
	@echo "🍏 Compilando para macOS..."
	mkdir -p $(BUILD_DIR)
	GOOS=darwin GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(BINARY_NAME)-macos-amd64
	@echo "✅ Build para macOS concluído: $(BUILD_DIR)/$(BINARY_NAME)-macos-amd64"

build-win:
	@echo "🖥️ Compilando para Windows..."
	mkdir -p $(BUILD_DIR)
	GOOS=windows GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe
	@echo "✅ Build para Windows concluído: $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe"

release: clean
	@echo "🚀 Criando release para versão $(VERSION)..."
	mkdir -p $(BUILD_DIR)
	$(foreach platform, $(PLATFORMS), \
		$(eval GOOS=$(word 1,$(subst /, ,$(platform)))) \
		$(eval GOARCH=$(word 2,$(subst /, ,$(platform)))) \
		echo "🔹 Compilando para $(GOOS)/$(GOARCH)..." && \
		GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(BINARY_NAME)-$(GOOS)-$(GOARCH)$(if $(filter windows,$(GOOS)),.exe,); \
	)
	@echo "✅ Todos os binários foram gerados em $(BUILD_DIR)/"

clean:
	@echo "🧹 Removendo arquivos de build..."
	rm -rf $(BUILD_DIR)
	@echo "✅ Build limpo!"

version:
	@echo "🔹 TermOTP - Versão $(VERSION)"
