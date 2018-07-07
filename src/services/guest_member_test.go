package services

import (
	"testing"
	"net/http"
	"bytes"
	"net/http/httptest"
	"github.com/golang/mock/gomock"
	"github.com/gin-gonic/gin"
	"github.com/metalscreame/GoToBoox/src/mocks"
	"github.com/metalscreame/GoToBoox/src/dataBase/repository"
	"errors"
	"time"
	"encoding/json"
)

const (
	regularInputBody = `{"nickname":"s","email":"email@email.com","password":"pass","new_passwordd":"1","has_book_for_exchange":false,"notification_get_new_books":false,"notification_get_when_book_reserved":false,"notification_daily":false,"taken_book_id":0,"role":""}`
	emailString      = "email@email.com"
)

func TestUserService_UserGetHandler(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	testCases := []struct {
		name            string
		inputBody       string
		expResponseBody string
		needError       bool
		needEmailCookie bool
		needServerError bool
	}{
		{
			name:            "regular",
			inputBody:       ``,
			expResponseBody: `{"nickname":"","email":"","password":"","new_passwordd":"","has_book_for_exchange":false,"notification_get_new_books":false,"notification_get_when_book_reserved":false,"notification_daily":false,"taken_book_id":0,"role":""}`,
			needError:       false,
			needEmailCookie: true,
		},
		{
			name:            "error no email cookie",
			inputBody:       ``,
			needError:       true,
			needEmailCookie: false,
		},
		{
			name:            "internal error",
			inputBody:       ``,
			needError:       true,
			needEmailCookie: false,
			needServerError: true,
		},
	}

	t.Parallel()
	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockUsersRepo := mocks.NewMockUserRepository(mockCtrl)
			mockService := NewUserService(mockUsersRepo)

			router := gin.New()
			router.GET("/getUser", mockService.UserGetHandler)

			requestBody := bytes.NewReader([]byte(testCase.inputBody))
			req, err := http.NewRequest("GET", "/getUser", requestBody)
			if err != nil {
				t.Fatal(err)
			}

			if testCase.needEmailCookie {
				req.AddCookie(&http.Cookie{Name: "email", Value: "email%40email.com"})
				mockUsersRepo.EXPECT().GetUserByEmail("email@email.com").Return(repository.User{}, nil)
			}

			if testCase.needServerError {
				req.AddCookie(&http.Cookie{Name: "email", Value: "email%40email.com"})
				mockUsersRepo.EXPECT().GetUserByEmail("email@email.com").Return(repository.User{}, errors.New("needed error"))
			}

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusOK && !testCase.needError {
				t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
			}

			result := rr.Body.String()
			if result != testCase.expResponseBody && !testCase.needError {
				t.Errorf("handler returned unexpected body: \n wanted: %v\n but got %v", testCase.expResponseBody, result)
			}
		})

	}
}

func TestUserService_UserDeleteHandler(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	testCases := []struct {
		name            string
		inputBody       string
		expResponseBody string
		needError       bool
		needEmailCookie bool
		needServerError bool
	}{
		{
			name:            "regular",
			expResponseBody: `{"status":"ok"}`,
			needError:       false,
			needEmailCookie: true,
		},
		{
			name:            "error no email cookie",
			needError:       true,
			needEmailCookie: false,
		},
		{
			name:            "internal error",
			needError:       true,
			needEmailCookie: false,
			needServerError: true,
		},
	}

	t.Parallel()
	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockUsersRepo := mocks.NewMockUserRepository(mockCtrl)
			mockService := NewUserService(mockUsersRepo)

			router := gin.New()
			router.GET("/deleteUser", mockService.UserDeleteHandler)

			requestBody := bytes.NewReader([]byte(testCase.inputBody))
			req, err := http.NewRequest("GET", "/deleteUser", requestBody)
			if err != nil {
				t.Fatal(err)
			}

			if testCase.needEmailCookie {
				req.AddCookie(&http.Cookie{Name: "email", Value: "email%40email.com"})
				mockUsersRepo.EXPECT().DeleteUserByEmail("email@email.com").Return(nil)
			}

			if testCase.needServerError {
				req.AddCookie(&http.Cookie{Name: "email", Value: "email%40email.com"})
				mockUsersRepo.EXPECT().DeleteUserByEmail("email@email.com").Return(errors.New("needed error"))
			}

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusOK && !testCase.needError {
				t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
			}

			result := rr.Body.String()
			if result != testCase.expResponseBody && !testCase.needError {
				t.Errorf("handler returned unexpected body: \n wanted: %v\n but got %v", testCase.expResponseBody, result)
			}
		})
	}
}

