clean:
	rm bin/sql-client*.*
build:
lambda:
	env GOOS=linux go build -ldflags="-s -w" -o bin/sql-client main.go
linux:
	env GOOS=linux go build -ldflags="-s -w" -o bin/sql-client main.go
linux-arm:
	env GOOS=linux GOARCH=arm go build -ldflags="-s -w" -o bin/sql-client main.go
local:
	go build -ldflags="-s -w" -o bin/sql-client main.go
mac:
	env GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o bin/sql-client main.go
windows:
	env GOOS=windows go build -ldflags="-s -w" -o bin/sql-client.exe main.go