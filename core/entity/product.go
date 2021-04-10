package entity

type Product struct {
	Code      string `bson:"code"`
	Name        string   `bson:"name"`
	Price    int32 `bson:"price"`
}
