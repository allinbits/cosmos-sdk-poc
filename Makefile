.PHONY: proto testapp

proto:
	docker build -t dev:proto-build -f contrib/devc/proto.dockerfile .
	docker run -v "$(CURDIR):/genproto" -w /genproto dev:proto-build ./scripts/genproto.sh

test:
	go test ./...

testapp:
	rm -rf testapp/app/config/data/blockstore.db
	rm -rf testapp/app/config/data/evidence.db
	rm -rf testapp/app/config/data/state.db
	rm -rf testapp/app/config/data/tx_index.db
	echo '{}' > testapp/app/config/data/priv_validator_state.json
	go build -o testapp/testapp.exe testapp/main.go
	testapp/testapp.exe