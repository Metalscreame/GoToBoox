package services

import (
	"log"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/metalscreame/GoToBoox/src/services/midlwares"
	"github.com/metalscreame/GoToBoox/src/dataBase"
	"github.com/metalscreame/GoToBoox/src/dataBase/postgres"
	"os"
)

const (
	apiRoute = "/api/v1"
)
var router *gin.Engine

//Start is a function that starts server and initializes all the routes.
func Start() {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		println("PORT is required\nFor localhosts setup sys env \"PORT\" as 8080 and reload IDE")
		log.Fatal("PORT is required\nFor localhosts setup sys env \"PORT\" as 8080 and reload IDE")
	}

	gin.SetMode(gin.ReleaseMode)
	router = gin.New()
	router.Use(gin.Logger())
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*.html")


	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
			"title": "GoToBooX",
			"page" : "main",
		})
	})
	router.GET("/location", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
			"title": "GoToBooX - location",
			"page" : "location",
		})
	})



	router.GET(apiRoute, IndexHandler)
	initUserProfileRoutes()
	initBooksRoutes()
	router.Run(":" + port)
}






func initUserProfileRoutes() {
	// Use the SetUserStatus middleware for every route to set a flag
	// indicating whether the request was from an authenticated user or not
	router.Use(midlwares.SetUserStatus())

	service := NewUserService(postgres.NewPostgresUsersRepo(dataBase.Connection))
	userRoutes := router.Group(apiRoute)
	{
		// Handle POST requests at /api/v1/login
		// Ensure that the user is not logged in by using the middleware
		userRoutes.POST("/login", midlwares.EnsureNotLoggedIn(), service.PerformLoginHandler)

		// Handle GET requests at /api/v1/logout
		// Ensure that the user is logged in by using the middleware
		userRoutes.GET("/logout", midlwares.EnsureLoggedIn(), service.LogoutHandler)

		// Handle POST requests at /api/v1/register
		// Ensure that the user is not logged in by using the middleware
		userRoutes.POST("/register", midlwares.EnsureNotLoggedIn(), service.UserCreateHandler)

		// Handle the GET requests at /api/v1/userProfile
		// Show the user's profile page
		// Ensure that the user is logged in by using the middleware
		userRoutes.GET("/userProfile", midlwares.EnsureLoggedIn(), service.UserGetHandler)

		// Handle the GET requests at /api/v1/register
		// Show the user's profile page
		// Ensure that the user is logged in by using the middleware
		userRoutes.PUT("/userProfile", midlwares.EnsureLoggedIn(), service.UserUpdateHandler)

		// Handle the GET requests at /api/v1/userProfile
		// Show the user's profile page
		// Ensure that the user is logged in by using the middleware
		userRoutes.DELETE("/userProfile", midlwares.EnsureLoggedIn(), service.UserDeleteHandler)
	}

	// Show the login page
	// Ensure that the user is not logged in by using the middleware
	router.GET("/login", midlwares.EnsureNotLoggedIn(), ShowLoginPage)

	// Show the registration page
	// Ensure that the user is not logged in by using the middleware
	router.GET("/register", midlwares.EnsureNotLoggedIn(), ShowRegistrPage)

	// Show the user's profile page or login page
	router.GET("/userProfile", UserProfileHandler)

	//Shows the lock page
	router.GET("/uploadPage/:book_id", midlwares.EnsureLoggedIn(), ShowUploadBookPage)

	// Show the user's profile page
	// Ensure that the user is logged in by using the middleware
	router.GET("/userProfilePage", midlwares.EnsureLoggedIn(), service.ShowUsersProfilePage)

}

func initBooksRoutes() {
	bookService := NewBookService(postgres.NewBooksRepository(dataBase.Connection), postgres.NewPostgresUsersRepo(dataBase.Connection))
	//get all books in certain category
	router.GET("categories/:cat_id/books", bookService.getBooks)
	//get all books
	router.GET("/books", bookService.showAllBooks)
	//get books by it's ID
	//router.GET("/categories/:cat_id/book/:book_id", bookService.getBook)
	router.GET("/books/m/mostPopularBooks", bookService.FiveMostPop)

	router.GET("/api/v1/book/:book_id", bookService.getBook)
	router.GET("/book/:book_id", ShowBook)

	router.GET("/api/v1/books/showReserved", bookService.showReservedBooksByUser)

	router.GET("/api/v1/books/taken", bookService.showAllTakenBooks)
	router.GET("/api/v1/books/taken/0",bookService.showTakenBookByUser)
	router.GET("/books/taken/:id", ShowTakenBooksPage)

	router.POST("/api/v1/insertNewBook/:book_id", midlwares.EnsureLoggedIn(), bookService.insertNewBook)
	router.GET("/api/v1/updateBookStatus/:book_id", bookService.UpdateBookStatusToReturningFromTaken)
	router.GET("/api/v1/updateBookStatusReturn/:book_id/:reserved_book_id", bookService.UpdateBookStatusToReturning)
	router.GET("/api/v1/makeBookCross", bookService.ExchangeBook)
}
