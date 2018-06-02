package services

import (
	"github.com/gin-gonic/gin"
	//"github.com/metalscreame/GoToBoox/src/models"
)


func IndexHandler(c *gin.Context) {
	type Category struct{
		Id int64
		Title string
	}
	type Books struct{
		Id int64
		Title string
		Description string
		Popularity float64
		CategoryId int64
	}
	type PopularBooks struct{      // select * from books order by desc limit N (exmpl: n = 5)
		Id int64
		Title string
		Description string
		Popularity float64
		CategoryId int64
	}

	type Data struct{
		Category
		Books
		PopularBooks
	}

	/* EXAMPLE
	categories := models.GetCategories()
			it means:
	categories := []Category{
		{0, "Title-1"},
		{1, "Title-2"},
		{2, "Title-3"},
	}
	*/



	//output := Data{cat, books, popularbooks}
	//c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": ouput})
}