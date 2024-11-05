.PHONY: ui
.PHONY: server
.PHONY: db

ui:
	@cd web && bun run dev -- --open

server:
	@if sudo lsof -i :8080 >/dev/null; then \
		echo "Port 8080 is already in use"; \
		sudo lsof -t -i :8080 | xargs sudo kill && echo "Port 8080 is now free"; \
	fi
	
	@cd server && air

db:
	docker pull mysql:latest

	docker run --name mysql-container -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=testdb -p 3306:3306 -d mysql:latest