func TestUserService_UserUpdateHandler(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	mockedUser := repository.User{1, "s", "email@email", "1a1dc91c907325c69271ddf0c944bc72", "passs",
		0, false, 0, false,
		false, false, time.Now(),
		1, 0, ""}

	testCases := []struct {
		name                           string
		inputBody                      string
		expResponseBody                string
		needError                      bool
		needEmailCookie                bool
		needServerErrorFromUpdateRepo  bool
		needServerErrorFromGetUserRepo bool
		needUpdateRepo                 bool
		user                           repository.User
	}{
		{
			name:            "regular",
			inputBody:       regularInputBody,
			expResponseBody: `{"status":"ok"}`,
			needEmailCookie: true,
			needUpdateRepo:  true,
			user:            mockedUser,
		},
		{
			name:      "error no email cookie",
			inputBody: regularInputBody,
			needError: true,
			user:      mockedUser,
		},
		{
			name:                          "internal error from repo",
			inputBody:                     regularInputBody,
			needError:                     true,
			needEmailCookie:               true,
			needServerErrorFromUpdateRepo: true,
			user:                          mockedUser,
		},
		{
			name:            "passwords doesnt mach",
			inputBody:       `{"nickname":"s","email":"email@email.com","password":"wrong pass","new_passwordd":"pass","has_book_for_exchange":false,"notification_get_new_books":false,"notification_get_when_book_reserved":false,"notification_daily":false,"taken_book_id":0,"role":""}`,
			needError:       true,
			needEmailCookie: true,
			user:            mockedUser,
		},
		{
			name:      "bad request, wrong json",
			inputBody: `{"somethign":"asd"}`,
			needError: true,
		},
		{
			name:                           "bad request, coudnt get user from db",
			inputBody:                      regularInputBody,
			needError:                      true,
			needServerErrorFromGetUserRepo: true,
		},
	}

	t.Parallel()
	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockUsersRepo := mocks.NewMockUserRepository(mockCtrl)
			mockService := NewUserService(mockUsersRepo)

			router := gin.New()
			router.PUT("/update", mockService.UserUpdateHandler)

			requestBody := bytes.NewReader([]byte(testCase.inputBody))
			req, err := http.NewRequest("PUT", "/update", requestBody)
			if err != nil {
				t.Fatal(err)
			}

			if testCase.needEmailCookie {
				req.AddCookie(&http.Cookie{Name: "email", Value: "email%40email.com"})
				mockUsersRepo.EXPECT().GetUserByEmail(emailString).Return(testCase.user, nil)
			}

			if testCase.needUpdateRepo {
				var u repository.User
				json.Unmarshal([]byte(testCase.inputBody), &u)
				u.NewPassword = "c4ca4238a0b923820dcc509a6f75849b"
				mockUsersRepo.EXPECT().UpdateUserByEmail(u, emailString).Return(nil)
			}

			if testCase.needServerErrorFromGetUserRepo {
				req.AddCookie(&http.Cookie{Name: "email", Value: "email%40email.com"})
				mockUsersRepo.EXPECT().GetUserByEmail(emailString).Return(testCase.user, errors.New("test error"))
			}

			if testCase.needServerErrorFromUpdateRepo {
				var u repository.User
				json.Unmarshal([]byte(testCase.inputBody), &u)
				u.NewPassword = "c4ca4238a0b923820dcc509a6f75849b"
				mockUsersRepo.EXPECT().UpdateUserByEmail(u, emailString).Return(errors.New("test error"))
			}

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusOK && !testCase.needError {
				t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
			}

			result := rr.Body.String()
			if result != testCase.expResponseBody && !testCase.needError {
				t.Errorf("handler returned unexpected body: \n wanted: %v\n but got %v", testCase.expResponseBody, result)
			}
		})
	}
}

