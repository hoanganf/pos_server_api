package persistence

import (
	"github.com/hoanganf/pos_domain/entity"
	"github.com/hoanganf/pos_domain/repository"
	"gopkg.in/gorp.v1"
)

type RestaurantRepositoryImpl struct {
	DbMap *gorp.DbMap
}

func NewRestaurantRepository(dbMap *gorp.DbMap) repository.RestaurantRepository {
	return &RestaurantRepositoryImpl{DbMap: dbMap}
}

func (r *RestaurantRepositoryImpl) FindById(id int64) (*entity.Restaurant, error) {
	var restaurant entity.Restaurant
	err := r.DbMap.SelectOne(&restaurant, "SELECT * FROM restaurant WHERE id=?", id)

	if err == nil {
		return &restaurant, nil
	}
	return nil, err
}
func (r *RestaurantRepositoryImpl) FindAll() ([]entity.Restaurant, error) {
	var restaurants []entity.Restaurant
	_, err := r.DbMap.Select(&restaurants, "SELECT * FROM restaurant")

	if err == nil {
		return restaurants, nil
	}

	return nil, err

}
