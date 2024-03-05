package main

import (
	"golang-poc/design_pattern/adapter/gorm"
	"golang-poc/design_pattern/adapter/mongo"
)

func main() {
	var a ProductRepository
	var b ProductRepository
	a = gorm.GormRepository{}
	b = mongo.MongoRepository{}

	a.GetProduct()
	b.GetProduct()
}