func TestUserService_UserCreateHandler(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	mockedUser := repository.User{0, "s", "email@email", "1a1dc91c907325c69271ddf0c944bc72", "1",
		0, false, 0, false,
		false, false, time.Now(),
		1, 0, ""}

	testCases := []struct {
		name                           string
		inputBody                      string
		expResponseBody                string
		needError                      bool
		needEmailCookie                bool
		needServerErrorFromInsertRepo  bool
		needServerErrorFromGetUserRepo bool
		needUpdateRepo                 bool
		user                           repository.User
	}{
		//{
		//	name:      "bad request, wrong json",
		//	inputBody: `{"somethign":"asd"}`,
		//	needError: true,
		//},
		{
			name:                          "could'nt insert user",
			inputBody:                     regularInputBody,
			needError:                     true,
			needServerErrorFromInsertRepo: true,
			user:                          mockedUser,
		},
		//{
		//	name:            "regular",
		//	inputBody:       regularInputBody,
		//	expResponseBody: `{"status":"ok"}`,
		//	needEmailCookie: true,
		//	needUpdateRepo:  true,
		//	user:            mockedUser,
		//},
		//{
		//	name:      "error no email cookie",
		//	inputBody: regularInputBody,
		//	needError: true,
		//	user:      mockedUser,
		//},
		//{
		//	name:                          "internal error from repo",
		//	inputBody:                     regularInputBody,
		//	needError:                     true,
		//	needEmailCookie:               true,
		//	needServerErrorFromUpdateRepo: true,
		//	user:                          mockedUser,
		//},
		//{
		//	name:            "passwords doesnt mach",
		//	inputBody:       `{"nickname":"s","email":"email@email.com","password":"wrong pass","new_passwordd":"pass","has_book_for_exchange":false,"notification_get_new_books":false,"notification_get_when_book_reserved":false,"notification_daily":false,"taken_book_id":0,"role":""}`,
		//	needError:       true,
		//	needEmailCookie: true,
		//	user:            mockedUser,
		//},
		//{
		//	name:      "bad request, wrong json",
		//	inputBody: `{"somethign":"asd"}`,
		//	needError: true,
		//},
		//{
		//	inputBody:                      regularInputBody,
		//	needError:                      true,
		//	needServerErrorFromGetUserRepo: true,
		//},
	}

	t.Parallel()
	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockUsersRepo := mocks.NewMockUserRepository(mockCtrl)
			mockService := NewUserService(mockUsersRepo)

			router := gin.New()
			router.POST("/createTest", mockService.UserCreateHandler)

			requestBody := bytes.NewReader([]byte(testCase.inputBody))
			req, err := http.NewRequest("POST", "/createTest", requestBody)
			if err != nil {
				t.Fatal(err)
			}

			if testCase.needEmailCookie {
				req.AddCookie(&http.Cookie{Name: "email", Value: "email%40email.com"})
				mockUsersRepo.EXPECT().GetUserByEmail(emailString).Return(testCase.user, nil)
			}

			if testCase.needUpdateRepo {
				var u repository.User
				json.Unmarshal([]byte(testCase.inputBody), &u)
				u.NewPassword = "c4ca4238a0b923820dcc509a6f75849b"
				mockUsersRepo.EXPECT().UpdateUserByEmail(u, emailString).Return(nil)
			}

			if testCase.needServerErrorFromGetUserRepo {
				req.AddCookie(&http.Cookie{Name: "email", Value: "email%40email.com"})
				mockUsersRepo.EXPECT().GetUserByEmail(emailString).Return(testCase.user, errors.New("test error"))
			}

			if testCase.needServerErrorFromInsertRepo {
				mockUsersRepo.EXPECT().InsertUser(testCase.user).Return(1,errors.New("test error"))
			}

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusOK && !testCase.needError {
				t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
			}

			result := rr.Body.String()
			if result != testCase.expResponseBody && !testCase.needError {
				t.Errorf("handler returned unexpected body: \n wanted: %v\n but got %v", testCase.expResponseBody, result)
			}
		})
	}
}

func TestUserService_LogoutHandler(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUsersRepo := mocks.NewMockUserRepository(mockCtrl)
	mockService := NewUserService(mockUsersRepo)

	router := gin.New()
	router.GET("/logout", mockService.LogoutHandler)

	requestBody := bytes.NewReader([]byte(""))
	req, err := http.NewRequest("GET", "/logout", requestBody)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
