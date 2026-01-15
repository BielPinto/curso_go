.PHONY: help build up down logs clean test

help:
	@echo "Comandos disponíveis:"
	@echo "  make build      - Build das imagens Docker"
	@echo "  make up         - Inicia os serviços"
	@echo "  make down       - Para os serviços"
	@echo "  make logs       - Mostra logs dos serviços"
	@echo "  make clean      - Remove containers e imagens"
	@echo "  make test       - Executa testes dos endpoints"
	@echo "  make setup      - Setup inicial (.env)"

setup:
	@cp .env.example .env
	@echo "✓ Arquivo .env criado. Adicione sua WEATHER_API_KEY!"

build:
	@echo "Building Docker images..."
	docker-compose build

up:
	@echo "Starting services..."
	docker-compose up -d
	@echo "✓ Services started!"
	@echo "Service A: http://localhost:8081"
	@echo "Service B: http://localhost:8080"

down:
	@echo "Stopping services..."
	docker-compose down

logs:
	docker-compose logs -f

logs-a:
	docker-compose logs -f service-a

logs-b:
	docker-compose logs -f service-b

clean:
	@echo "Cleaning up..."
	docker-compose down -v
	docker system prune -f

test:
	@echo "Testing Service A endpoints..."
	@echo ""
	@echo "1. Health Check:"
	@curl -s http://localhost:8081/health | jq '.' || echo "Service not available"
	@echo ""
	@echo "2. Temperature for CEP 01310100 (São Paulo):"
	@curl -s -X POST http://localhost:8081/ \
	  -H "Content-Type: application/json" \
	  -d '{"cep": "01310100"}' | jq '.' || echo "Service not available"
	@echo ""
	@echo "3. Invalid CEP (should return 422):"
	@curl -s -X POST http://localhost:8081/ \
	  -H "Content-Type: application/json" \
	  -d '{"cep": "123"}' | jq '.' || echo "Service not available"
	@echo ""

restart: down up
	@echo "✓ Services restarted!"
