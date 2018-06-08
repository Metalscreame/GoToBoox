package postgres

import (
	"github.com/metalscreame/GoToBoox/src/dataBase"
	"github.com/metalscreame/GoToBoox/src/dataBase/repository"
	_"database/sql"
	"errors"
)

type CategoryRepoPq struct {}

//Function GetAllCategories creates a list of all categories currently available and order them alphabetically
func (cr CategoryRepoPq) GetAllCategories ( ) ([]repository.Categories, error) {
	rows, err := dataBase.Connection.Query("SELECT id, title FROM categories LIMIT 20")
	if err != nil {
		return nil, errors.New("Failed to create to get the access to the database")
	}
	defer rows.Close()
	var allCategories []repository.Categories
	var categoryRepo repository.Categories
	for rows.Next() {
		if err := rows.Scan(&categoryRepo.ID,&categoryRepo.Title); err != nil {
			return nil, errors.New("Failed to create the struct/read the data")
		}
		allCategories = append(allCategories, categoryRepo)
	}
	return allCategories, nil
}

