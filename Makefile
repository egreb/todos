generate:
	rm -rf generated && mkdir generated
	oto -template ./templates/server.go.plush \
			-out ./generated/oto.gen.go \
			-ignore Ignorer \
			-pkg generated \
			./definitions
	gofmt -w ./generated/oto.gen.go ./generated/oto.gen.go
	oto -template ./templates/client.ts.plush \
			-out ./src/types/oto.gen.ts \
			-ignore Ignorer \
			./definitions
download:
	@echo Download go.mod dependencies
	@go mod download

install-tools: download
	@echo Installing tools from tools.go
	@cat tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

test:
	go test ./... -v