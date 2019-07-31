b_s:
	go build -o bin/contacterServer contacterServer.go
b_s_l:
	GOOS=linux go build -o bin/contacterServer_linux contacterServer.go
b_c:
	go build -o bin/client client/client.go