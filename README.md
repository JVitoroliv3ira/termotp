# TermOTP

[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

**TermOTP** é uma ferramenta de linha de comando (CLI) para gerenciar e gerar códigos **TOTP** (Time-based One-Time Password).  
O objetivo é unificar tokens de diversos serviços (Google Authenticator, Microsoft Authenticator etc.) e permitir a criação de novos tokens diretamente no terminal.

---

## Recursos Principais

- **Importar tokens** via links `otpauth-migration://offline?data=...`.  
- **Listar** contas TOTP e gerar códigos atualizados no terminal.  
- **Criar** novas contas (secrets) para habilitar 2FA em qualquer serviço compatível com TOTP.  
- Possível **integração** com armazenamento seguro (ex.: GPG, `pass`, Keepass).  
- **Extensibilidade** para diferentes fluxos e necessidades de autenticação.

---

## Roadmap

- [ ] Implementar importação de tokens usando Protobuf.  
- [ ] Suportar múltiplas contas em um arquivo ou vault (listagem detalhada).  
- [ ] Gerar QR codes para facilitar configuração em outros dispositivos.  
- [ ] Integrações com serviços de armazenamento seguro de senhas.  
- [ ] Interface CLI aprimorada (cores, prompts interativos, etc.).

---

## Contribuindo

1. Faça um **fork** deste repositório.  
2. Crie um branch para sua nova feature ou correção de bug:
   ```bash
   git checkout -b minha-feature
   ```
3. Após realizar as alterações, faça commits com mensagens claras:
   ```bash
   git commit -m "Descreva a mudança"
   ```
4. Envie seu branch para o GitHub:
   ```bash
   git push origin minha-feature
   ```
5. Abra um **Pull Request** descrevendo as mudanças propostas.

Contribuições em forma de código, documentação, testes ou sugestões são bem-vindas!

---

## Licença

Este projeto está disponível sob a **Licença GPLv3**. Consulte o arquivo [LICENSE](./LICENSE) ou acesse a página da [GNU General Public License v3.0](https://www.gnu.org/licenses/gpl-3.0.html) para mais detalhes.

---

**TermOTP** é um projeto **open source** desenvolvido para simplificar e centralizar a gestão de tokens TOTP. Fique à vontade para explorar, contribuir e compartilhar!
