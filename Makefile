.PHONY: test
test:
	go test -p=1 ./...

.PHONY: ogen
ogen:
	ogen -target ui/api -clean docs/openapi.yaml
