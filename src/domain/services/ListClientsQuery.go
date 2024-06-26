package services

import(
	"github.com/ProgramandoComAndre/holistic-therapy-clientes2/src/domain/repositories"
    "github.com/ProgramandoComAndre/holistic-therapy-clientes2/src/domain/entities")

type ListClientsQuery struct {
	clientRepository repositories.ClientRepository
}

func NewListClientsQuery(clientRepository repositories.ClientRepository) *ListClientsQuery {
	return &ListClientsQuery{
		clientRepository: clientRepository,
	}
}

func (ccf *ListClientsQuery) Execute(page int, limit int) (*entities.PaginatedClients, error) {
	repository := ccf.clientRepository

	client, err := repository.GetClients(limit, page)
	if err != nil {
		return nil, err
	}
	return client, nil
}