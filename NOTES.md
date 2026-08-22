## 22/8

- Comecei o projeto entendendo um pouco melhor sobre o gin, um framework para APIs REST em Go
    - Segui o tutorial simplezinho do [web-service-gin](https://go.dev/doc/tutorial/web-service-gin)
- Iniciei a implementação real agora, comecei no simples ainda, nas rotas de `health check`
    - Os endpoints `/healthz` e `/readyz`
    - Gerenciador de estado com RWMutex, com flags booleanas sobre o estado da aplicação e Snapshot de estado, para prover ao endpoint de `/readyz`
    - Inicialização do cliente do `Flagsmith`
- Pensando em Dev-Ex, também criei um `Makefile` e um `lefthook.yml` que formata o código ao commitar

```
.
│   .gitignore
│   Dockerfile
│   go.mod
│   go.sum
│   lefthook.yml
│   Makefile
│   PROJECT.md
│   README.md
├───cmd
│   └───api
│           main.go
└───internal
    ├───handlers
    │       apphandler.go
    └───state
            state.go
```