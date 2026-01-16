# Temperature System - Fase 1

Este projeto implementa um sistema de consulta de temperatura por CEP usando dois serviços Go em containers Docker com **tracing distribuído via OpenTelemetry e Zipkin**.

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
- **Tracing**: Instrumentado com OpenTelemetry

### Serviço B (Service B)
- **Porta**: 8080
- **Função**: Consulta localidade via CEP (ViaCEP API) e retorna temperatura (WeatherAPI API)
- **Endpoint**: `GET /?cep={cep}`
- **Resposta**: `{ "city": "São Paulo", "temp_C": 25.5, "temp_F": 77.9, "temp_K": 298.65 }`
- **Tracing**: Instrumentado com OpenTelemetry, mede tempo de ViaCEP e Weather API

### OTEL Collector
- **Porta**: 4317 (gRPC), 4318 (HTTP)
- **Função**: Coleta e processa traces dos serviços

### Zipkin
- **Porta**: 9411
- **Função**: Visualiza e analisa traces distribuídos
- **URL**: `http://localhost:9411`

## Pré-requisitos

- Docker
- Docker Compose
- Uma chave de API da WeatherAPI
   - https://www.weatherapi.com/docs/
## Setup

1. Clone o arquivo `.env.example` para `.env`:
```bash
cp .env.example .env
```

2. Adicione sua chave de API da WeatherAPI no arquivo `.env`:
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
    │   ├── storage/
    │   └── telemetry/     # OTEL configuration
    ├── Dockerfile
    ├── docker-compose.yml
    └── go.mod
```

## Observabilidade - OpenTelemetry + Zipkin

Este projeto inclui implementação completa de **tracing distribuído**:

### Acessar Zipkin
```bash
# Após iniciar com docker-compose, acesse:
open http://localhost:9411
```

### Visualizar Traces
1. No Zipkin, selecione o serviço (service-a ou service-b)
2. Clique em "Find Traces"
3. Veja o fluxo completo das requisições com timings

### Spans Capturados

**Service A:**
- `POST /` - Requisição principal
  - `call_service_b` - Chamada HTTP ao Service B

**Service B:**
- `GET /` - Requisição principal
  - `search_viacep` - Busca de CEP
  - `get_temperature` - Busca de temperatura

**Métricas Capturadas:**
- Tempo total de processamento
- Tempo de cada serviço externo (ViaCEP, Weather API)
- Status das operações
- Informações como CEP, cidade e temperatura

Para mais detalhes, veja [OTEL_ZIPKIN.md](./OTEL_ZIPKIN.md)

## Parar os serviços

```bash
docker-compose down
```

