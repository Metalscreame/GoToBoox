package postgres

import (
	"log"
	"github.com/metalscreame/GoToBoox/src/dataBase"
	"github.com/metalscreame/GoToBoox/src/dataBase/repository"
)

type CategoryRepoPq struct{}

var CategoryRepo repository.CategoryRepository

//Function GetAllCategories creates a list of all categories currently available and order them alphabetically
func (cr CategoryRepoPq) GetAllCategories ( ) ([]repository.Categories, error) {

	rows, err := dataBase.Connection.Query("SELECT id, title FROM gotoboox.categories")
	if err != nil {
		log.Println("Unknown error occurred")
		return nil, err
	}
	defer rows.Close()
	var allCategories []repository.Categories
	for rows.Next() {
		var id int
		var title string
		if err := rows.Scan(&id, &title); err != nil {
			log.Fatal(err)
		}
		category := repository.Categories{id, title}
		allCategories = append(allCategories, category)
	}
	return allCategories, err
}


