

build:
 go build -o exchanger ./cmd/main.go
 
tests:
   go test ./...

usage:
   ./exchanger 12223.45 RUB USD
   
output:
   12223.45 RUB equals 156.99 USD



  
