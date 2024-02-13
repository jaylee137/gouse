package connection

// import (
// 	"context"

// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// 	"go.mongodb.org/mongo-driver/mongo/readpref"
// )

// type MongoConnection struct {
// 	client *mongo.Client
// 	ctx    context.Context
// }

// func NewMongo(context context.Context, uri string) (*MongoConnection, error) {
// 	client, err := Connect(uri)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &MongoConnection{
// 		client: client,
// 		ctx:    context,
// 	}, nil
// }

// func (connection *MongoConnection) Connect(uri string) (*mongo.Client, error) {
// 	client, err := mongo.Connect(connection.context.TODO(), options.Client().ApplyURI(uri))
// 	if err != nil {
// 		return nil, err
// 	}

// 	if err := connection.Ping(); err != nil {
// 		return nil, err
// 	}

// 	return client, nil
// }

// func (connection *MongoConnection) Ping() error {
// 	return connection.Client.Ping(connection.context.Background(), readpref.Primary())
// }

// func (connection *MongoConnection) Close() error {
// 	return connection.Client.Disconnect(connection.context.Background())
// }

// func (connection *MongoConnection) CreateDatabase(dbName string) *mongo.Database {
// 	return connection.Client.Database(dbName)
// }

// func (connection *MongoClient) CreateCollection(dbName string, collectionName string) (*mongo.Collection, error) {
// 	return connection.Client.Database(dbName).Collection(collectionName), nil
// }

// func (connection *MongoConnection) CountDocuments(collection *mongo.Collection, filter interface{}) (int64, error) {
// 	return collection.CountDocuments(connection.context.Background(), filter)
// }

// func (connection *MongoConnection) InsertOne(collection *mongo.Collection, document interface{}) (*mongo.InsertOneResult, error) {
// 	return collection.InsertOne(connection.context.Background(), document)
// }

// func (connection *MongoConnection) InsertMany(collection *mongo.Collection, documents []interface{}) (*mongo.InsertManyResult, error) {
// 	return collection.InsertMany(connection.context.Background(), documents)
// }

// func (connection *MongoConnection) FindOne(collection *mongo.Collection, filter interface{}) *mongo.SingleResult {
// 	return collection.FindOne(connection.context.Background(), filter)
// }

// func (connection *MongoConnection) FindMany(collection *mongo.Collection, filter interface{}) (*mongo.Cursor, error) {
// 	return collection.Find(connection.context.Background(), filter)
// }

// func (connection *MongoConnection) UpdateOne(collection *mongo.Collection, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
// 	return collection.UpdateOne(connection.context.Background(), filter, update)
// }

// func (connection *MongoConnection) UpdateMany(collection *mongo.Collection, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
// 	return collection.UpdateMany(connection.context.Background(), filter, update)
// }

// func (connection *MongoConnection) DeleteOne(collection *mongo.Collection, filter interface{}) (*mongo.DeleteResult, error) {
// 	return collection.DeleteOne(connection.context.Background(), filter)
// }

// func (connection *MongoConnection) DeleteMany(collection *mongo.Collection, filter interface{}) (*mongo.DeleteResult, error) {
// 	return collection.DeleteMany(connection.context.Background(), filter)
// }

// func (connection *MongoConnection) Aggregate(collection *mongo.Collection, pipeline interface{}) (*mongo.Cursor, error) {
// 	return collection.Aggregate(connection.context.Background(), pipeline)
// }
