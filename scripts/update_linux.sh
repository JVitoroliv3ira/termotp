#!/bin/bash
set -e

INSTALL_PATH="/usr/local/bin/totp"
BACKUP_PATH="/usr/local/bin/totp.bak"
NEW_BINARY="/tmp/totp_new"
DOWNLOAD_URL="https://github.com/JVitoroliv3ira/termotp/releases/latest/download/totp-linux-amd64"

rollback() {
    echo "Ocorreu um erro. Restaurando a versão anterior..."
    if [ -f "$BACKUP_PATH" ]; then
        sudo mv "$BACKUP_PATH" "$INSTALL_PATH"
        echo "Rollback realizado: binário restaurado."
    else
        echo "Nenhum backup encontrado para restaurar."
    fi
    exit 1
}

trap 'rollback' ERR

echo "Iniciando atualização do TermOTP..."

if [ -f "$INSTALL_PATH" ]; then
    sudo cp "$INSTALL_PATH" "$BACKUP_PATH"
    echo "Backup criado em $BACKUP_PATH"
fi

wget "$DOWNLOAD_URL" -O "$NEW_BINARY"
echo "Novo binário baixado para $NEW_BINARY"

chmod +x "$NEW_BINARY"
echo "Permissões de execução aplicadas ao novo binário"

sudo mv "$NEW_BINARY" "$INSTALL_PATH"
echo "Novo binário movido para $INSTALL_PATH"

if [ -f "$BACKUP_PATH" ]; then
    sudo rm -f "$BACKUP_PATH"
    echo "Backup removido."
fi

echo "Atualização concluída com sucesso!"
