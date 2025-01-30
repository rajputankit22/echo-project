package mongodb

// MongoDBInterface - interface for MongoDB
type MongoDBInterface interface {
	Disconnect() error
	FindOne(collection string, filter interface{}) (interface{}, error)
	InsertOne(collection string, document interface{}) error
}
