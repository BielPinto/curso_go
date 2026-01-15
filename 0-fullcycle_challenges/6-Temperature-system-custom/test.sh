#!/bin/bash

# Script para testar os serviços

echo "Testing Service A (API Gateway)..."
echo "=================================="
echo ""

# Health check
echo "1. Health Check:"
curl -s http://localhost:8081/health | jq '.'
echo ""
echo ""

# CEP válido - São Paulo
echo "2. Consultando temperatura para CEP 01310100 (São Paulo):"
curl -s -X POST http://localhost:8081/ \
  -H "Content-Type: application/json" \
  -d '{"cep": "01310100"}' | jq '.'
echo ""
echo ""

# CEP válido - Rio de Janeiro
echo "3. Consultando temperatura para CEP 20040020 (Rio de Janeiro):"
curl -s -X POST http://localhost:8081/ \
  -H "Content-Type: application/json" \
  -d '{"cep": "20040020"}' | jq '.'
echo ""
echo ""

# CEP inválido - menos de 8 dígitos
echo "4. Testando com CEP inválido (menos de 8 dígitos - deve retornar 422):"
curl -s -X POST http://localhost:8081/ \
  -H "Content-Type: application/json" \
  -d '{"cep": "12345"}' | jq '.'
echo ""
echo ""

# CEP inválido - contém letras
echo "5. Testando com CEP inválido (contém letras - deve retornar 422):"
curl -s -X POST http://localhost:8081/ \
  -H "Content-Type: application/json" \
  -d '{"cep": "0131010a"}' | jq '.'
echo ""
echo ""

# CEP válido mas inexistente
echo "6. Testando com CEP válido mas inexistente (deve retornar 404):"
curl -s -X POST http://localhost:8081/ \
  -H "Content-Type: application/json" \
  -d '{"cep": "00000000"}' | jq '.'
echo ""
echo ""

echo "Tests completed!"
