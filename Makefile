.PHONY: ui
.PHONY: server
ui:
	@cd web && bun run dev -- --open

server:
	@if sudo lsof -i :8080 >/dev/null; then \
		echo "Port 8080 is already in use"; \
		sudo lsof -t -i :8080 | xargs sudo kill && echo "Port 8080 is now free"; \
	fi
	
	@cd server && air
