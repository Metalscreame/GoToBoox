package services

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
)

func TestStaticPages(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	testCases := []struct {
		name          string
		handlerFunc   func(*gin.Context)
		needRedirect  bool
		needLoggedOff bool
	}{
		{
			name:        "SearchHandler",
			handlerFunc: SearchHandler,
		},
		{
			name:        "ShowLoginPage",
			handlerFunc: ShowLoginPage,
		},
		{
			name:        "ShowRegistrPage",
			handlerFunc: ShowRegistrPage,
		},
		{
			name:         "UserProfileHandler",
			handlerFunc:  UserProfileHandler,
			needRedirect: true,
		},
		{
			name:          "UserProfileHandler logged off",
			handlerFunc:   UserProfileHandler,
			needRedirect:  true,
			needLoggedOff: true,
		},
		{
			name:        "ShowUsersProfilePage",
			handlerFunc: ShowUsersProfilePage,
		},
		{
			name:        "ShowBook",
			handlerFunc: ShowBook,
		},
		{
			name:        "ShowUploadBookPage",
			handlerFunc: ShowUploadBookPage,
		},
		{
			name:        "ShowTakenBooksPage",
			handlerFunc: ShowTakenBooksPage,
		},
		{
			name:        "ShowCommentsPage",
			handlerFunc: ShowCommentsPage,
		},
		{
			name:        "SearchHandler",
			handlerFunc: SearchHandler,
		},
		{
			name:        "LocationHandler",
			handlerFunc: LocationHandler,
		},
	}

	t.Parallel()
	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			router := getTestRouter()
			router.GET("/staticTest", testCase.handlerFunc)

			req, _ := http.NewRequest("GET", "/staticTest", nil)

			if !testCase.needLoggedOff {
				req.AddCookie(&http.Cookie{Name: "is_logged_in", Value: "true"})
			}

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			if rr.Code != http.StatusOK && !testCase.needRedirect {
				t.Errorf("handler returned unexpected status: \n wanted: %v\n but got %v", http.StatusOK, rr.Code)
			} else if rr.Code != http.StatusFound && testCase.needRedirect {
				t.Errorf("handler returned unexpected status: \n wanted: %v\n but got %v", http.StatusOK, rr.Code)
			}
		})
	}

}

func BenchmarkShowTakenBooksPage(b *testing.B) {
	router := getTestRouter()
	router.GET("/staticTest", ShowTakenBooksPage)
	req, _ := http.NewRequest("GET", "/staticTest", nil)

	for n := 0; n < b.N; n++ {
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			b.Errorf("handler returned unexpected status: \n wanted: %v\n but got %v", http.StatusOK, rr.Code)
		}
	}
}

func BenchmarkSearchHandler(b *testing.B) {
	router := getTestRouter()
	router.GET("/staticTest", SearchHandler)
	req, _ := http.NewRequest("GET", "/staticTest", nil)

	for n := 0; n < b.N; n++ {
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		if rr.Code != http.StatusOK {
			b.Errorf("handler returned unexpected status: \n wanted: %v\n but got %v", http.StatusOK, rr.Code)
		}
	}
}

func getTestRouter() *gin.Engine {
	r := gin.Default()
	path := filepath.ToSlash(os.Getenv("GOPATH"))
	r.LoadHTMLGlob(path + "/src/github.com/metalscreame/GoToBoox/templates/*")
	return r
}
