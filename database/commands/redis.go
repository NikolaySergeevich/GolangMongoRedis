package commands

import (
	"github.com/go-redis/redis/v8"

	"redis-cache/database"
)

var _ Repository = (*redisRepository)(nil)

func NewRedisRepository(db *redis.Client) Repository {
	return &redisRepository{db: db}
}

type redisRepository struct {
	db *redis.Client
}

func (r redisRepository) AddCommand(command database.Command) error {
	// TODO implement me
	panic("implement me")
}

func (r redisRepository) FindByCommand() (database.Command, error) {
	// TODO implement me
	panic("implement me")
}
