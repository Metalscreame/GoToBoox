package services

import (
	"github.com/metalscreame/GoToBoox/src/dataBase/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)


//TagsService is a struct with books and tags repository
type TagsService struct {
	BooksRepo repository.BookRepository
	TagsRepo repository.TagsRepository

}
//NewTagsService is a func that initialize TagsService struct
func NewTagsService(repository repository.BookRepository, tagsRepo repository.TagsRepository) *TagsService {
	return &TagsService{
		BooksRepo: repository,
		TagsRepo: tagsRepo,
	}
}

func (b *TagsService) getTags(c *gin.Context) {
	type Data struct {
		Tags []repository.Tags
	}
		if tags, err := b.TagsRepo.GetListOfTags();
			err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else {
			output := Data{tags}
			c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": output})
			return
		}
	}
