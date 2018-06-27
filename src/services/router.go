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
		isLoggedIn := midlwares.CheckLoggedIn(c)
		c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
			"title": "GoToBooX",
			"page" : "main",
			"isLoggedIn": isLoggedIn,
		})
	})
	router.GET("/location", func(c *gin.Context) {
		isLoggedIn := midlwares.CheckLoggedIn(c)
		c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
			"title": "GoToBooX - location",
			"page" : "location",
			"isLoggedIn": isLoggedIn,

		})
	})
	router.GET("/search", func(c *gin.Context) {
		isLoggedIn := midlwares.CheckLoggedIn(c)
		c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
			"title": "GoToBooX - search",
			"page" : "search",
			"isLoggedIn": isLoggedIn,

		})
	})
	router.GET(apiRoute, IndexHandler)
	initUserProfileRoutes()
	initBooksRoutes()
	initTagsRoutes()
	router.Run(":" + port)
}






func initUserProfileRoutes() {
	// Use the SetUserStatus middleware for every route to set a flag
	// indicating whether the request was from an authenticated user or not
	router.Use(midlwares.SetUserStatus())

	userService := NewUserService(postgres.NewPostgresUsersRepo(dataBase.Connection))
	userRoutes := router.Group(apiRoute)
	{
		// Handle POST requests at /api/v1/login
		// Ensure that the user is not logged in by using the middleware
		userRoutes.POST("/login", midlwares.EnsureNotLoggedIn(), userService.PerformLoginHandler)

		// Handle GET requests at /api/v1/logout
		// Ensure that the user is logged in by using the middleware
		userRoutes.GET("/logout", midlwares.EnsureLoggedIn(), userService.LogoutHandler)

		// Handle POST requests at /api/v1/register
		// Ensure that the user is not logged in by using the middleware
		userRoutes.POST("/register", midlwares.EnsureNotLoggedIn(), userService.UserCreateHandler)

		// Handle the GET requests at /api/v1/userProfile
		// Show the user's profile page
		// Ensure that the user is logged in by using the middleware
		userRoutes.GET("/userProfile", midlwares.EnsureLoggedIn(), userService.UserGetHandler)

		// Handle the GET requests at /api/v1/register
		// Show the user's profile page
		// Ensure that the user is logged in by using the middleware
		userRoutes.PUT("/userProfile", midlwares.EnsureLoggedIn(), userService.UserUpdateHandler)

		// Handle the GET requests at /api/v1/userProfile
		// Show the user's profile page
		// Ensure that the user is logged in by using the middleware
		userRoutes.DELETE("/userProfile", midlwares.EnsureLoggedIn(), userService.UserDeleteHandler)
	}

	// Show the login page
	// Ensure that the user is not logged in by using the middleware
	router.GET("/login", midlwares.EnsureNotLoggedIn(), ShowLoginPage)

	// Show the registration page
	// Ensure that the user is not logged in by using the middleware
	router.GET("/register/:bookid", midlwares.EnsureNotLoggedIn(), ShowRegistrPage)

	// Show the user's profile page or login page
	router.GET("/userProfile", UserProfileHandler)

	//Shows the lock page
	router.GET("/uploadPage/:book_id", midlwares.EnsureLoggedIn(), ShowUploadBookPage)

	// Show the user's profile page
	// Ensure that the user is logged in by using the middleware
	router.GET("/userProfilePage", midlwares.EnsureLoggedIn(), userService.ShowUsersProfilePage)

	router.GET("/userComments/:nickname", ShowCommentsPage)
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
	router.POST("/api/v1/books/search", bookService.getBookBySearch)
	router.GET("/api/v1/book/:book_id", bookService.getBook)

	router.GET("/book/:book_id", ShowBook)
	router.GET("/api/v1/books/taken", bookService.showAllTakenBooks)


	router.GET("/api/v1/books/showReserved", bookService.ShowReservedBooksByUser)
	router.GET("/api/v1/books/taken/0",bookService.ShowTakenBookByUser)
	router.GET("/books/taken/:id", ShowTakenBooksPage)

	router.POST("/api/v1/insertNewBook/:book_id", midlwares.EnsureLoggedIn(), bookService.InsertNewBook)

	router.GET("/api/v1/updateBookStatus/:book_id", bookService.UpdateBookStatusToReturningFromTaken)
	router.GET("/api/v1/updateBookStatusReturn/:book_id/:reserved_book_id", bookService.UpdateBookStatusToReturning)
	router.GET("/api/v1/makeBookCross", bookService.ExchangeBook)
	
	commentsService:= NewCommentsService(postgres.NewCommentsRepository(dataBase.Connection))
	router.GET("/api/v1/bookComments/:book_id",commentsService.BookCommentsHandler)
	router.POST("/api/v1/addBookComment/:book_id",commentsService.AddBookCommentHandler)
	router.GET("/api/v1/allCommentsByNickname/:nickname",commentsService.AllCommentsByNicknameHandler)

	router.GET("api/v1/tag/:book_id", bookService.getTags)
}

func initTagsRoutes(){

	tagsService  := NewTagsService(postgres.NewBooksRepository(dataBase.Connection), postgres.NewTagsRepository(dataBase.Connection))
	router.GET("/api/v1/tags", tagsService.getTags)
}