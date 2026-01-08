# Sistema de Temperatura por CEP - Desafio Go

Este projeto é um sistema em Go que recebe um CEP, identifica a cidade correspondente e retorna o clima atual (temperatura em graus Celsius, Fahrenheit e Kelvin). O projeto foi desenvolvido para ser implantado no Google Cloud Run.

## 📋 Requisitos do Desafio

**Objetivo:** Desenvolver um sistema que receba um CEP válido de 8 dígitos, encontre a localização e retorne as temperaturas formatadas.

**Cenários de Resposta:**
* **Sucesso (HTTP 200):**
  * Response Body: `{ "temp_C": 28.5, "temp_F": 83.3, "temp_K": 301.5 }`
* **Falha - CEP Inválido (HTTP 422):**
  * Mensagem: `invalid zipcode`
* **Falha - CEP Não Encontrado (HTTP 404):**
  * Mensagem: `can not find zipcode`

**Fórmulas de Conversão:**
* Fahrenheit: `F = C * 1,8 + 32`
* Kelvin: `K = C + 273`

## 🚀 Como Executar

### Pré-requisitos
* Go instalado.
* Docker e Docker Compose instalados.
* Chave de API da WeatherAPI.
  * https://www.weatherapi.com/docs/

### Configuração
É necessário configurar a chave da API de clima. Você pode fazer isso através de um arquivo `.env` ou variáveis de ambiente.

```env
WEATHER_API_KEY=sua_chave_aqui
```

### 🐳 Executando com Docker (Recomendado)

Para subir a aplicação e garantir que todas as dependências estejam corretas, utilize o Docker Compose:

```bash
docker compose up --build
```

A aplicação estará disponível em `http://localhost:8080`.

**Exemplo de uso:**
```bash
curl "http://localhost:8080?cep=01153000"
```

### Executando Localmente (Sem Docker)
```bash
go run cmd/server/main.go
```

## 🧪 Testes Automatizados

Os testes automatizados demonstram o funcionamento do sistema e podem ser executados isoladamente via Docker:

```bash
docker compose run --rm test
```

## ☁️ Deploy no Google Cloud Run

A aplicação está configurada para deploy no Google Cloud Run.

**Endereço Ativo:**
> `https://weather-fullcyclev2-117536311839.us-central1.run.app`

Para testar em produção:
```bash
curl "https://weather-fullcyclev2-117536311839.us-central1.run.app/?cep=01153000"
```
