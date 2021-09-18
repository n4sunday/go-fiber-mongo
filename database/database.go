package database

import(
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client * mongo.Client
	Db * mongo.Database
}

var DB MongoInstance

const dbName = "sundaydb"
const mongoURI = "mongodb://root:123456@localhost:27017/sundaydb?authSource=admin"

func Connect() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))

	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()

	err = client.Connect(ctx)
	db:= client.Database(dbName)

	if err != nil {
		return err
	}

	DB = MongoInstance{
		Client: client,
			Db: db,
	}

	return nil
}