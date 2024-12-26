# Rate Limiter em Go

Este projeto implementa um rate limiter em Go que controla o tráfego de requisições para um serviço web, permitindo a limitação de requisições com base em dois critérios: endereço IP e token de acesso.

## Estrutura do Projeto

```
rate-limiter
├── cmd
│   └── main.go               # Ponto de entrada da aplicação
├── internal
│   ├── limiter
│   │   ├── limiter.go        # Lógica principal do rate limiter
│   │   └── middleware.go     # Implementação do middleware do rate limiter
│   ├── config
│   │   └── config.go         # Gerenciamento de configuração
│   └── storage
│       ├── redis.go          # Implementação do armazenamento com Redis
│       └── strategy.go       # Estratégia de persistência
├── .env                       # Variáveis de ambiente
├── docker-compose.yml         # Configuração do serviço Redis
├── go.mod                     # Gerenciamento de dependências
└── go.sum                     # Somas de verificação das dependências
```

## Instalação

1. Clone o repositório:
   ```
   git clone <URL_DO_REPOSITORIO>
   cd rate-limiter
   ```

2. Crie um arquivo `.env` na raiz do projeto com as seguintes variáveis:
   ```
   RATE_LIMIT_IP=5
   RATE_LIMIT_TOKEN=10
   BLOCK_TIME=300
   REDIS_HOST=localhost
   REDIS_PORT=6379
   ```

3. Inicie o Redis usando Docker:
   ```
   docker-compose up -d
   ```

4. Execute a aplicação:
   ```
   go run cmd/main.go
   ```

## Uso

O rate limiter pode ser utilizado como middleware em um servidor web. Ele limita o número de requisições por IP e por token de acesso, respondendo com um código HTTP 429 quando os limites são excedidos.

## Contribuição

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues ou pull requests.

## Licença

Este projeto está licenciado sob a MIT License. Veja o arquivo LICENSE para mais detalhes.