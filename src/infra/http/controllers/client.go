package controllers

import(
	"net/http"
	"encoding/json"
	"github.com/ProgramandoComAndre/holistic-therapy-clientes2/src/domain/services"
	"github.com/ProgramandoComAndre/holistic-therapy-clientes2/src/domain/entities"
	"strconv"
)

type ClientController struct {
	CreateClientFileCommand *services.CreateClientFileCommand
	ListClientsQuery *services.ListClientsQuery
}

type errorBody struct {
	Message string `json:"message"`
}

func NewClientController(CreateClientFileCommand *services.CreateClientFileCommand, ListClientsQuery *services.ListClientsQuery) *ClientController {
	return &ClientController{
		CreateClientFileCommand: CreateClientFileCommand,
		ListClientsQuery: ListClientsQuery,
	}
}

func (cc *ClientController) CreateClient(w http.ResponseWriter, r *http.Request) {
	
	authorizedUser, ok := r.Context().Value("AuthorizedUser").(*entities.AuthorizedUser)

	if !ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&errorBody{Message: "Unauthorized"})
		return
	}

	var createClientRequest entities.CreateClientRequest
	err := json.NewDecoder(r.Body).Decode(&createClientRequest)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&errorBody{Message: err.Error()})
		return
	}



	err = createClientRequest.Validate()

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&errorBody{Message: err.Error()})
		return
	}


	client, err := cc.CreateClientFileCommand.Execute(*authorizedUser,createClientRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(client)
}

func (cc *ClientController) GetClients(w http.ResponseWriter, r *http.Request) {

	_, ok := r.Context().Value("AuthorizedUser").(*entities.AuthorizedUser)
	if !ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&errorBody{Message: "Unauthorized"})
		return
	}
	
	queryParams := r.URL.Query()
	page := queryParams.Get("page")
	limit := queryParams.Get("limit")

	pageConverted, err := strconv.Atoi(page)
	
	if err != nil {
		pageConverted = 1
	}
	limitConverted, err := strconv.Atoi(limit)
	if err != nil {
		limitConverted = 5
	}

	
	
	client, err := cc.ListClientsQuery.Execute(pageConverted,limitConverted)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(client)
}
