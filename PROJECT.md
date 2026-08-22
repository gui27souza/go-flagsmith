# 📦 Microsserviço Resiliente com Flagsmith & Go

## Objetivo

Desenvolver uma API REST em Go que consuma Feature Flags e Remote Config de maneira não bloqueante, resiliente e segura para ambientes de alta volumetria.

## Stack Técnica

- **Linguagem:** Go (1.22+)
- **Router:** `gin`
- **SDK:** Flagsmith Go SDK (com cache local e polling em background)
- **Container:** Docker (Multi-stage build com imagem `distroless` ou `alpine`)

## Principais Entregas & Requisitos

- [ ]  **Integração com SDK:** Configurar cliente Flagsmith com cache em memória e intervalo de sincronização.
- [ ]  **Padrão Fallback Seguro:** Garantir que falhas de rede ou timeout na API do Flagsmith não causem HTTP 500 no serviço (uso de valores default predefinidos).
- [ ]  **Dynamic Remote Config:** Criar endpoint que responde com regras e limites operacionais (ex: `rate_limit_per_second`, `maintenance_mode`) alteráveis em tempo real.
- [ ]  **Observabilidade & Health Check:** Endpoints `/healthz` (liveness) e `/readyz` (readiness) refletindo a integridade da aplicação.
- [ ]  **Testes Automatizados:** Testes unitários utilizando *interfaces* e *mocks* para simular comportamento com flags ativas, inativas e falhas de conexão.

---

## Notas

### 22/8

- Comecei o projeto entendendo um pouco melhor sobre o gin, um framework para APIs REST em Go
    - Segui o tutorial simplezinho do web-service-gin