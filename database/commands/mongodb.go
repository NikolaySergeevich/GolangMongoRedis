package commands

import (
	"go.mongodb.org/mongo-driver/mongo"

	"redis-cache/database"
)

var _ Repository = (*mongoDbRepository)(nil)

func NewMongoDbRepository(db *mongo.Client) Repository {
	return &mongoDbRepository{db: db}
}

type mongoDbRepository struct {
	db *mongo.Client
}

func (m mongoDbRepository) AddCommand(command database.Command) error {
	// TODO implement me
	panic("implement me")
}

func (m mongoDbRepository) FindByCommand() (database.Command, error) {
	// TODO implement me
	panic("implement me")
}
