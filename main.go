package main

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type CustomerDto struct {
	Id    uint32 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	CPF   string `json:"cpf"`
}

type InputEvent struct {
	CPF   string `json:"cpf"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// URL das APIs
var getUserAPI = "https://fastfood/v1/customer/cpf?cpf="
var postUserAPI = "https://fastfood/v1/customer"

func main() {
	lambda.Start(handler)
}

func handler(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("Iniciando lambda-authenticator")

	// Verificar se o body está vazio
	log.Println("VALOR BODY"+event.Body)
	if event.Body == "" {
		log.Println("Erro: Body vazio na requisição")
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Body vazio na requisição",
		}, nil
	}

	log.Printf("Body recebido: %s", event.Body)

	// Decodificar o JSON recebido
	var input InputEvent
	err := json.Unmarshal([]byte(event.Body), &input)
	if err != nil {
		log.Printf("Erro ao decodificar o body: %v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Formato inválido no body",
		}, err
	}

	// Acessar os valores decodificados
	log.Printf("Dados recebidos: CPF=%s, Name=%s, Email=%s", input.CPF, input.Name, input.Email)

	// Criar resposta com status 200
	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       input.CPF,
	}

	return response, nil
}
