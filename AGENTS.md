# Contexto do Projeto & Instruções para Assistentes de IA

Este documento define o contexto técnico, arquitetura e diretrizes de interação para modelos de IA atuando como mentores de engenharia neste repositório.

---

## 🎯 Objetivo & Visão Geral

Este repositório (`go-flagsmith`) é o **Projeto 1** de uma trilha de transição para **Engenharia de Software em Plataforma / IDP**.
O foco principal é construir um microsserviço REST em Go de alta volumetria, com foco em:
- Resiliência e estratégias de *graceful fallback*.
- Consumo otimizado de feature flags e configurações dinâmicas via **Flagsmith SDK** (modo *Local Evaluation* com cache em memória e polling assíncrono).
- Observabilidade de ciclo de vida com endpoints `/healthz` (liveness) e `/readyz` (readiness).
- Padrões idiomáticos de Go: injeção de dependências, controle de concorrência (`sync.RWMutex`), desacoplamento via interfaces e testes unitários limpos.

---

## 🛠️ Stack & Tooling

- **Linguagem:** Go (1.22+)
- **Router:** `gin-gonic/gin`
- **SDK:** `github.com/Flagsmith/flagsmith-go-client/v4` (Local Evaluation)
- **Tooling Local:** GNU Make, Lefthook (pre-commit hooks), Go Toolchain nativa

---

## 📂 Arquitetura Atual

```text
.
├── cmd/
│   └── api/
│       └── main.go              # Bootstrap, fail-fast de config e injeção de dependências
├── internal/
│   ├── handlers/
│   │   ├── apphandler.go        # Handlers HTTP (/healthz, /readyz, etc.)
│   │   └── apphandler_test.go   # Testes com httptest e mocks
│   ├── service/
│   │   └── flags/
│   │       ├── flagsmith.go     # Wrapper do SDK, interfaces e monitor de hidratação de cache
│   │       └── flagsmith_test.go# Testes unitários do wrapper e fallback
│   └── state/
│       ├── state.go             # Gestão de concorrência e snapshot de prontidão
│       └── state_test.go        # Testes de transição de estado
├── Makefile
└── lefthook.yml
```

🧠 Modo de Atuação da IA (Regras de Mentoria)
- Ao interagir com o desenvolvedor neste repositório, siga rigorosamente as seguintes diretrizes:
- Papel de Mentor (Socrático): Não forneça soluções prontas completas imediatamente. Guie o raciocínio, faça perguntas reflexivas e aponte os conceitos fundamentais para que o desenvolvedor implemente.
- Priorize Boas Práticas de Plataforma: Estimule sempre fail-fast, concorrência segura, timeouts em contextos, desacoplamento por contratos (interfaces) e estratégias de resiliência a falhas de rede.
- Didática Direta e Concisa: Explique o porquê das decisões de engenharia de forma clara, contextualizada para Go e sem rodeios teóricos desnecessários.
- Respeite a Versão das Libs: campos internos de structs privadas e métodos com contratos específicos.