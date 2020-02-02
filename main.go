package main

import (
  "github.com/gin-gonic/gin"
  "./config"
  "./application"
  "./infrastructure/persistence"
)

func main() {
  //connect to DB
  dbMap := config.InitDB()
  defer dbMap.Db.Close()
  service := application.NewIngredientService(persistence.NewIngredientRepository(dbMap))
  //setup router
  router := gin.Default()
  v1 := router.Group("v1")
  {
    v1.GET("/ingredients/:id", service.FindById)
    v1.GET("/ingredients", service.FindByCategoryId)

  /*  v1.GET("/instructions/:id", app.GetInstruction)
    v1.POST("/instructions", app.PostInstruction)
    v1.PUT("/instructions/:id", app.UpdateInstruction)
    v1.DELETE("/instructions/:id", app.DeleteInstruction)*/
  }
  router.Run(":8080")
}
