dev:
	go run cmd/main.go

cert:
	rm -rf certs/cert*
	openssl genrsa -out certs/cert 2014
	openssl rsa -in certs/cert -pubout > certs/cert.pub