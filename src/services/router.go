package services

import (
	"github.com/gin-gonic/gin"

	"github.com/metalscreame/GoToBoox/src/services/midlwares"
	"github.com/metalscreame/GoToBoox/src/dataBase"
	"github.com/metalscreame/GoToBoox/src/dataBase/postgres"
	"log"
	"net/http"
	"gopkg.in/appleboy/gin-jwt.v2"
	"time"
	"os"
)

const (
	apiRoute = "/api/v1"
)

//Router is a global router variable
var Router *gin.Engine
var jwtMiddleware *jwt.GinJWTMiddleware

//Shutdown Is a channel to shutdown the router in runtime
var Shutdown chan int

//TestCaseFlag is a variable that used in router test. Is true when it's a test
var TestCaseFlag bool

//Start is a function that starts server and initializes all the routes.
func Start() {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		println("PORT is required\nFor localhosts setup sys env \"PORT\" as 8080 and reload IDE")
		log.Fatal("PORT is required\nFor localhosts setup sys env \"PORT\" as 8080 and reload IDE")
	}

	gin.SetMode(gin.ReleaseMode)
	Router = gin.Default()
	Router.Use(gin.Logger())
	Router.Use(gin.Recovery())
	if !TestCaseFlag {
		Router.Static("/static", "./static")
		Router.LoadHTMLGlob("templates/*.html")
	}

	Router.GET("/", IndexHandler)
	Router.GET("/location", LocationHandler)
	Router.GET("/search", SearchHandler)

	service := NewUserService(postgres.NewPostgresUsersRepo(dataBase.Connection))
	jwtMiddleware = &jwt.GinJWTMiddleware{

		Realm:         "Name",
		Key:           []byte(dataBase.TokenKeyLookUp()),
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour * 24,
		Authenticator: service.CheckCredentials,
		Authorizator:  service.Authorization,
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup: "cookie:token",
		//	TokenHeadName: "Bearer",
		TimeFunc:    time.Now,
		PayloadFunc: service.Payload,
	}

	Router.GET(apiRoute, ApIIndexHandler)
	initUserProfileRoutes()
	initBooksRoutes()
	initTagsRoutes()

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: Router,
	}
	Router.GET("/serverStatus", ServerIsOn)

	go func() {
		// Starting server
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	Shutdown = make(chan int)
	<-Shutdown
	log.Println("Shutdown Server ...")
}

func initUserProfileRoutes() {

	// Use the SetUserStatus middleware for every route to set a flag
	// indicating whether the request was from an authenticated user or not
	Router.Use(midlwares.SetUserStatus())

	userService := NewUserService(postgres.NewPostgresUsersRepo(dataBase.Connection))
	userRoutes := Router.Group(apiRoute)
	{

		// Handle POST requests at /api/v1/login
		// Ensure that the user is not logged in by using the middleware
		userRoutes.POST("/login", midlwares.EnsureNotLoggedIn(), jwtMiddleware.LoginHandler)

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

	auth := Router.Group("/auth")
	auth.Use(jwtMiddleware.MiddlewareFunc())
	{
		auth.GET("/vip", helloHandler)
		auth.GET("/refresh_token", jwtMiddleware.RefreshHandler)
	}

	// Show the login page
	// Ensure that the user is not logged in by using the middleware
	Router.GET("/login", midlwares.EnsureNotLoggedIn(), ShowLoginPage)

	// Show the registration page
	// Ensure that the user is not logged in by using the middleware
	Router.GET("/register/:bookid", midlwares.EnsureNotLoggedIn(), ShowRegistrPage)

	// Show the user's profile page or login page
	Router.GET("/userProfile", UserProfileHandler )

	//Shows the lock page
	Router.GET("/uploadPage/:book_id", midlwares.EnsureLoggedIn(), ShowUploadBookPage)

	// Show the user's profile page
	// Ensure that the user is logged in by using the middleware
	Router.GET("/userProfilePage", midlwares.EnsureLoggedIn(),midlwares.TokenChecking(), ShowUsersProfilePage)

	Router.GET("/userComments/:nickname", ShowCommentsPage)
}

func initBooksRoutes() {
	bookService := NewBookService(postgres.NewBooksRepository(dataBase.Connection), postgres.NewPostgresUsersRepo(dataBase.Connection))
	//get all books in certain category
	Router.GET("categories/:cat_id/books", bookService.getBooks)
	//get all books
	Router.GET("/books", bookService.showAllBooks)
	//get books by it's ID
	//Router.GET("/categories/:cat_id/book/:book_id", bookService.getBook)
	Router.GET("/books/m/mostPopularBooks", bookService.FiveMostPop)
	Router.POST("/api/v1/books/search", bookService.getBookBySearch)
	Router.GET("/api/v1/book/:book_id", bookService.getBook)

	Router.GET("/book/:book_id", ShowBook)
	Router.GET("/api/v1/books/taken", bookService.showAllTakenBooks)

	Router.GET("/api/v1/books/showReserved", bookService.ShowReservedBooksByUser)
	Router.GET("/api/v1/books/taken/0", bookService.ShowTakenBookByUser)
	Router.GET("/books/taken/:id", ShowTakenBooksPage)

	Router.POST("/api/v1/insertNewBook/:book_id", midlwares.EnsureLoggedIn(), bookService.InsertNewBook)

	Router.GET("/api/v1/updateBookStatus/:book_id", bookService.UpdateBookStatusToReturningFromTaken)
	Router.GET("/api/v1/updateBookStatusReturn/:book_id/:reserved_book_id", bookService.UpdateBookStatusToReturning)
	Router.GET("/api/v1/makeBookCross", bookService.ExchangeBook)

	commentsService := NewCommentsService(postgres.NewCommentsRepository(dataBase.Connection))
	Router.GET("/api/v1/bookComments/:book_id", commentsService.BookCommentsHandler)
	Router.POST("/api/v1/addBookComment/:book_id", commentsService.AddBookCommentHandler)
	Router.GET("/api/v1/allCommentsByNickname/:nickname", commentsService.AllCommentsByNicknameHandler)

	Router.GET("api/v1/tag/:book_id", bookService.getTags)
}

func initTagsRoutes() {

	tagsService := NewTagsService(postgres.NewBooksRepository(dataBase.Connection), postgres.NewTagsRepository(dataBase.Connection))
	Router.GET("/api/v1/tags", tagsService.getTags)
}

func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	c.JSON(200, gin.H{
		"userID": claims["id"],
		"text":   "Hello! U see this cause u'r vip.",
	})
}
