package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/metalscreame/GoToBoox/src/dataBase/repository"
)

func NewCategoriesService(repository repository.CategoryRepository) *CategoriesService {
	return &CategoriesService{
		CategoriesRepoPq: repository,
	}
}

type CategoriesService struct {
	CategoriesRepoPq repository.CategoryRepository
}

func (cs *CategoriesService) AllCategories (c *gin.Context) {
	allCategories, _ := cs.CategoriesRepoPq.GetAllCategories()
	if len(allCategories) > 0 {
		c.JSON(http.StatusOK,  allCategories)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "No categories found"} )
	}
}
