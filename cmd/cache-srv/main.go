package main

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"redis-cache/database"
	"redis-cache/database/commands"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	mongoDB, err := mongo.Connect(
		ctx, &options.ClientOptions{
			Hosts: []string{fmt.Sprintf("%s:%d", "127.0.0.1", 27018)},
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	if err := mongoDB.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	redisClient := redis.NewClient(&redis.Options{DB: 1})

	commandsRepository := commands.NewMongoDbRepository(mongoDB.Database("test-db"))

	cachedRepository := commands.NewCacheRepository(redisClient)

	if err := cachedRepository.AddCommand(
		ctx, database.Command{
			ID:        uuid.New(),
			Command:   "ls -la",
			CreatedAt: time.Now().UTC(),
		},
	); err != nil {
		return
	}

	command, err := cachedRepository.FindByCommand(ctx, "ls -la")
	if err != nil {
		log.Fatal(err)
	}

	_, err = cachedRepository.FindByCommand(ctx, "ls -lat")
	if err != nil {
		log.Println(err)
	}

	fmt.Println(command.Command)

	_ = commandsRepository
	_ = cachedRepository
	_ = mongoDB

	<-ctx.Done()
}
