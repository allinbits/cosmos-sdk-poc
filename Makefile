.PHONY: proto

proto:
	docker build -t dev:proto-build -f contrib/devc/proto.dockerfile .
	docker run -v "$(CURDIR):/genproto" -w /genproto dev:proto-build ./scripts/genproto.sh

test:
	go test ./...

