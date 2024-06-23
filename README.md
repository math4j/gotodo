# gotodo

### Run project
```sh
go run cmd/main.go
```

### Run tests
```sh
go test -v ./... -coverprofile=coverage.out
```

### See test results
```sh
go tool cover -html=coverage.out
```