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

const (
	senderEmailAddr = "GoToBooX@gmail.com"
	emailUsername   = "GoToBooX"
	emailPassword   = "hjvfhekbn"
)

var dialer *gomail.Dialer

//ConfigureEmailDialer is a function that configures dialer for GoToBooX email sender
func ConfigureEmailDialer() {
	dialer = gomail.NewDialer("smtp.gmail.com", 587, emailUsername, emailPassword)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
}

//DailyEmailNotifications is a functions that sends emails every 24 hours with lists of
func DailyEmailNotifications() {
	s := NewBookService(postgres.NewBooksRepository(dataBase.Connection), postgres.NewPostgresUsersRepo(dataBase.Connection))
	for {
		time.Sleep(24 * time.Hour)

		users, err := s.UsersRepo.GetAllUsers()
		if err != nil {
			log.Printf("Error at daily email notify while getting all users %v\n", time.Now())
			log.Println(err)
			continue
		}

		books, err := s.BooksRepo.GetAll()
		if err != nil {
			log.Println(err)
			log.Printf("Error at daily email notify while getting all books at %v\n", time.Now())
			continue
		}

		sendCloser, _ := dialer.Dial()
		prepearedMsg := prepareMsgAllAvailebleBooksEveryDay(books)
		msg := gomail.NewMessage()

		for _, user := range users {

			msg.SetHeader("From", senderEmailAddr)
			msg.SetAddressHeader("To", user.Email, user.Nickname)
			msg.SetHeader("Subject", "A book has been reserved")
			msg.SetBody("text/html", prepearedMsg)
			msg.Attach("/static/images/logo.jpg")

			if err := gomail.Send(sendCloser, msg); err != nil {
				log.Printf("Could not send email to %q: %v \n at %v\n", user.Email, err, time.Now())
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
		log.Printf("Error at notifyAllBookReserved email notify while getting all users %v\n", time.Now())
		log.Println(err)
		return
	}

	sc, _ := dialer.Dial()
	msg := prepareMsgNewBook(bookTitle, bookDescription)
	m := gomail.NewMessage()

	for _, user := range listOfUsersEmails {

		m.SetHeader("From", senderEmailAddr)
		m.SetAddressHeader("To", user.Email, user.Nickname)
		m.SetHeader("Subject", "A book has been reserved")
		m.SetBody("text/html", msg)
		m.Attach("/static/images/logo.jpg")

		if err := gomail.Send(sc, m); err != nil {
			log.Printf("Could not send email to %q: %v \n at %v\n", user.Email, err, time.Now())
		}
		m.Reset()
	}
}

//NotifyAllNewBook is a func that notifies by email everyone when the book is added
func NotifyAllNewBook(bookTitle, bookDescription string) {
	s := NewUserService(postgres.NewPostgresUsersRepo(dataBase.Connection))
	listOfUsersEmails, err := s.UsersRepo.GetUsersEmailToNotifyNewBook()
	if err != nil {
		log.Println(err)
		return
	}
	sc, _ := dialer.Dial()

	msg := prepareMsgBookReserved(bookTitle, bookDescription)

	m := gomail.NewMessage()
	for _, user := range listOfUsersEmails {
		dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

		m.SetHeader("From", senderEmailAddr)
		m.SetAddressHeader("To", user.Email, user.Nickname)
		m.SetHeader("Subject", "A book has been reserved")
		m.SetBody("text/html", msg)
		m.Attach("/static/images/logo.jpg")

		if err := gomail.Send(sc, m); err != nil {
			log.Printf("Could not send email to %q: %v \n at %v", user.Email, err, time.Now())
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
