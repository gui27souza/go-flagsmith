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

## 23/8

- De início, tive uma conversa com o Gemini para entender os benefícios de configurações específicas na inicialização do sdk do `flagsmith`
    - A ideia é que ele baixe todas as flags localmente ao inicializar e tenha um melhor gerenciamento da leitura dessas flags, em memória
- Além disso, isolei de fato o inicializador do sdk por trás de uma struct/interface, em `internal/services/flags.go` , facilitando a manutenção e testabilidade
- Descobri 3 ferramentas para tornar a saída de testes mais agradável:
```
// gotestsum
go install gotest.tools/gotestsum@latest
gotestsum --format pkgname

// tparse
go install github.com/mfridman/tparse@latest
go test -json ./... | tparse -all

// coverage nativo
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## 24/8

- Fiz uns refinamentos no DevEx, cons uns ajustezinhos no `Makefile`, e com checagem de env e alguns encadeamentos de comando
- Pude finalmente testar a aplicação! E deu tudo certo!!
    - Funcionou o start do servidor, hidradação das flags e os endpoints de saúde `/healthz` e `/readyz`!!

- Para o encor