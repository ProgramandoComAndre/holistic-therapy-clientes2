package repositories

import (
	"github.com/ProgramandoComAndre/holistic-therapy-clientes2/src/domain/entities"
)
type ClientRepository interface {
	CreateClient(*entities.Client) (*entities.Client, error)
	GetClientById(string) (*entities.Client, error)
	GetClients(int, int) (*entities.PaginatedClients, error)
	UpdateClient(string, entities.UpdateClientRequest) (*entities.Client, error)
	DeleteClient(string) error
}
