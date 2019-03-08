cc := go build
dist := ../public/cgi-bin/index.cgi

main : main.go
	GOOS=linux GOARCH=amd64 $(cc) -o $(dist) main.go
