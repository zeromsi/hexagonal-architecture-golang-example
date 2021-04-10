package mongo

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
	"hexagonal-architecture-example/core/entity"
	"hexagonal-architecture-example/core/repository"
	"log"
	"time"
)

var (
	ProductCollection="productCollection"
	NO_KEYS_FOUND_ERROR=errors.New("NO_KEYS_FOUND")
)
type productRepository struct {
	manager  *dmManager
	timeout  time.Duration
}

func (p productRepository) Store(product entity.Product) error {
	coll := p.manager.Db.Collection(ProductCollection)
	_, err := coll.InsertOne(p.manager.Ctx, product)
	if err != nil {
		log.Println("[ERROR] Insert document:", err.Error())
		return err
	}
	return nil
}

func (p productRepository) Update(product entity.Product) error {
	filter := bson.M{
		"$and": []bson.M{
			{"code": product.Code},
		},
	}

	update := bson.M{
		"$set": product,
	}

	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	coll := p.manager.Db.Collection(ProductCollection)
	err := coll.FindOneAndUpdate(p.manager.Ctx, filter, update, &opt)
	if err.Err() != nil {
		log.Println("[ERROR]", err.Err())
		return err.Err()
	}
	return nil
}

func (p productRepository) FindAll() []entity.Product {
	data := []entity.Product{}
	query:= bson.M{"$and": []bson.M{}}
	coll := p.manager.Db.Collection(ProductCollection)
	curser, _ := coll.Find(p.manager.Ctx, query)
	for curser.Next(context.TODO()) {
		elemValue := new(entity.Product)
		err := curser.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		data = append(data, *elemValue)
	}
	return data
}

func (p productRepository) FindByCode(code string) entity.Product {
	query := bson.M{
		"$and": []bson.M{
			{"code": code},
		},
	}
	temp := new(entity.Product)
	coll := p.manager.Db.Collection(ProductCollection)
	result := coll.FindOne(p.manager.Ctx, query)
	err := result.Decode(temp)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return *temp
}

func (p productRepository) Delete(code string) error {
	query := bson.M{
		"$and": []bson.M{
			{"code": code},
		},
	}
	coll := p.manager.Db.Collection(ProductCollection)
	_,err:=coll.DeleteOne(p.manager.Ctx,query)
	if err != nil {
		log.Println("[ERROR]", err)
		return err
	}
	return nil
}

func NewProductRepository(mongoTimeout int) repository.ProductRepository {
	repo := &productRepository{
		timeout:  time.Duration(mongoTimeout) * time.Second,
	}
	manager:=GetDmManager()
	repo.manager = manager
	return repo
}
