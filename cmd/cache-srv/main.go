package main

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"syscall"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

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

	commandsRepository := commands.NewMongoDbRepository(mongoDB)
	cachedRepository := commands.NewRedisRepository(redisClient)

	_ = commandsRepository
	_ = cachedRepository
	_ = mongoDB

	<-ctx.Done()
}
