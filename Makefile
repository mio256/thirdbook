BIN_DIR:=$(shell pwd)/bin

.PHONY: test
test:
	go test -p=1 ./...

.PHONY: ogen
ogen:
	ogen -target ui/api -clean docs/openapi.yaml

.PHONY: tools
tools:
	go get .
	GOBIN=$(BIN_DIR) go install github.com/sqldef/sqldef/cmd/psqldef@$(shell go list -m -f "{{.Version}}" github.com/sqldef/sqldef)
	GOBIN=$(BIN_DIR) go install github.com/sqlc-dev/sqlc/cmd/sqlc@$(shell go list -m -f "{{.Version}}" github.com/sqlc-dev/sqlc)

.PHONY: migrate
migrate:
	$(BIN_DIR)/psqldef -U postgres -W postgres -p 5432 -f ./db/core.sql --enable-drop-table thirdbook

.PHONY: sqlc
sqlc:
	$(BIN_DIR)/sqlc generate