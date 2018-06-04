package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/metalscreame/GoToBoox/src/dataBase/postgres"
)

func (cs *CategoriesService) AllCategories (c *gin.Context) {
	allCategories, _ := postgres.CategoryRepo.GetAllCategories()
	if len(allCategories) > 0 {
		c.JSON(http.StatusOK,  allCategories)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "No categories found"} )
	}
}
