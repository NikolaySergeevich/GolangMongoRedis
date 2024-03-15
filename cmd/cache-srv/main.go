package main

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-redis/redis/v8"
	// "github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"redis-cache/database"
	"redis-cache/database/commands"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	mongoClient, err := mongo.Connect(
		ctx, &options.ClientOptions{
			Hosts: []string{fmt.Sprintf("%s:%d", "127.0.0.1", 27018)},
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	if err := mongoClient.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	redisClient := redis.NewClient(&redis.Options{DB: 1, Addr: ":6381"})
	status := redisClient.Ping(ctx)
	if err := status.Err(); err != nil{
		log.Fatal(err)
	}


	mongoDb := mongoClient.Database("test-db")
	commandsRepository := commands.NewMongoDbRepository(mongoDb)

	cachedRepository := commands.NewCacheRepository(redisClient, commandsRepository)

	if err := cachedRepository.AddCommand(
		ctx, database.Command{
			ID:        primitive.NewObjectID(),//это для того что бы генерировать id
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
	_ = mongoClient

	<-ctx.Done()
}
