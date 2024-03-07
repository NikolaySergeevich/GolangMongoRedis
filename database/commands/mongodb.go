package commands

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"redis-cache/database"
)

var _ Repository = (*mongoDbRepository)(nil)

const collection = "commands"

func NewMongoDbRepository(db *mongo.Database) Repository {
	return &mongoDbRepository{db: db}
}

type mongoDbRepository struct {
	db *mongo.Database
}

func (m mongoDbRepository) AddCommand(ctx context.Context, command database.Command) error {
	// TODO implement me
	panic("implement me")
}

func (m mongoDbRepository) FindByCommand(ctx context.Context, command string) (database.Command, error) {
	// TODO implement me
	panic("implement me")
}
