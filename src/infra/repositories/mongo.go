package repositories

import (
	"context"
	"errors"

	"github.com/ProgramandoComAndre/holistic-therapy-clientes2/src/domain/entities"
	r "github.com/ProgramandoComAndre/holistic-therapy-clientes2/src/domain/repositories"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClientRepository struct {
	clientCollection *mongo.Collection
}

func NewClienteRepository(database *mongo.Database) r.ClientRepository {
	return &MongoClientRepository{
		clientCollection: database.Collection("clients"),
	}
}

func (r *MongoClientRepository) CreateClient(client *entities.Client) (*entities.Client, error) {
	result, err := r.clientCollection.InsertOne(context.Background(), client)
	if err != nil {
		return nil, err
	}
	newClient := entities.Client{
		Name:        client.Name,
		Birthdate:   client.Birthdate,
		Email:       client.Email,
		Mobilephone: client.Mobilephone,
		Address:     client.Address,
		Diseases:    client.Diseases,
		OtherInfo:   client.OtherInfo,
		TherapistsAccess: client.TherapistsAccess,
	}


	newClient.ID = result.InsertedID.(primitive.ObjectID).Hex() //result.InsertedID.(string)

	return &newClient, nil
}

func (r *MongoClientRepository) GetClientById(id string) (*entities.Client, error) {
	var client entities.Client
	err := r.clientCollection.FindOne(context.Background(), bson.D{{"_id", id}}).Decode(&client)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (r *MongoClientRepository) GetClients(limit int, page int, username string) (*entities.PaginatedClients, error) {
	var clients []entities.Client
	clients = make([]entities.Client, 0)
	l := int64(limit)
	skip := int64(page * limit - limit)
	fOpt := options.FindOptions { Limit: &l, Skip: &skip }

	filter := bson.M{"therapistsaccess": bson.M{"$in": []string{username}}}


	cursor, err := r.clientCollection.Find(context.Background(), filter, &fOpt)
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.Background()) {
		var client entities.Client
		err := cursor.Decode(&client)
		if err != nil {
			return nil, err
		}
		clients = append(clients, client)
	}
	return &entities.PaginatedClients{Clients: clients, Page: page, Limit: limit}, nil
}

func (r *MongoClientRepository) UpdateClient(id string, client entities.UpdateClientRequest) (*entities.Client, error) {
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{
		{"name", client.Name},
		{"birthdate", client.Birthdate},
		{"email", client.Email},
		{"mobilephone", client.Mobilephone},
		{"address", client.Address},
		{"diseases", client.Diseases},
		{"otherinfo", client.OtherInfo},
	}}}
	result, err := r.clientCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}
	if result.MatchedCount == 0 {
		return nil, errors.New("client not found")
	}
	return nil, nil
}

func (r *MongoClientRepository) DeleteClient(id string) error {
	_, err := r.clientCollection.DeleteOne(context.Background(), bson.D{{"_id", id}})
	if err != nil {
		return err
	}
	return nil
}
