package repo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Database stores MongoDB connection
var Database *mongo.Database

//Open connects to MongDB database.
func Open() error {
	ctx := context.Background()
	db, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err = db.Ping(ctx, nil); err != nil {
		return err
	}
	Database = db.Database("cryptocurrency")
	return nil
}
