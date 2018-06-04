package services

import (
	"github.com/gin-gonic/gin"
	"github.com/metalscreame/GoToBoox/src/dataBase/repository/categories"
	"net/http"
)

func (cs *CategoriesService) AllCategories (c *gin.Context) {
	allCategories, _ := categories.CategoryRepo.GetAllCategories()
	if len(allCategories) > 0 {
		c.JSON(http.StatusOK,  allCategories)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "No categories found"} )
	}
}
