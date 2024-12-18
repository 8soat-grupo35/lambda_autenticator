package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var fastfoodAPI = os.Getenv("FASTFOOD_APP_URL")

// URL das APIs
var getUserAPI = fastfoodAPI + "/v1/customer/cpf/"
var postUserAPI = fastfoodAPI + "/v1/customer"

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

func main() {
	lambda.Start(handler)
}

func handler(request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	log.Println("Iniciando lambda-authenticator")
	log.Println("FASTFOOD_API_URL", fastfoodAPI)

	// Verificar se o body está vazio
	if request.Body == "" {
		log.Println("Erro: Body vazio na requisição")
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 400,
			Body:       "Body vazio na requisição",
		}, nil
	}

	log.Printf("Body recebido: %s", request.Body)

	// Decodificar o JSON recebido
	var input InputEvent
	err := json.Unmarshal([]byte(request.Body), &input)
	if err != nil {
		log.Printf("Erro ao decodificar o body: %v", err)
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 400,
			Body:       "Formato inválido no body",
		}, nil
	}

	// Acessar os valores decodificados
	log.Printf("Dados recebidos: CPF=%s, Name=%s, Email=%s", input.CPF, input.Name, input.Email)

	existeCliente, response, err := buscarCliente(input.CPF)
	if existeCliente {
		return response, err
	}

	return cadastrarCliente(input.CPF, input.Name, input.Email)
}

func cadastrarCliente(name string, email string, cpf string) (events.APIGatewayV2HTTPResponse, error) {
	newUser := CustomerDto{
		Name:  name,
		Email: email,
		CPF:   cpf,
	}

	jsonData, err := json.Marshal(newUser)
	if err != nil {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 400,
			Body:       "Erro ao criar payload de cadastro",
		}, err
	}

	postResp, err := http.Post(postUserAPI, "application/json", bytes.NewBuffer(jsonData))
	log.Println("Retorno API Cadastrar Cliente")
	log.Println("Cadastrar Cliente Resp",postResp)
	
	if err != nil {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 400,
			Body:       "Erro ao cadastrar usuário",
		}, err
	}
	defer postResp.Body.Close()

	if postResp.StatusCode != http.StatusCreated {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 400,
			Body:       "Erro ao cadastrar usuário:",
		}, err
	}

	return events.APIGatewayV2HTTPResponse{
		StatusCode: 200,
		Body:       "Usuário cadastrado com sucesso!",
	}, nil
}

func buscarCliente(cpf string) (bool, events.APIGatewayV2HTTPResponse, error) {
	resp, err := http.Get(getUserAPI + "" + cpf)
	log.Println("Retorno API Buscar Cliente")
	log.Println("Buscar Cliente", resp)
	if err != nil {
		return true, events.APIGatewayV2HTTPResponse{
			StatusCode: 500,
			Body:       "Erro ao verificar usuário.",
		}, err

	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		//return true, "Usuário já cadastrado!", nil

		return true, events.APIGatewayV2HTTPResponse{
			StatusCode: 422,
			Body:       "Usuário Autenticado com sucesso!",
		}, err
	}

	if resp.StatusCode != http.StatusInternalServerError {
		return true, events.APIGatewayV2HTTPResponse{
			StatusCode: 422,
			Body:       "Erro inesperado ao verificar usuário.",
		}, err
	}
	return false, events.APIGatewayV2HTTPResponse{}, nil
}
