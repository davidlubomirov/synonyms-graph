run-dev-apiserver:
	go run cmd/apiserver/apiserver.go

run-all-tests:
	go clean --testcache && go test -v ./...