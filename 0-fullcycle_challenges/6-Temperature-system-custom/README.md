# Temperature System - Fase 1

Este projeto implementa um sistema de consulta de temperatura por CEP usando dois serviços Go em containers Docker.

## Arquitetura

### Serviço A (Service A)
- **Porta**: 8081
- **Função**: API que recebe CEP via POST e consulta o Serviço B
- **Endpoint**: `POST /`
  - Request: `{ "cep": "01310100" }`
  - Resposta de sucesso (HTTP 200): `{ "city": "São Paulo", "temp_C": 25.5, "temp_F": 77.9, "temp_K": 298.65 }`
  - CEP inválido (HTTP 422): `{ "message": "invalid zipcode" }`
  - CEP não encontrado (HTTP 404): `{ "message": "can not find zipcode" }`
- **Validação**: CEP deve ter exatamente 8 dígitos e ser uma STRING
- **Health Check**: `GET /health`

### Serviço B (Service B)
- **Porta**: 8080
- **Função**: Consulta localidade via CEP (ViaCEP API) e retorna temperatura (OpenWeatherMap API)
- **Endpoint**: `GET /?cep={cep}`
- **Resposta**: `{ "city": "São Paulo", "temp_C": 25.5, "temp_F": 77.9, "temp_K": 298.65 }`

## Pré-requisitos

- Docker
- Docker Compose
- Uma chave de API da OpenWeatherMap

## Setup

1. Clone o arquivo `.env.example` para `.env`:
```bash
cp .env.example .env
```

2. Adicione sua chave de API da OpenWeatherMap no arquivo `.env`:
```
WEATHER_API_KEY=sua_chave_aqui
```

## Como executar

Execute com Docker Compose:
```bash
docker-compose up --build
```

## Testando os serviços

### Teste Service A (recomendado):

**CEP válido:**
```bash
curl -X POST http://localhost:8081/ \
  -H "Content-Type: application/json" \
  -d '{"cep": "01310100"}'
```

Resposta esperada:
```json
{
  "city": "São Paulo",
  "temp_C": 25.5,
  "temp_F": 77.9,
  "temp_K": 298.65
}
```

**CEP inválido (menos de 8 dígitos):**
```bash
curl -X POST http://localhost:8081/ \
  -H "Content-Type: application/json" \
  -d '{"cep": "123"}'
```

Resposta esperada (HTTP 422):
```json
{
  "message": "invalid zipcode"
}
```

**CEP válido mas não encontrado:**
```bash
curl -X POST http://localhost:8081/ \
  -H "Content-Type: application/json" \
  -d '{"cep": "00000000"}'
```

Resposta esperada (HTTP 404):
```json
{
  "message": "can not find zipcode"
}
```

### Teste Service B diretamente:
```bash
curl "http://localhost:8080/?cep=01310100"
```

### Health Check do Service A:
```bash
curl "http://localhost:8081/health"
```

## Estrutura do projeto

```
6-Temperature-system-custom/
├── docker-compose.yml     # Configuração dos serviços
├── .env.example          # Exemplo de variáveis de ambiente
├── README.md             # Este arquivo
├── service-a/            # Serviço A (API Gateway)
│   ├── cmd/
│   │   └── main.go
│   ├── Dockerfile
│   └── go.mod
└── service-b/            # Serviço B (Temperature Service)
    ├── cmd/
    │   └── server/
    │       └── main.go
    ├── internal/
    │   ├── core/
    │   ├── dto/
    │   └── storage/
    ├── Dockerfile
    ├── docker-compose.yml
    └── go.mod
```

## Parar os serviços

```bash
docker-compose down
```

## Melhorias futuras

- Adicionar cache de resultados
- Implementar rate limiting
- Adicionar autenticação e autorização
- Melhorar tratamento de erros
- Adicionar testes unitários
- Implementar logging estruturado
