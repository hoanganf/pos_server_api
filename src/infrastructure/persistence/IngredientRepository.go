package persistence

import (
	"github.com/hoanganf/pos_domain/entity"
	"github.com/hoanganf/pos_domain/repository"
	"gopkg.in/gorp.v1"
)

type IngredientRepositoryImpl struct {
	Table string
	DbMap *gorp.DbMap
}

func NewIngredientRepository(dbMap *gorp.DbMap) repository.IngredientRepository {
	return &IngredientRepositoryImpl{Table: "ingredient", DbMap: dbMap}
}

func (r *IngredientRepositoryImpl) FindById(id int64) (*entity.Ingredient, error) {
	var ingredient entity.Ingredient
	err := r.DbMap.SelectOne(&ingredient, "SELECT * FROM "+r.Table+" WHERE id=? AND available=1", id)

	if err == nil {
		return &ingredient, nil
	}
	return nil, err
}
func (r *IngredientRepositoryImpl) FindAll() ([]entity.Ingredient, error) {
	var ingredients []entity.Ingredient
	_, err := r.DbMap.Select(&ingredients, "SELECT * FROM "+r.Table+" WHERE available=1")

	if err == nil {
		return ingredients, nil
	}

	return nil, err

}
