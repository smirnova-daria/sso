.PHONY: sso

sso:
	go run ./cmd/sso/main.go --config=./config/local.yaml
