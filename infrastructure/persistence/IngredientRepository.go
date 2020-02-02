package persistence

import (
	_ "github.com/go-sql-driver/mysql" // driver
	"github.com/hoanganf/pos_domain/entity"
	"github.com/hoanganf/pos_domain/repository"
  "gopkg.in/gorp.v1"
)

type IngredientRepositoryImpl struct {
	DbMap *gorp.DbMap
}

func NewIngredientRepository(dbMap *gorp.DbMap) repository.IngredientRepository {
	return &IngredientRepositoryImpl{DbMap: dbMap}
}

func (r *IngredientRepositoryImpl) FindById(id int64) (*entity.Ingredient, error) {
  var ingredient entity.Ingredient
  err := r.DbMap.SelectOne(&ingredient, "SELECT * FROM ingredient WHERE id=?", id)

  if err == nil {
		return &ingredient, nil
  }
	return nil, err
}
func (r *IngredientRepositoryImpl) FindByCategoryId(categoryId int64) ([]entity.Ingredient, error) {
  var ingredients []entity.Ingredient
  _,err := r.DbMap.Select(&ingredients, "SELECT * FROM ingredient WHERE category_id=?", categoryId)

  if err == nil {
		return ingredients, nil
  }

	return nil, err

}
