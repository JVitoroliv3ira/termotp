# **TermOTP** 🛡️🔑

[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)
[![Última versão](https://img.shields.io/github/v/release/JVitoroliv3ira/termotp?label=Download)](https://github.com/JVitoroliv3ira/termotp/releases/latest)

**TermOTP** é uma ferramenta de linha de comando (**CLI**) para **gerenciar e gerar códigos TOTP** (Time-based One-Time Password) de forma **segura** e **offline**.  
Com **criptografia AES-256** e suporte a múltiplas contas, ele permite que você centralize tokens de autenticação de diversos serviços diretamente no terminal.

---

## **📥 Instalação**
A versão mais recente do **TermOTP** está disponível em **[Releases](https://github.com/JVitoroliv3ira/termotp/releases/latest)**.

### **🔹 Linux**
```sh
# Remover versão antiga (se existir)
sudo rm -f /usr/local/bin/termotp

# Baixar a nova versão
wget https://github.com/JVitoroliv3ira/termotp/releases/latest/download/termotp-linux-amd64 -O termotp

# Tornar executável e mover para /usr/local/bin
chmod +x termotp
sudo mv termotp /usr/local/bin/
```
Agora, basta rodar `termotp` de qualquer lugar no terminal! 🚀

### **🔹 Windows**
```powershell
# Remover versão antiga (se existir)
Remove-Item "C:\Program Files\TermOTP\termotp.exe" -ErrorAction SilentlyContinue

# Criar diretório se não existir
mkdir "C:\Program Files\TermOTP" -ErrorAction SilentlyContinue

# Baixar a nova versão
Invoke-WebRequest -Uri "https://github.com/JVitoroliv3ira/termotp/releases/latest/download/termotp-windows-amd64.exe" -OutFile "C:\Program Files\TermOTP\termotp.exe"

# Adicionar ao PATH (se ainda não estiver)
[System.Environment]::SetEnvironmentVariable("Path", $Env:Path + ";C:\Program Files\TermOTP", [System.EnvironmentVariableTarget]::Machine)
```
Após isso, reinicie o terminal e rode `termotp` de qualquer lugar! 🎉

---

## **🚀 Como Usar**
Após instalar o **TermOTP**, você pode rodar o seguinte comando para ver todas as opções disponíveis:

```sh
termotp --help
```

### **Comandos Disponíveis**
- **Gerar um código TOTP:** `termotp generate`
- **Copiar um código sem exibir:** `termotp copy`
- **Listar todas as contas e códigos:** `termotp list`
- **Adicionar uma nova conta:** `termotp setup`
- **Ver a versão instalada:** `termotp version`

Para mais detalhes sobre os comandos, acesse a **[documentação completa](https://github.com/JVitoroliv3ira/termotp/wiki)**.

---

## **📦 Releases**
A versão mais recente do **TermOTP** pode ser encontrada em **[Releases](https://github.com/JVitoroliv3ira/termotp/releases/latest)**.

Cada versão inclui:
- Binários pré-compilados para **Linux** e **Windows**.
- Histórico de mudanças e novas funcionalidades.

---

## **📜 Licença**
Este projeto é distribuído sob a **Licença GPLv3**.  
Consulte o arquivo [LICENSE](./LICENSE) ou acesse a [GNU General Public License v3.0](https://www.gnu.org/licenses/gpl-3.0.html) para mais detalhes.

---

## **🌟 Apoie este projeto!**
Se você gostou do **TermOTP**, deixe uma ⭐ no repositório!  
Isso ajuda o projeto a crescer e alcançar mais pessoas. 😃🚀