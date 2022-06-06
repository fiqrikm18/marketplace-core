dev:
	go run cmd/main.go

cert:
	mkdir certs && rm -rf certs/* && openssl genrsa -out certs/app.rsa 1024 && openssl rsa -in certs/app.rsa -pubout > certs/app.rsa.pub