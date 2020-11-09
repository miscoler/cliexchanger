

build:
 go build ./cmd/main.go
 
tests:
   go build -o exchanger test ./...

usage:
   ./exchanger 12223.45 RUB USD
  
