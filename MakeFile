.PHONY: run
run:
	@go build -o run && ./run

install:
	@cd app; npm i

runweb:
	@cd app; npm run start

buildapp:
	@cd app; npm run build
