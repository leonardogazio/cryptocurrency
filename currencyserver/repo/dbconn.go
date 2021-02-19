package repo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

//Database stores MongoDB connection
var Database *mongo.Database

//Creates database indexes
func createIndexes() (err error) {
	collation := options.Collation{
		Locale:   "en",
		Strength: 2,
	}

	models := []mongo.IndexModel{
		{
			Keys:    bsonx.Doc{{Key: "code", Value: bsonx.Int32(1)}},
			Options: options.Index().SetUnique(true).SetCollation(&collation),
		},
		{
			Keys:    bsonx.Doc{{Key: "name", Value: bsonx.Int32(1)}},
			Options: options.Index().SetUnique(true).SetCollation(&collation),
		},
	}

	opts := options.CreateIndexes().SetMaxTime(10 * time.Second)
	_, err = Database.Collection("currencies").Indexes().CreateMany(context.Background(), models, opts)
	return err
}

//Open connects to MongoDB database.
func Open() error {
	ctx := context.Background()
	db, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err = db.Ping(ctx, nil); err != nil {
		return err
	}
	Database = db.Database("cryptocurrency")

	if err := createIndexes(); err != nil {
		return err
	}

	return nil
}
