.PHONY: todone
todone:
	@go run cmd/todone/main.go

install
	go install ./cmd/todone/...