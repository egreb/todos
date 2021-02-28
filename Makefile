generate:
	rm -rf && mkdir generated
	oto -template ./templates/server.go.plush \
			-out ./generated/oto.gen.go \
			-ignore Ignorer \
			-pkg generated \
			./definitions
	gofmt -w ./generated/oto.gen.go ./generated/oto.gen.go
	oto -template ./templates/client.ts.plush \
			-out ./generated/oto.gen.ts \
			-ignore Ignorer \
			./definitions