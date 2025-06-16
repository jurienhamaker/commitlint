tests:
	@go test ./... | grep -v "no test files"

build:
	@go build -ldflags="-s -w -X=main.Version=1.0.0" -o dist/commitlint ./cmd/commitlint/main.go 

build-windoos:
	@GOOS=windows GOARCH=amd64 go build -ldflags="-s -w -X=main.Version=1.0.0" -o dist/commitlint.exe ./cmd/commitlint/main.go 
