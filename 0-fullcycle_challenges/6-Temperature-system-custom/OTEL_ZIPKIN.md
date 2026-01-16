# OTEL + Zipkin - Guia de Implementação

## Overview

Este projeto implementa **OpenTelemetry (OTEL)** com **Zipkin** para tracing distribuído entre os serviços A e B.

## Componentes

### 1. **OTEL Collector**
- Porta: 4317 (gRPC) e 4318 (HTTP)
- Responsável por receber e processar traces dos serviços
- Encaminha para o Zipkin

### 2. **Zipkin**
- Porta: 9411
- Interface web para visualizar e analisar traces
- URL: `http://localhost:9411`

### 3. **Service A** (API Gateway)
- Instrumentado com OTEL
- Cria spans para requisições POST `/`
- Propaga trace context ao chamar Service B
- Mede tempo total de processamento

### 4. **Service B** (Temperature Service)
- Instrumentado com OTEL
- Cria spans para:
  - Busca de CEP (ViaCEP API)
  - Busca de temperatura (Weather API)
- Recebe trace context do Service A

## Iniciar o Sistema

```bash
# Build e inicia todos os serviços com OTEL
docker-compose up --build

# Em outro terminal, visualizar Zipkin
open http://localhost:9411
```

## Visualizar Traces no Zipkin

1. Acesse `http://localhost:9411`
2. Na aba "Service Name", selecione:
   - `service-a` ou `service-b`
3. Clique em "Find Traces"
4. Veja o fluxo completo das requisições e timings

## Exemplo de Trace

Quando você faz uma requisição:

```bash
curl -X POST http://localhost:8081/ \
  -H "Content-Type: application/json" \
  -d '{"cep": "01310100"}'
```

O fluxo de tracing será:

```
Service A (POST /)
├── Span: "POST /" - tempo total
├── Span: "call_service_b" - tempo da chamada HTTP
│   └── Service B (GET /)
│       ├── Span: "GET /" - tempo total
│       ├── Span: "search_viacep" - tempo da API ViaCEP
│       └── Span: "get_temperature" - tempo da API Weather
```

## Estrutura de Spans

### Service A
- **Root Span**: `POST /` - Requisição principal
  - **Child Span**: `call_service_b` - Chamada HTTP ao Service B

### Service B
- **Root Span**: `GET /` - Requisição principal
  - **Child Span**: `search_viacep` - Busca de CEP
  - **Child Span**: `get_temperature` - Busca de temperatura

## Informações Capturadas nos Spans

### Events
- CEP recebido
- Validação de CEP
- Erro ou sucesso da operação

### Attributes
- `status_code` - Código HTTP da resposta
- `city` - Nome da cidade encontrada
- `temp_C` - Temperatura em Celsius

## Variáveis de Ambiente

```bash
# Endpoint do OTEL Collector (já configurado no docker-compose)
OTEL_EXPORTER_OTLP_ENDPOINT=http://otel-collector:4317

# Nome do serviço (já configurado no docker-compose)
OTEL_SERVICE_NAME=service-a (ou service-b)
```

## Troubleshooting

### Zipkin não mostra traces

1. Verifique se o OTEL Collector está rodando:
   ```bash
   docker ps | grep otel-collector
   ```

2. Verifique os logs do collector:
   ```bash
   docker logs otel-collector
   ```

3. Verifique se os serviços conseguem se conectar ao collector:
   ```bash
   docker logs service-a
   docker logs service-b
   ```

### Spans não aparecem em tempo real

- Zipkin pode levar alguns segundos para exibir os traces
- Tente aguardar 5-10 segundos após fazer a requisição
- Atualize a página do Zipkin (F5)

## Performance Insights do Zipkin

Use o Zipkin para analisar:

1. **Latência total** - Tempo total da requisição
2. **Latência por serviço** - Quanto tempo cada serviço leva
3. **Gargalos** - Quais chamadas externas demoram mais
4. **Distribuição de erros** - Quais serviços geram erros

## Parar os Serviços

```bash
docker-compose down
```

## Referências

- [OpenTelemetry Go Documentation](https://opentelemetry.io/docs/languages/go/getting-started/)
- [OTEL Spans Documentation](https://opentelemetry.io/docs/languages/go/instrumentation/#creating-spans)
- [OTEL Collector](https://opentelemetry.io/docs/collector/quick-start/)
- [Zipkin Documentation](https://zipkin.io/)
