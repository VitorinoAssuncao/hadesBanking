# Desafio da Empresa Stone.

Esta pequena api foi criada conforme os requisitos do [desafio stone](https://gist.github.com/guilhermebr/fb0d5896d76634703d385a4c68b730d8),com o objetivo de permitir a criação de contas financeiras e as transferencias entre elas. 

Para tal foram selecionadas as seguintes técnologias:

* [Go](https://golang.org) Uma linguagem robusta e rápida, ela foi utilizada como base para o desenvolvimento da aplicação, por sua escalabilidade e capacidade de processamento de dados.

* [ORM - Gorm](https://gorm.io/index.html) Foi utilizado o ORM Gorm, por sua simplicidade e eficiência, pois ao abstrair as consultas nos permite maior customização e possíveis alterações futuras fáceis no banco de dados.

* [Banco de Dados Postgresql](https://www.postgresql.org) Como dito anteriormente, com o ORM a seleção pelo banco de dados se tornou secundário, neste caso selecionei o postgresql por alinhamento de usos internos.

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
`{
    "Account_id": 1,
    "Account_cpf": "383333333",
    "Account_name": "Joao",
    "Created_at": "2021-10-23T00:24:16.902971-03:00"
 }`
 
---

### Request
` "POST /account" : Rota para a criação do usuário`

- Body (JSON)
`{
	"name": "Vitorino",
	"cpf": "57857751011",
	"secret":"12345",
	"balance": 1000
}`
---
### Response
`{
  "ID": 3,
  "Name": "Vitorino",
  "Cpf": "57857751011",
  "Secret": "$2a$14$W2wx0ynuJa9wRA9CX65VL./nPgtmgMD.0Mmzz5YsZIIbPNJipYam6",
  "Balance": 1000,
  "Created_at": "2021-10-22T14:47:01.064005-03:00"
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
  "balance": 1000
}`

### Request
` "POST /login" : Rota para realizar o login na conta`

- Body (JSON)
`{
	"cpf": "57857751099",
	"secret":"12345"
}`

### Response
`{
  "accountToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzQ4OTI0NzIsImlzcyI6InZpdG9yaW5vIiwiVXNlcl9pZCI6MX0.dlyrFzbfBz7QPBQOaq9c1_gCVmv2JcjkI0SGWZ6ZsVU"
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
    "id": 1,
    "acount_origin_id": 10,
    "acount_destination_id": 1,
    "amount": 100,
    "created_at": "2021-10-15T21:10:45-03:00"
  },
  {
    "id": 2,
    "acount_origin_id": 10,
    "acount_destination_id": 1,
    "amount": 100,
    "created_at": "2021-10-15T21:11:58-03:00"
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
	"acount_destination_id":2,
	"amount": 100
}`

### Response

`{
  "id": 1,
  "acount_origin_id": 1,
  "acount_destination_id": 2,
  "amount": 100,
  "created_at": "2021-10-22T03:48:24.903575-03:00"
}`
