test:
	go test -v ./...

check:
	go fmt ./... && goimports -w ./ && go mod tidy
	golint -set_exit_status ./...
	go vet ./...