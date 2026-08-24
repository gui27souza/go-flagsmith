# go-flagsmith

Microsserviço REST em Go projetado para demonstrar padrões de resiliência, observabilidade e consumo de feature flags e configurações dinâmicas via Flagsmith SDK.

Este projeto faz parte da construção de uma base para Internal Developer Platform (IDP).

---

## 🎯 Objetivos do Projeto

- **Feature Flags & Remote Config:** Consumo eficiente com foco em baixa latência e resiliência a falhas de rede.
- **Probes de Ciclo de Vida:** Endpoints `/healthz` (liveness) e `/readyz` (readiness) refletindo o estado real dos subsistemas da aplicação.
- **Resiliência:** Estratégias de fallback seguro para evitar indisponibilidade em caso de instabilidade de serviços externos.
- **Boas Práticas em Go:** Concorrência segura (`sync.RWMutex`), injeção de dependências e testes automatizados.

---

## 🛠️ Stack Técnica

- **Linguagem:** Go (1.22+)
- **Router HTTP:** Gin Web Framework
- **SDK:** Flagsmith Go Client (Local Evaluation / Polling)
- **Containerização:** Docker

---

## 🚀 Como Executar

### Pré-requisitos
- Go 1.22+
- Make (opcional, para automação de tarefas)

### Variáveis de Ambiente
Crie um arquivo `.env` ou exporte a chave da API do Flagsmith:

```bash
export FLAGSMITH_API_KEY="sua_chave_aqui"
```

### Executando Localmente
```bash
# Baixar dependências
go mod tidy

# Executar a API
go run cmd/api/main.go
```

---

## 📌 Status do Projeto

> ⚠️ **Em desenvolvimento ativo.** Estruturas de pacotes, rotas e integrações estão sendo refinadas continuamente.