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
	@if [ -z "$$(docker images -q mysql:9.1.0)" ]; then \
		docker pull mysql:9.1.0; \
	fi

	@if [ $$(docker ps -aq -f name=mysql-container) ]; then \
		docker start mysql-container; \
	else \
		docker run --name mysql-container -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=bats -p 3306:3306 -d mysql:9.1.0; \
	fi