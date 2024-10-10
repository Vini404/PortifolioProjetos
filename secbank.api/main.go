package main

import (
	"net/http"
)

func main() {
	http.ListenAndServe(":8080", ChiRouter().InitRouter())
}

//To run database in docker local, run this command:

//docker run -p 5432:5432 --name postgres-db -e POSTGRES_PASSWORD=1234 -d postgres

//TODO
//-Implementar sistema de log com gerenciamento de erro eficiente
//padronizar response em formato json de forma gennerica, consultar as boas praticas de retorno
