package models

import (
	"time"
	"github.com/metalscreame/GoToBoox/src/dataBase/postgres"
	"github.com/metalscreame/GoToBoox/src/services"
	"github.com/metalscreame/GoToBoox/src/dataBase"
	"log"
	"github.com/metalscreame/GoToBoox/src/dataBase/repository"
)

const preReservedTimeAllowedSecs = 1000

//PreReservedTimer is a funtion that changes the status of a book to free if user didnt propose book to exchange
//must be started from "go PreReservedTimer"
func PreReservedTimer(bookId int){
	time.Sleep(preReservedTimeAllowedSecs)
	s:=services.NewBookService(postgres.NewBooksRepository(dataBase.Connection))
	book,err:=s.BooksRepo.GetByID(bookId)
	if err!=nil{
		log.Println(err)
		return
	}

	if book.State == repository.BookStatePreReserved{
		s.BooksRepo.UpdateBookState(bookId,repository.BookStateFree)
		return
	}
	return
}
