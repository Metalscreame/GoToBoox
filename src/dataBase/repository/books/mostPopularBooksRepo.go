package books

import (
	"github.com/metalscreame/GoToBoox/src/dataBase/repository/entity"
	_"github.com/lib/pq"
	"github.com/metalscreame/GoToBoox/src/dataBase"
)

type BookRepo interface {
	GetMostPopularBooks(id int) (entity.Book, error)
}

type BookRepository struct{}

//Function GetMostPopulareBooks iterates over the DB using the SQL SELECT Request and return 5 top-rated books.
func (br BookRepository) GetMostPopularBooks (id int) ([]entity.Book, error) {
	db := dataBase.Connection
		rows, err := db.Query("SELECT Id, Title, Popularity FROM gotoboox.books ORDER BY Popularity DESC LIMIT $ID", id)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		var popularBooks []entity.Book
		for rows.Next() {
			var id int
			var title string
			var popularity float32
			err = rows.Scan(&id, &title, &popularity)
			if err != nil {
				return nil, err
			}
			book := entity.Book{Id: id, Title: title, Popularity: popularity}
			popularBooks = append(popularBooks, book)
		}
		return popularBooks, err
}