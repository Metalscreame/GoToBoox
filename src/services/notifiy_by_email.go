package services

import (
	"github.com/metalscreame/GoToBoox/src/dataBase/postgres"
	"github.com/metalscreame/GoToBoox/src/dataBase"
	"log"
	"gopkg.in/gomail.v2"
	"crypto/tls"
	"bytes"
	"time"
	"github.com/metalscreame/GoToBoox/src/dataBase/repository"
)

var d = gomail.NewDialer("smtp.gmail.com", 587, "GoToBooX", "hjvfhekbn")

//DailyEmailNotifications is a functions that sends emails every 24 hours with lists of
func DailyEmailNotifications() {
	s:=NewBookService(postgres.NewBooksRepository(dataBase.Connection),postgres.NewPostgresUsersRepo(dataBase.Connection))
	for true{
		time.Sleep(24*time.Hour)

		users,err:=s.UsersRepo.GetAllUsers()
		if err!=nil{
			log.Println(err)
			continue
		}

		books,err:=s.BooksRepo.GetAll()
		if err!=nil{
			log.Println(err)
			continue
		}

		sendCloser, err := d.Dial()
		if err != nil {
			log.Println(err)
			continue
		}
		prepearedMsg := prepareMsgAllAvailebleBooksEveryDay(books)
		msg := gomail.NewMessage()
		for _, user := range users {
			d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

			msg.SetHeader("From", "GoToBooX@gmail.com")
			msg.SetAddressHeader("To", user.Email, user.Nickname)
			msg.SetHeader("Subject", "A book has been reserved")
			msg.SetBody("text/html", prepearedMsg)
			msg.Attach("/static/images/logo.jpg")

			if err := gomail.Send(sendCloser, msg); err != nil {
				log.Printf("Could not send email to %q: %v", user.Email, err)
			}
			msg.Reset()
		}
	}
}

//NofityAllBookReserved is a func that notifies by email everyone when the book is reserved
func NofityAllBookReserved(bookTitle, bookDescription string) {
	s := NewUserService(postgres.NewPostgresUsersRepo(dataBase.Connection))
	listOfUsersEmails, err := s.UsersRepo.GetUsersEmailToNotifyReserved()
	if err != nil {
		log.Println(err)
		return
	}

	sc, err := d.Dial()
	if err != nil {
		log.Println(err)
		return
	}
	msg := prepareMsgNewBook(bookTitle, bookDescription)

	m := gomail.NewMessage()
	for _, user := range listOfUsersEmails {
		d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

		m.SetHeader("From", "GoToBooX@gmail.com")
		m.SetAddressHeader("To", user.Email, user.Nickname)
		m.SetHeader("Subject", "A book has been reserved")
		m.SetBody("text/html", msg)
		m.Attach("/static/images/logo.jpg")

		if err := gomail.Send(sc, m); err != nil {
			log.Printf("Could not send email to %q: %v", user.Email, err)
		}
		m.Reset()
	}
}

func NotifyAllNewBook(bookTitle, bookDescription string) {
	s := NewUserService(postgres.NewPostgresUsersRepo(dataBase.Connection))
	listOfUsersEmails, err := s.UsersRepo.GetUsersEmailToNotifyNewBook()
	if err != nil {
		log.Println(err)
		return
	}
	sc, err := d.Dial()
	if err != nil {
		log.Println(err)
		return
	}
	msg := prepareMsgBookReserved(bookTitle, bookDescription)

	m := gomail.NewMessage()
	for _, user := range listOfUsersEmails {
		d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

		m.SetHeader("From", "GoToBooX@gmail.com")
		m.SetAddressHeader("To", user.Email, user.Nickname)
		m.SetHeader("Subject", "A book has been reserved")
		m.SetBody("text/html", msg)
		m.Attach("/static/images/logo.jpg")

		if err := gomail.Send(sc, m); err != nil {
			log.Printf("Could not send email to %q: %v", user.Email, err)
		}
		m.Reset()
	}
}

func prepareMsgNewBook(bookTitle, bookDescription string) string {
	var b bytes.Buffer
	b.WriteString("Hello from GoToBooX!\r\nWe have a new book at GoToBooX for you!\r\n")
	b.WriteString(bookTitle)
	b.WriteString("\r\n\n\n")
	b.WriteString("Description\n\n")
	b.WriteString(bookDescription)
	b.WriteString("\r\n\nHave a nice day! Don't forget to read the BooX :)")
	return b.String()
}

func prepareMsgBookReserved(bookTitle, bookDescription string) string {
	var b bytes.Buffer
	b.WriteString("Hello from GoToBooX!\nThis book has been reserved!\r\n")
	b.WriteString(bookTitle)
	b.WriteString("\r\n\n\n")
	b.WriteString("Description\n\n")
	b.WriteString(bookDescription)
	b.WriteString("\r\n\nHave a nice day! Don't forget to read the BooX :)")
	return b.String()
}

func prepareMsgAllAvailebleBooksEveryDay(books []repository.Book) string {
	var b bytes.Buffer
	b.WriteString("Hello from GoToBooX!\nThere are the list of books that can be read!\r\n")
	b.WriteString("\r\n\n")
	b.WriteString("Take a look: \r\n\n")
	for _, book := range books {
		b.WriteString(book.Title)
		b.WriteString("\r\n")
	}
	b.WriteString("\r\n\nHave a nice day! Don't forget to read the BooX :)")
	return b.String()
}
