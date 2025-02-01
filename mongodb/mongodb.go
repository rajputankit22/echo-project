package mongodb

import (
	"context"
	"echo-project/config"
	"echo-project/logger"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ctx              = context.Background()
	newMongoDBClient = newMongoDB
)

// MongoDBInterface - interface for MongoDB
type MongoDBInterface interface {
	Disconnect() error
	FindOne(collection string, filter interface{}) (interface{}, error)
	InsertOne(collection string, document interface{}) error
}

// MongoDB - struct for MongoDB
type mongoDB struct {
	client *mongo.Client
}

// Disconnect - disconnects the MongoDB client
func (m *mongoDB) Disconnect() error {
	return m.client.Disconnect(ctx)
}

// Constructor function for MongoDB
func NewMongoDBAdapter() (MongoDBInterface, error) {
	return &mongoDB{
		client: newMongoDBClient(),
	}, nil
}

// FindOne - find one document in the collection
func (m *mongoDB) FindOne(collection string, filter interface{}) (interface{}, error) {
	return nil, nil
}

// InsertOne - insert one document in the collection
func (m *mongoDB) InsertOne(collection string, document interface{}) error {
	return nil
}

// newMongoDB - create a new MongoDB client
func newMongoDB() *mongo.Client {
	// opts := options.Client().ApplyURI(config.Config().Mongo.URI)
	opts := options.Client().ApplyURI(config.Config().Mongo.URI)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		logger.Error("", "Failed to connect to MongoDB:", err)
	} else {
		logger.Trace("Successfully connected to MongoDB")
	}

	// Get a reference to the database "firstdb"
	dbName := config.Config().Mongo.DBName
	db := client.Database(dbName)
	collectionName := config.Config().Mongo.CollectionName

	// Get a reference to a collection
	collection := db.Collection(collectionName)
	logger.Trace("Successfully created collection: " + collection.Name())
	if collection.Name() == collectionName {
		createUserIndexes(collection)
	}
	return client
}

// Create indexes for the "users" collection
func createUserIndexes(collection *mongo.Collection) {
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"email": 1},              // 1 means ascending order
		Options: options.Index().SetUnique(true), // Ensuring unique emails
	}

	indexName, err := collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		logger.Error("", "Failed to create index:", err)
	}
	logger.Info("", "Index created", indexName)
}
