package application
import (
  "github.com/gin-gonic/gin"
  "github.com/hoanganf/pos_domain/repository"
  "strconv"
)
type IngredientService struct {
  IngredientRepository repository.IngredientRepository
}

func NewIngredientService(ingredientRepository repository.IngredientRepository) *IngredientService {
	return &IngredientService{IngredientRepository: ingredientRepository}
}

func (s *IngredientService) FindById(c *gin.Context){
  id,paramErr := strconv.ParseInt(c.Params.ByName("id"), 0, 64)
  if(paramErr!=nil){
    c.JSON(400, gin.H{"error": "bad request"})
    return
  }

  var ingredient,err=s.IngredientRepository.FindById(id)
  if(err!=nil && ingredient==nil){
    c.JSON(404, gin.H{"error": "not found"})
    return
  }

  c.JSON(404, gin.H{"error": "instruction not found"})

}

func (s *IngredientService) FindByCategoryId(c *gin.Context){
  id,paramErr := strconv.ParseInt(c.Query("categoryId"), 0, 64)
  if(paramErr!=nil){
    c.JSON(400, gin.H{"error": "bad request"})
    return
  }
  var ingredients,err=s.IngredientRepository.FindByCategoryId(id)
  if(err!=nil && ingredients==nil){
    c.JSON(404, gin.H{"error": "not found"})
    return
  }

  c.JSON(200, ingredients)

}
