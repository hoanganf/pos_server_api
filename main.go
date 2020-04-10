package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"./src"
)

func main() {
	//connect to DB
	bean, err := src.InitBean()
	defer bean.DestroyBean()
	if err != nil {
		log.Fatalln("can not create bean", err)
	}
	//setup router
	router := gin.Default()
	v1 := router.Group("v1")
	{
		v1.GET("/ingredients/:id", bean.IngredientService.GetById)
		v1.GET("/ingredients", bean.IngredientService.GetAll)
		v1.GET("/restaurants/:id", bean.RestaurantService.GetById)
		v1.GET("/restaurants", bean.RestaurantService.GetAll)
		v1.GET("/restaurants/:id/products", bean.RestaurantService.GetProducts)
		v1.POST("/user", bean.UserService.Login)

		/*  v1.GET("/instructions/:id", app.GetInstruction)
		    v1.POST("/instructions", app.PostInstruction)
		    v1.PUT("/instructions/:id", app.UpdateInstruction)
		    v1.DELETE("/instructions/:id", app.DeleteInstruction)*/
	}
	router.Run(":8080")
}
