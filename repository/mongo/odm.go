package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"hexagonal-architecture-example/config"
	"log"
	"reflect"
	"sync"
)

type SortParam struct {
	SortBy string // bson field name
	Type   int    // -1 or 1
}

type Data interface{}


type dmManager struct {
	Ctx context.Context
	Db  *mongo.Database
}

var singletonDmManager *dmManager
var onceDmManager sync.Once

func GetDmManager() *dmManager {
	onceDmManager.Do(func() {
		log.Println("[INFO] Starting Initializing Singleton DB Manager")
		singletonDmManager = &dmManager{}
		singletonDmManager.initConnection()
	})
	return singletonDmManager
}

func (dm *dmManager) initConnection() {
	// Base context.
	ctx := context.Background()
	dm.Ctx = ctx
	clientOpts := options.Client().ApplyURI(config.DatabaseConnectionString)
	//clientOpts := options.Client().ApplyURI("mongodb+srv://MyDataBase:MongoDBDatabase@harun.wzryz.mongodb.net/<dbname>?retryWrites=true&w=majority")
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Println("[ERROR] DB Connection error:", err.Error())
		return
	}

	db := client.Database(config.DatabaseName)
	dm.Db = db

	log.Println("[INFO] Initialized Singleton DB Manager")
}

func (dm *dmManager) FindOne(collectionName string, filter interface{}, objType reflect.Type) interface{} {
	coll := dm.Db.Collection(collectionName)

	findResult := coll.FindOne(dm.Ctx, filter)
	if err := findResult.Err(); err != nil {
		return nil
	}

	objValue := reflect.New(objType)
	obj := objValue.Interface()
	err := findResult.Decode(obj)
	if err != nil {
		log.Println("[ERROR] Find document decoding:", err.Error())
		return nil
	}
	return obj
}

