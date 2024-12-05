# lambda_autenticator

Compile o código Go
Compile o código para a arquitetura Linux e AMD64 (ou ARM64, dependendo da configuração do Lambda) porque o runtime do AWS Lambda usa um ambiente Linux.

```ssh
GOOS=linux GOARCH=amd64 go build -o bootstrap main.go
```

Empacote o binário
Coloque o arquivo executável em um arquivo zip
```ssh
zip bootstrap.zip bootstrap
```