package gorm

type GormRepository struct {
	db *gorm.DB
}

func (g GormRepository) GetProduct() {
	g.db.FindById()
}
