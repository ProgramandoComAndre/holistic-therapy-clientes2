package services

import(
	"github.com/ProgramandoComAndre/holistic-therapy-clientes2/src/domain/repositories"
    "github.com/ProgramandoComAndre/holistic-therapy-clientes2/src/domain/entities")

type CreateClientFileCommand struct {
	clientRepository repositories.ClientRepository
}

func NewCreateClientFileCommand(clientRepository repositories.ClientRepository) *CreateClientFileCommand {
	return &CreateClientFileCommand{
		clientRepository: clientRepository,
	}
}

func (ccf *CreateClientFileCommand) Execute(authUser entities.AuthorizedUser, createClientRequest entities.CreateClientRequest) (*entities.Client, error) {
	repository := ccf.clientRepository
	newClient, err := authUser.CreateClient(createClientRequest)
	if err != nil {
		return nil, err
	}
	
	client, err := repository.CreateClient(newClient)
	if err != nil {
		return nil, err
	}
	return client, nil
}