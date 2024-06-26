package main

import (
	"context"
	 "net/http" // import gorilla/mux
	 "github.com/gorilla/mux"

	"github.com/ProgramandoComAndre/holistic-therapy-clientes2/src/infra/repositories"
	r "github.com/ProgramandoComAndre/holistic-therapy-clientes2/src/domain/repositories"
	"github.com/ProgramandoComAndre/holistic-therapy-clientes2/src/domain/services"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/ProgramandoComAndre/holistic-therapy-clientes2/src/infra/http/controllers"
	"github.com/ProgramandoComAndre/holistic-therapy-clientes2/src/infra/middlewares"
)

func main() {
	mongoURI:= "mongodb://localhost:27017"

	clientOptions := options.Client().ApplyURI(mongoURI)
	mongoConnection, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		panic(err)
	}
	clientsDatabase := mongoConnection.Database("clients")
	var clientRepository r.ClientRepository
	clientRepository = repositories.NewClienteRepository(clientsDatabase)
	
	createClientFileCommand := services.NewCreateClientFileCommand(clientRepository)
	listClientsQuery := services.NewListClientsQuery(clientRepository)
	clientsController := controllers.NewClientController(createClientFileCommand, listClientsQuery)

	r := mux.NewRouter()
	r.Use(middlewares.AuthMiddleware)
	r.HandleFunc("/clients", clientsController.CreateClient).Methods("POST")
	r.HandleFunc("/clients", clientsController.GetClients).Methods("GET")
	http.ListenAndServe(":8080", r)


}