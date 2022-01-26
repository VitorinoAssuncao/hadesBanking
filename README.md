Esta pequena api foi criada conforme os requisitos do desafio,com o objetivo de permitir a criação de contas financeiras e as transferencias entre elas. 

Para tal foram selecionadas as seguintes técnologias:

* [Go](https://golang.org) Uma linguagem robusta e rápida, ela foi utilizada como base para o desenvolvimento da aplicação, por sua escalabilidade e capacidade de processamento de dados.

* [Banco de Dados Postgresql](https://www.postgresql.org) É um banco de dados rápido, leve e eficiente, atendendo assim todos os requisitos necessários.

# Requisitos para rodar o projeto de forma Local:

Para rodar o projeto de forma local é necessário se possuir ao menos o Go instalado e configurado.

Atualmente o banco utilizado é uma base PostgreSQL disponível através de imagem Docker, porém caso se deseje configurar uma de forma local, basta se instalar a base apropriada e realizar as adequações em seu conector.

Para se baixar o arquivo diretamente do git pode se utilizar o comando abaixo:

 ``` git clone https://github.com/VitorinoAssuncao/stoneBanking.git ```

GitHub CLI
 
 ``` gh repo clone VitorinoAssuncao/stoneBanking ```


# Gerando Imagem Docker:

Para se gerar a imagem docker, após clonado/baixado o programa, basta acessar o terminal de sua prefêrencia, e rodar o comando abaixo:

 ```  docker-compose up ```

Dessa forma o sistema irá gerar a imagem docker já com os dados de banco criados.

# Rodando o Programa:

Após ter a imagem docker configurada, ou possuir um banco de dados local configurado de forma apropriada, basta abrir no se programa de preferencia o comando abaixo:

``` go build ```

# Estruturas Relevantes:

Este projeto consiste em uma aplicação de backend, a qual não possui uma rota raiz (/) atualmente, possuindo apenas 3 estruturas de rotas, conforme a necessidade do usuário:

• accounts: Referente aos dados de conta, gerais e individuais. E a partir do ID do usuário que será possível acessar o saldo da conta através das rotas (/accounts/balancce) e uma listagem de contas gerais (/accounts). Além disso, é possível se fazer nessa rota o login do usuário em questão (/accounts/login)

• transfers: Rota responsável pelas ações de transfêrencia entre duas contas, só sendo permitido realizar as mesmas quando realizado previamente um login. Deverá obrigatoriamente enviar um token no header das requisições.

# EndPoints

Segue abaixo rotas principais liberadas atualmente no projeto:

## Accounts

### Request  
` "GET /accounts" : Rota que retorna a listagem de contas cadastradas.`

### Response  
`[
  {
		"id": "abc7cf9f-b984-4d0c-9dfa-8c89eaf1bfb0",
		"name": "Joselino das Neves",
		"cpf": "74202445023",
		"balance": 99.99,
		"created_at": "2022-01-18T17:28:53Z"
	},
  {
	"id": "7b17f816-ef30-4096-95ef-c6118068ade1",
	"name": "João",
	"cpf": "34760400036",
	"balance": 99.99,
	"created_at": "2022-01-21T13:12:41Z"
  }
]`
 
---

### Request
` "POST /account" : Rota para a criação do usuário`

- Body (JSON)
`{
	"name":"João",
	"cpf":"347.604.000-36",
	"secret":"12344",
	"balance":9999
}`
---
### Response
`  {
	"id": "7b17f816-ef30-4096-95ef-c6118068ade1",
	"name": "João",
	"cpf": "34760400036",
	"balance": 99.99,
	"created_at": "2022-01-21T13:12:41Z"
}`


## Balance
### Request
` "GET /account/balance" : Rota validar o saldo atual da conta, deve-se passar o token de acesso, recebido ao logar`

- Header 
`{
	Authorization: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzQ4OTI0NzIsImlzcyI6InZpdG9yaW5vIiwiVXNlcl9pZCI6MX0.dlyrFzbfBz7QPBQOaq9c1_gCVmv2JcjkI0SGWZ6ZsVU"
}`


### Response
`{
  "balance": 10.00
}`

### Request
` "POST /account/login" : Rota para realizar o login na conta`

- Body (JSON)
`{
	"cpf": "57857751099",
	"secret":"12345"
}`

### Response
`{
	"authorization": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiJjMDM2NDc1Zi1iN2EwLTRmMzQtOGYxZi1jNDM1MTVkMzE3MjQifQ.Vzl8gI6gYbDMTDPhq878f_Wq_b8J0xz81do8XmHeIFI"
}`

## Transfer

### Request
` "GET /transfers" : Rota para retornar os dados de todas as transfêrencias (recebidas ou realizadas) do usuário`

- Header 
`{
	Authorization: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzQ4OTI0NzIsImlzcyI6InZpdG9yaW5vIiwiVXNlcl9pZCI6MX0.dlyrFzbfBz7QPBQOaq9c1_gCVmv2JcjkI0SGWZ6ZsVU"
}`

### Response  
`[
	{
		"id": "5ccbb3f0-9351-45c2-80cb-cfffbee767c2",
		"account_origin_id": "c036475f-b7a0-4f34-8f1f-c43515d31724",
		"account_destiny_id": "4cc7fe98-9996-408c-bff7-06cee3e6c519",
		"value": 0.01,
		"created_at": "2022-01-10T11:46:37Z"
	},
  	{
		"id": "5ccbb3f0-9351-45c2-80cb-cfffbee767c2",
		"account_origin_id": "c036475f-b7a0-4f34-8f1f-c43515d31724",
		"account_destiny_id": "4cc7fe98-9996-408c-bff7-06cee3e6c519",
		"value": 0.01,
		"created_at": "2022-01-10T11:46:37Z"
	}
]`

---

### Request
` "POST /transfer" : Cria uma nova transfêrencia do usuário informado no token, para o destinatário.`

- Header 
`{
	Authorization: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzQ4OTI0NzIsImlzcyI6InZpdG9yaW5vIiwiVXNlcl9pZCI6MX0.dlyrFzbfBz7QPBQOaq9c1_gCVmv2JcjkI0SGWZ6ZsVU"
}`

- Body (JSON):
`{
	"account_destiny_id":"4cc7fe98-9996-408c-bff7-06cee3e6c519",
	"amount":1
}`

### Response

`{
	"id": "5ccbb3f0-9351-45c2-80cb-cfffbee767c2",
	"account_origin_name": "c036475f-b7a0-4f34-8f1f-c43515d31724",
	"account_destiny_name": "4cc7fe98-9996-408c-bff7-06cee3e6c519",
	"value": 0.01,
	"created_at": "2022-01-10T11:46:37Z"
}`
