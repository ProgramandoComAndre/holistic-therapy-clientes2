package main

import (
	"context"
	"net/http" // import gorilla/mux
	"os"
	r "github.com/ProgramandoComAndre/holistic-therapy-clientes2/src/domain/repositories"
	"github.com/ProgramandoComAndre/holistic-therapy-clientes2/src/domain/services"
	"github.com/ProgramandoComAndre/holistic-therapy-clientes2/src/infra/http/controllers"
	"github.com/ProgramandoComAndre/holistic-therapy-clientes2/src/infra/middlewares"
	"github.com/ProgramandoComAndre/holistic-therapy-clientes2/src/infra/repositories"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	err := godotenv.Load(".env")
 	
	mongoURI:= os.Getenv("MONGO_URI")

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
	handler := cors.AllowAll().Handler(r)
	http.ListenAndServe(":3004", handler)
	println("Listening on port 3004")

}

