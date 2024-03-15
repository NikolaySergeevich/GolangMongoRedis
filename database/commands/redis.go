package commands

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"

	"redis-cache/database"
)

var _ Repository = (*redisRepository)(nil) //это строка значт соответсвие интерфейсу

func NewCacheRepository(db *redis.Client, repository Repository) Repository {
	return &redisRepository{db: db, dbLayer: repository}
}

type redisRepository struct {
	db      *redis.Client
	dbLayer Repository
}

func (r redisRepository) AddCommand(ctx context.Context, command database.Command) error {
	if err := r.dbLayer.AddCommand(ctx, command); err != nil {
		return fmt.Errorf("db layer AddCommand: %w", err)
	}
	return nil
}

func (r redisRepository) FindByCommand(ctx context.Context, command string) (database.Command, error) {
	var cmd database.Command

	bytes, err := r.db.Get(ctx, command).Bytes()
	if err != nil{
		if err == redis.Nil{
			// тут если в redis не нашли команду
			byCommand, err := r.dbLayer.FindByCommand(ctx, command)
			if err != nil{
				return cmd, fmt.Errorf("db Layer Find: %w", err)
			}

			encoder, err := json.Marshal(byCommand)
			if err != nil{
				return cmd, fmt.Errorf("json Marshal: %w", err)
			}
			status := r.db.Set(ctx, command, encoder, -1)// -1 это время жизни записи в кэше. Бесконечно должго в этом случае
			if err := status.Err(); err != nil{
				return cmd, fmt.Errorf("redis Err: %w", err)
			}
			//и после анесения в redis значения, можно вернуть искомую команду
			return byCommand, nil

		}
			
		return cmd, fmt.Errorf("redis Get: %w", err)
	}

	if err := json.Unmarshal(bytes, &cmd); err != nil{
		return cmd, fmt.Errorf("json Unmarshal: %w", err)
	}
	 return cmd, nil
}
