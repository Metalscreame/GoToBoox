package commonFunctions

import (
	"github.com/metalscreame/GoToBoox/src/services"
	"github.com/metalscreame/GoToBoox/src/dataBase/postgres"
	"github.com/metalscreame/GoToBoox/src/dataBase"
	"log"
	"gopkg.in/gomail.v2"
	"crypto/tls"
	"bytes"
)

var d = gomail.NewDialer("smtp.gmail.com", 587, "GoToBooX", "hjvfhekbn")

func NofityAllBookReserved(bookTitle, bookDescription string) {
	s := services.NewUserService(postgres.NewPostgresUsersRepo(dataBase.Connection))
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
		m.Attach("/static/images/logo.png")

		if err := gomail.Send(sc, m); err != nil {
			log.Printf("Could not send email to %q: %v", user.Email, err)
		}
		m.Reset()
	}
}

func NotifyAllNewBook(bookTitle, bookDescription string) {
	s := services.NewUserService(postgres.NewPostgresUsersRepo(dataBase.Connection))
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
		m.Attach("/static/images/logo.png")

		if err := gomail.Send(sc, m); err != nil {
			log.Printf("Could not send email to %q: %v", user.Email, err)
		}
		m.Reset()
	}

}

func prepareMsgNewBook(bookTitle, bookDescription string) string {
	var b bytes.Buffer
	b.WriteString("Hello from GoToBooX!\nTWe have a new book at GoToBooX for you!\n")
	b.WriteString(bookTitle)
	b.WriteString("n\n\n")
	b.WriteString("Descrition:\n\n")
	b.WriteString(bookDescription)
	b.WriteString("\n\nHave a nice day! Dont forget to read the BooX :)")
	return b.String()
}

func prepareMsgBookReserved(bookTitle, bookDescription string) string {
	var b bytes.Buffer
	b.WriteString("Hello from GoToBooX!\nThis book has been reserved!\n")
	b.WriteString(bookTitle)
	b.WriteString("n\n\n")
	b.WriteString("Descrition:\n\n")
	b.WriteString(bookDescription)
	b.WriteString("\n\nHave a nice day! Dont forget to read the BooX :)")
	return b.String()
}