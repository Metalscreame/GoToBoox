package categories

import (
	"log"
	"github.com/metalscreame/GoToBoox/src/dataBase/repository/entity"
	_"github.com/lib/pq"
	"github.com/metalscreame/GoToBoox/src/dataBase"
)

type CategoryRepository interface{
	GetAllCategories () ([]entity.Categories, error)
}

type CategoryRepoPq struct{}

var CategoryRepo CategoryRepository

//Function GetAllCategories creates a list of all categories currently available and order them alphabetically
func (cr CategoryRepoPq) GetAllCategories ( ) ([]entity.Categories, error) {

	rows, err := dataBase.Connection.Query("SELECT id, title FROM gotoboox.categories")
	if err != nil {
		log.Println("Unknown error occurred")
	}
	defer rows.Close()
	var allCategories []entity.Categories
	for rows.Next() {
		var id int
		var title string
		if err := rows.Scan(&id, &title); err != nil {
			log.Fatal(err)
		}
		category := entity.Categories{id, title}
		allCategories = append(allCategories, category)
	}
	return allCategories, err
}

