boiler:
	sqlboiler psql
vet:
	go vet ./...
linter:
	golangci-lint run
tidy:
	go mod tidy && go mod vendor
update:
	go get -u ./...
