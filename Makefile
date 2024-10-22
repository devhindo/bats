.PHONY: ui
.PHONY: server
ui:
	@cd web && bun run dev -- --open

server:
	@cd server && air
