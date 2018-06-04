package services

import (
	"os"
	"log"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/metalscreame/GoToBoox/src/services/authentification/midlware"
	"github.com/metalscreame/GoToBoox/src/dataBase/repository/users"
	"github.com/metalscreame/GoToBoox/src/dataBase"
	"github.com/metalscreame/GoToBoox/src/dataBase/postgres"
)

const (
	apiRoute = "/api/v1"
)

type UserService struct {
	UsersRepo users.UserRepository
}

func NewUserService(repository users.UserRepository) *UserService {
	return &UserService{
		UsersRepo: repository,
	}
}

var router *gin.Engine

func InitializeRouter() {
	//Used for heroku
	port := os.Getenv("PORT")

	//Uncomment for local machine   !!!!
	//port = "8080"

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	gin.SetMode(gin.ReleaseMode)
	router = gin.New()
	router.Use(gin.Logger())
	//router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "/")
	router.LoadHTMLGlob("templates/*.html")


	router.GET("/", func(c *gin.Context) {
		isLoggedIn := midlware.CheckLoggedIn(c)
			if !isLoggedIn{
				guest := true
				c.HTML(http.StatusOK, "index.tmpl.html", guest)
			}else{
				c.HTML(http.StatusOK, "index.tmpl.html", nil)
			}
		})


	//The place for handlers routes
	// exm router.GET("/", showIndex.ShowIndexPage)
	router.GET(apiRoute, IndexHandler)
	initUserProfileRouters()

	router.Run(":" + port)
}

func initUserProfileRouters() {

	// Use the SetUserStatus middleware for every route to set a flag
	// indicating whether the request was from an authenticated user or not
	router.Use(midlware.SetUserStatus())

	service := NewUserService(postgres.NewPostgresUsersRepo(dataBase.Connection))
	userRoutes := router.Group(apiRoute)
	{
		// Handle POST requests at /api/v1/login
		// Ensure that the user is not logged in by using the middleware
		userRoutes.POST("/login", midlware.EnsureNotLoggedIn(), service.PerformLoginHandler)

		// Handle GET requests at /api/v1/logout
		// Ensure that the user is logged in by using the middleware
		userRoutes.GET("/logout", midlware.EnsureLoggedIn(), service.LogoutHandler)

		// Handle POST requests at /api/v1/register
		// Ensure that the user is not logged in by using the middleware
		userRoutes.POST("/register", midlware.EnsureNotLoggedIn(), service.UserCreateHandler)

		// Handle the GET requests at /api/v1/userProfile
		// Show the user's profile page
		// Ensure that the user is logged in by using the middleware
		userRoutes.GET("/userProfile", midlware.EnsureLoggedIn(), service.UserGetHandler)

		// Handle the GET requests at /api/v1/register
		// Show the user's profile page
		// Ensure that the user is logged in by using the middleware
		userRoutes.PUT("/userProfile", midlware.EnsureLoggedIn(), service.UserUpdateHandler)

		// Handle the GET requests at /api/v1/userProfile
		// Show the user's profile page
		// Ensure that the user is logged in by using the middleware
		userRoutes.DELETE("/userProfile", midlware.EnsureLoggedIn(), service.UserDeleteHandler)
	}

	// Show the login page
	// Ensure that the user is not logged in by using the middleware
	router.GET("/login", midlware.EnsureNotLoggedIn(), ShowLoginPage)

	// Show the registration page
	// Ensure that the user is not logged in by using the middleware
	router.GET("/register", midlware.EnsureNotLoggedIn(), ShowRegistrPage)

	// Show the user's profile page
	// Ensure that the user is logged in by using the middleware
	router.GET("/userProfile", midlware.EnsureLoggedIn(), ShowUsersProfilePage)
}
