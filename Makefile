server:
	go run main.go

test:
	go test -v -cover "$(path)"

mock:
	mockery --all --keeptree --disable-version-string
