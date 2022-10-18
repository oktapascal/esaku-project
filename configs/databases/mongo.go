package databases

import (
	"context"
	"esaku-project/configs"
	"esaku-project/helpers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
	"time"
)

func NewMongoDatabase(configuration configs.Config) *mongo.Database {
	ctx, cancel := NewMongoContext()
	defer cancel()

	mongoPoolMin, err := strconv.Atoi(configuration.Get("DB_MONGO_POOL_MIN"))
	helpers.PanicIfError(err)

	mongoPoolMax, err := strconv.Atoi(configuration.Get("DB_MONGO_POOL_MAX"))
	helpers.PanicIfError(err)

	mongoMaxIdleTime, err := strconv.Atoi(configuration.Get("DB_MONGO_MAX_IDLE_TIME_SECOND"))
	helpers.PanicIfError(err)

	option := options.Client().ApplyURI(configuration.Get("DB_MONGO_URI")).
		SetMinPoolSize(uint64(mongoPoolMin)).
		SetMaxPoolSize(uint64(mongoPoolMax)).
		SetMaxConnIdleTime(time.Duration(mongoMaxIdleTime) * time.Second)

	client, err := mongo.NewClient(option)
	helpers.PanicIfError(err)

	err = client.Connect(ctx)
	helpers.PanicIfError(err)

	database := client.Database(configuration.Get("DB_MONGO_DATABASE"))

	return database
}

func NewMongoContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
