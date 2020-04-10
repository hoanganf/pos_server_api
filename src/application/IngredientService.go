package application

import (
	"github.com/gin-gonic/gin"
	"github.com/hoanganf/pos_domain/entity"
	"github.com/hoanganf/pos_domain/repository"
	"github.com/hoanganf/pos_domain/entity/exception"
	"strconv"
)

type IngredientService struct {
	IngredientRepository repository.IngredientRepository
	IngredientFactory    entity.IngredientFactory
}

func NewIngredientService(ingredientRepository repository.IngredientRepository,
	IngredientFactory entity.IngredientFactory) *IngredientService {
	return &IngredientService{IngredientRepository: ingredientRepository,
		IngredientFactory: IngredientFactory}
}

func (s *IngredientService) GetById(c *gin.Context) {
	id, paramErr := strconv.ParseInt(c.Params.ByName("id"), 0, 64)
	if paramErr != nil {
		c.JSON(400, exception.CreateError(exception.CodeValueInvalid, "ingredientId invalid."))
		return
	}

	var ingredient, err = s.IngredientRepository.FindById(id)
	if err != nil && ingredient == nil {
		c.JSON(404, exception.CreateError(exception.CodeNotFound, "ingredient not found."))
		return
	}

	var fields = c.Query("fields")
	if fields == "" {
		c.JSON(200, ingredient)
	} else {
		c.JSON(200, s.IngredientFactory.Create(ingredient, fields))
	}
}

func (s *IngredientService) GetAll(c *gin.Context) {
	var ingredients, err = s.IngredientRepository.FindAll()
	if err != nil && ingredients == nil {
		c.JSON(404, exception.CreateError(exception.CodeNotFound, "ingredient not found."))
		return
	}

	var fields = c.Query("fields")
	if fields == "" {
		c.JSON(200, ingredients)
	} else {
		c.JSON(200, s.IngredientFactory.CreateList(ingredients, fields))
	}

}
