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
)

func TestUserService_UserGetHandler(t *testing.T) {

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
			expResponseBody: `{"nickname":"","email":"","password":"","new_passwordd":"","has_book_for_exchange":false,"notification_get_new_books":false,"notification_get_when_book_reserved":false,"notification_daily":false,"taken_book_id":0}`,
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
			needEmailCookie: true,
			needServerError: true,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase

		t.Run(testCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockUsersRepo := mocks.NewMockUserRepository(mockCtrl)
			mockService := NewUserService(mockUsersRepo)

			router := gin.New()
			router.GET("/getUser", mockService.UserGetHandler)

			req, err := http.NewRequest("GET", "/getUser", bytes.NewReader([]byte(testCase.inputBody)))
			if err != nil {
				t.Fatal(err)
			}

			if testCase.needEmailCookie {
				req.AddCookie(&http.Cookie{Name: "email", Value: "email%40email.com"})
				mockUsersRepo.EXPECT().GetUserByEmail("email@email.com").Return(repository.User{}, nil)
			}

			if testCase.needServerError{
				req.AddCookie(&http.Cookie{Name: "email", Value: "email%40email.com"})
				mockUsersRepo.EXPECT().GetUserByEmail("email@email.com").Return(repository.User{}, errors.New("err"))
			}

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusOK && !testCase.needError {
				t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
			}

			result := rr.Body.String()
			if result != testCase.expResponseBody && !testCase.needError {
				t.Errorf("handler returned unexpected body: %v", result)
			}
		})

	}
}
