.PHONY: proto testapp

proto:
	docker build -t dev:proto-build -f contrib/devc/proto.dockerfile .
	docker run -v "$(CURDIR):/genproto" -w /genproto dev:proto-build ./scripts/genproto.sh

codegen:
	docker build -t starportcodegen:dev -f contrib/devc/codegen.dockerfile .
	docker run -v "$(CURDIR):/gencode" -w /gencode starportcodegen:dev ./scripts/gencode.sh

test:
	go test ./...

testapp:
	rm -rf testapp/app/config/data
	mkdir testapp/app/config/data
	echo '{}' > testapp/app/config/data/priv_validator_state.json
	go build -o testapp/testapp.exe testapp/main.go
	testapp/testapp.exe