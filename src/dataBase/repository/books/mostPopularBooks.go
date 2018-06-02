package books

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"database/sql"
	"log"
	"net/http"
	"fmt"
	"github.com/metalscreame/GoToBoox/src/dataBase/repository/entity"
	db "github.com/metalscreame/GoToBoox/src/dataBase"
)

var B entity.Book

func GetMostPopularBooks(c *gin.Context) {
	var id sql.NullInt64
	var title sql.NullString
	row := db.GlobalDataBaseConnection.QueryRow("SELECT Id, Title FROM Books where Popularity > $1 ORDER BY Popularity", B.Popularity)
	err := row.Scan(&id, &title)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("By popularity rating", B.Popularity, "id: ", B.Id, "title: ", B.Title)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "The list of most popular books is provided", "Popularity": B.Popularity, "Book id": B.Id, "title ": B.Title})
}
