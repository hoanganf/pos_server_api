package persistence

import (
	"github.com/hoanganf/pos_domain/entity"
	"github.com/hoanganf/pos_domain/repository"
	"gopkg.in/gorp.v1"
)

type ProductRepositoryImpl struct {
	Table string
	DbMap *gorp.DbMap
}

func NewProductRepository(dbMap *gorp.DbMap) repository.ProductRepository {
	return &ProductRepositoryImpl{Table: "product", DbMap: dbMap}
}

func (r *ProductRepositoryImpl) FindById(id int64) (*entity.Product, error) {
	var product entity.Product
	err := r.DbMap.SelectOne(&product, "SELECT * FROM "+r.Table+" WHERE id=? AND available=?", id, 1)

	if err == nil {
		return &product, nil
	}
	return nil, err
}
func (r *ProductRepositoryImpl) FindByRestaurantId(resId int64) ([]entity.Product, error) {
	var products []entity.Product
	_, err := r.DbMap.Select(&products, "SELECT * FROM "+r.Table+" WHERE id IN(SELECT product_id FROM restaurant_product WHERE restaurant_id=?) AND available=?", resId, 1)

	if err == nil {
		return products, nil
	}

	return nil, err
}

func (r *ProductRepositoryImpl) FindAll() ([]entity.Product, error) {
	var products []entity.Product
	_, err := r.DbMap.Select(&products, "SELECT * FROM "+r.Table+" WHERE available=?", 1)

	if err == nil {
		return products, nil
	}

	return nil, err

}
