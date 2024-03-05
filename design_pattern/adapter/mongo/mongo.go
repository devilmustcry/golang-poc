package mongo

type MongoRepository struct {
	db *mongo.DB
}

func (g MongoRepository) GetProduct() {
	g.db.FindOne()
}
