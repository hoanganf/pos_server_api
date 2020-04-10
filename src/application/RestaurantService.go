package application

import (
	"github.com/gin-gonic/gin"
	"github.com/hoanganf/pos_domain/entity"
	"github.com/hoanganf/pos_domain/entity/exception"
	"github.com/hoanganf/pos_domain/repository"
	"strconv"
)

type RestaurantService struct {
	RestaurantRepository repository.RestaurantRepository
	ProductRepository    repository.ProductRepository
	RestaurantFactory    entity.RestaurantFactory
	ProductFactory       entity.ProductFactory
}

func NewRestaurantService(
	restaurantRepository repository.RestaurantRepository,
	productRepository repository.ProductRepository,
	restaurantFactory entity.RestaurantFactory,
	productFactory entity.ProductFactory) *RestaurantService {
	return &RestaurantService{RestaurantRepository: restaurantRepository,
		ProductRepository: productRepository,
		RestaurantFactory: restaurantFactory,
		ProductFactory:    productFactory}
}

func (s *RestaurantService) GetById(c *gin.Context) {
	id, paramErr := strconv.ParseInt(c.Params.ByName("id"), 0, 64)
	if paramErr != nil {
		c.JSON(400, exception.CreateError(exception.CodeValueInvalid, "restaurantId invalid."))
		return
	}

	var restaurant, err = s.RestaurantRepository.FindById(id)
	if err != nil && restaurant == nil {
		c.JSON(404, exception.CreateError(exception.CodeNotFound, "restaurant not found."))
		return
	}

	var fields = c.Query("fields")
	if fields == "" {
		c.JSON(200, restaurant)
	} else {
		c.JSON(200, s.RestaurantFactory.Create(restaurant, fields))
	}
}

func (s *RestaurantService) GetAll(c *gin.Context) {
	var restaurants, err = s.RestaurantRepository.FindAll()
	if err != nil && restaurants == nil {
		c.JSON(404, exception.CreateError(exception.CodeNotFound, "restaurant not found."))
		return
	}

	var fields = c.Query("fields")
	if fields == "" {
		c.JSON(200, restaurants)
	} else {
		c.JSON(200, s.RestaurantFactory.CreateList(restaurants, fields))
	}
}

func (s *RestaurantService) GetProducts(c *gin.Context) {
	id, paramErr := strconv.ParseInt(c.Params.ByName("id"), 0, 64)
	if paramErr != nil {
		c.JSON(400, exception.CreateError(exception.CodeValueInvalid, "restaurantId invalid."))
		return
	}

	var products, err = s.ProductRepository.FindByRestaurantId(id)
	if err != nil && products == nil {
		c.JSON(404, exception.CreateError(exception.CodeNotFound, "product not found."))
		return
	}

	var fields = c.Query("fields")
	if fields == "" {
		c.JSON(200, products)
	} else {
		c.JSON(200, s.ProductFactory.CreateList(products, fields))
	}
}
