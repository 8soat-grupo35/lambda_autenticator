package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

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
var	getUserAPI = "https://fastfood/v1/customer/cpf?cpf="
var	postUserAPI = "https://fastfood/v1/customer"

func handler(event InputEvent) (string, error) {
	cpf := event.CPF
	name := event.Name
	email := event.Email

	// Verificar se o usuário já existe (GET)
	shouldReturn, returnValue, returnValue1 := buscarCliente(cpf)
	if shouldReturn {
		return returnValue, returnValue1
	}

	// Usuário não encontrado, cadastrar (POST)
	return cadastrarCliente(name, email, cpf)
}

func cadastrarCliente(name string, email string, cpf string) (string, error) {
	newUser := CustomerDto{
		Name:  name,
		Email: email,
		CPF:   cpf,
	}

	jsonData, err := json.Marshal(newUser)
	if err != nil {
		return "", fmt.Errorf("erro ao criar payload de cadastro: %v", err)
	}

	postResp, err := http.Post(postUserAPI, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("erro ao cadastrar usuário: %v", err)
	}
	defer postResp.Body.Close()

	if postResp.StatusCode != http.StatusCreated {
		body, _ := ioutil.ReadAll(postResp.Body)
		return "", fmt.Errorf("erro ao cadastrar usuário: %s", string(body))
	}

	return "Usuário cadastrado com sucesso!", nil
}

func buscarCliente(cpf string) (bool, string, error) {
	resp, err := http.Get(getUserAPI+""+cpf)
	if err != nil {
		return true, "", fmt.Errorf("erro ao verificar usuário: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return true, "Usuário já cadastrado!", nil
	}

	if resp.StatusCode != http.StatusNotFound {
		body, _ := ioutil.ReadAll(resp.Body)
		return true, "", fmt.Errorf("erro inesperado ao verificar usuário: %s", string(body))
	}
	return false, "", nil
}

func main() {
	lambda.Start(handler)
}
