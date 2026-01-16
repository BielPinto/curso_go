# Quick Start Guide

## 1. Setup inicial

```bash
# Copie o arquivo .env.example para .env
cp .env.example .env

# Abra o arquivo .env e adicione sua chave de API da OpenWeatherMap
# WEATHER_API_KEY=sua_chave_aqui
```

## 2. Executar com Docker Compose

```bash
# Build e inicia ambos os serviços
docker-compose up --build

# Ou apenas inicia (se já foi feito build anteriormente)
docker-compose up
```

## 3. Testar os serviços

Em outro terminal:

```bash
# Teste rápido - Consultar temperatura para São Paulo
curl -X POST http://localhost:8081/ \
  -H "Content-Type: application/json" \
  -d '{"cep": "01310100"}'

# Resposta esperada:
# {"city":"São Paulo","temp_C":25.5,"temp_F":77.9,"temp_K":298.65}
```

## 4. Visualizar Traces (OpenTelemetry + Zipkin)

Em outro terminal ou navegador:

```bash
# Acesse Zipkin
open http://localhost:9411
```

1. Selecione o serviço (service-a ou service-b)
2. Clique em "Find Traces"
3. Veja o fluxo completo das requisições com timings de cada componente

## 5. Parar os serviços

```bash
docker-compose down
```

## Endpoints disponíveis

| Serviço | Método | Endpoint | Request | Resposta |
|---------|--------|----------|---------|----------|
| Service A | POST | `/` | `{"cep":"01310100"}` | `{"city":"São Paulo","temp_C":25.5,"temp_F":77.9,"temp_K":298.65}` |
| Service A | GET | `/health` | - | `{"status":"ok"}` |
| Service B | GET | `/?cep={cep}` | - | `{"city":"São Paulo","temp_C":25.5,"temp_F":77.9,"temp_K":298.65}` |

## Exemplos de CEPs para testar

- 01310100 - São Paulo, SP
- 20040020 - Rio de Janeiro, RJ
- 30130010 - Belo Horizonte, MG
- 70040902 - Brasília, DF

## Validações do Service A

- CEP deve ter **exatamente 8 dígitos**
- CEP deve ser uma **STRING**
- Caso inválido: HTTP 422 com mensagem `{"message": "invalid zipcode"}`
- CEP não encontrado: HTTP 404 com mensagem `{"message": "can not find zipcode"}`

## Serviços Disponíveis

| Serviço | URL | Porta |
|---------|-----|-------|
| Service A (API) | http://localhost:8081 | 8081 |
| Service B (Temperature) | http://localhost:8080 | 8080 |
| Zipkin (Tracing) | http://localhost:9411 | 9411 |
| OTEL Collector | grpc://localhost:4317 | 4317 |

Para documentação completa de OpenTelemetry + Zipkin, veja [OTEL_ZIPKIN.md](../OTEL_ZIPKIN.md)
