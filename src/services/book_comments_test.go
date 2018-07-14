package services

import (
	"testing"
	"github.com/metalscreame/GoToBoox/src/mocks"
	"github.com/golang/mock/gomock"

	"bytes"
	"net/http"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"errors"
	"github.com/metalscreame/GoToBoox/src/dataBase/repository"
)

func TestCommentsService_BookCommentsHandler(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	testCases := []struct {
		name            string
		inputBody       string
		expResponseBody string
		needError       bool
		needParam       bool
		needRepoErr     bool
		needRepoNoErr   bool
		needBadParam    bool
	}{
		{
			name:      "need error, no param",
			needError: true,
		},
		{
			name:         "need error, bad param",
			needError:    true,
			needBadParam: true,
		},
		{
			name:        "need error, cant get",
			needError:   true,
			needParam:   true,
			inputBody:   `{"commentText":"comment"}`,
			needRepoErr: true,
		},
		{
			name:            "regular",
			needParam:       true,
			needRepoNoErr:   true,
			expResponseBody: `{"data":{"Comments":null},"status":200}`,
		},
	}

	t.Parallel()
	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mocCommentsRepo := mocks.NewMockCommentsRepository(mockCtrl)
			mockService := NewCommentsService(mocCommentsRepo)

			router := gin.New()
			router.GET("/getAll/:book_id", mockService.BookCommentsHandler)

			requestBody := bytes.NewReader([]byte(testCase.inputBody))

			var req *http.Request
			if testCase.needParam {
				req, _ = http.NewRequest("GET", "/getAll/1", requestBody)
			} else {
				req, _ = http.NewRequest("GET", "/getAll", requestBody)
			}

			if testCase.needBadParam {
				req, _ = http.NewRequest("GET", "/getAll/string", requestBody)
			}

			if testCase.needRepoErr {
				var commen []repository.Comment
				mocCommentsRepo.EXPECT().GetAllCommentsByBookID(gomock.Any()).Return(commen, errors.New("needed error"))
			}

			if testCase.needRepoNoErr {
				var commen []repository.Comment
				mocCommentsRepo.EXPECT().GetAllCommentsByBookID(gomock.Any()).Return(commen, nil)
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

func TestCommentsService_AllCommentsByNicknameHandler(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	testCases := []struct {
		name            string
		inputBody       string
		expResponseBody string
		needError       bool
		needParam       bool
		needRepoErr     bool
		needRepoNoErr   bool
	}{
		{
			name:      "need error, no param",
			needError: true,
		},
		{
			name:        "need error, cant get",
			needError:   true,
			needParam:   true,
			inputBody:   `{"commentText":"comment"}`,
			needRepoErr: true,
		},
		{
			name:            "regular",
			needParam:       true,
			needRepoNoErr:   true,
			expResponseBody: `{"data":{"Comments":null},"status":200}`,
		},
	}

	t.Parallel()
	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mocCommentsRepo := mocks.NewMockCommentsRepository(mockCtrl)
			mockService := NewCommentsService(mocCommentsRepo)

			router := gin.New()
			router.GET("/getAll/:nickname", mockService.AllCommentsByNicknameHandler)

			requestBody := bytes.NewReader([]byte(testCase.inputBody))

			var req *http.Request
			if testCase.needParam {
				req, _ = http.NewRequest("GET", "/getAll/1", requestBody)
			} else {
				req, _ = http.NewRequest("GET", "/getAll", requestBody)
			}

			if testCase.needRepoErr {
				var commen []repository.Comment
				mocCommentsRepo.EXPECT().GetAllCommentsByNickname(gomock.Any()).Return(commen, errors.New("needed error"))
			}

			if testCase.needRepoNoErr {
				var commen []repository.Comment
				mocCommentsRepo.EXPECT().GetAllCommentsByNickname(gomock.Any()).Return(commen, nil)
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

func TestCommentsService_AddBookCommentHandler(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	testCases := []struct {
		name            string
		inputBody       string
		expResponseBody string
		needError       bool
		needParam       bool
		needBadParam    bool
		needCookie      bool
		needRepoErr     bool
		needRepoNoErr   bool
	}{
		{
			name:      "need error, no param",
			needError: true,
		},
		{
			name:      "need error, cant parse json",
			needError: true,
			needParam: true,
			inputBody: `{"im bad":"bad","ssad"'}`,
		},
		{
			name:      "need error, empty comment",
			needError: true,
			needParam: true,
			inputBody: `{"commentText":""}`,
		},
		{
			name:         "need error, param is string",
			needError:    true,
			needBadParam: true,
			inputBody:    `{"commentText":"comment"}`,
		},
		{
			name:      "need error, no cookie",
			needError: true,
			needParam: true,
			inputBody: `{"commentText":"comment"}`,
		},
		{
			name:        "need error, cant insert",
			needError:   true,
			needParam:   true,
			needCookie:  true,
			inputBody:   `{"commentText":"comment"}`,
			needRepoErr: true,
		},
		{
			name:            "regular",
			needParam:       true,
			needCookie:      true,
			inputBody:       `{"commentText":"comment"}`,
			needRepoNoErr:   true,
			expResponseBody: `{"status":"ok"}`,
		},
	}

	t.Parallel()
	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mocCommentsRepo := mocks.NewMockCommentsRepository(mockCtrl)
			mockService := NewCommentsService(mocCommentsRepo)

			router := gin.New()
			router.POST("/addCmt/:book_id", mockService.AddBookCommentHandler)

			requestBody := bytes.NewReader([]byte(testCase.inputBody))

			var req *http.Request
			if testCase.needParam {
				req, _ = http.NewRequest("POST", "/addCmt/1", requestBody)
			} else {
				req, _ = http.NewRequest("POST", "/addCmt", requestBody)
			}

			if testCase.needBadParam {
				req, _ = http.NewRequest("POST", "/addCmt/string", requestBody)
			}

			if testCase.needCookie {
				req.AddCookie(&http.Cookie{Name: "email", Value: "email%40email.com"})
				req.AddCookie(&http.Cookie{Name: "nickname", Value: "nickname"})
			}

			if testCase.needRepoErr {
				mocCommentsRepo.EXPECT().InsertNewComment(gomock.Any(), gomock.Any(), gomock.Any(), 1).Return(errors.New("needed error"))
			}

			if testCase.needRepoNoErr {
				mocCommentsRepo.EXPECT().InsertNewComment(gomock.Any(), gomock.Any(), gomock.Any(), 1).Return(nil)
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

func BenchmarkCommentsService_BookCommentsHandler(b *testing.B) {
	gin.SetMode(gin.ReleaseMode)

	type BenchCase struct {
		name string

		expResponseBody string
		needError       bool
		needRepoErr     bool
		needRepoNoErr   bool
	}

	benchCase := BenchCase{
		name:            "regular",
		needRepoNoErr:   true,
		expResponseBody: `{"data":{"Comments":null},"status":200}`,
	}

	for n := 0; n < b.N; n++ {

		mockCtrl := gomock.NewController(b)
		defer mockCtrl.Finish()

		mocCommentsRepo := mocks.NewMockCommentsRepository(mockCtrl)
		mockService := NewCommentsService(mocCommentsRepo)

		router := gin.New()
		router.GET("/getAll/:book_id", mockService.BookCommentsHandler)

		var req *http.Request
		req, _ = http.NewRequest("GET", "/getAll/1", nil)
		if benchCase.needRepoErr {
			var commen []repository.Comment
			mocCommentsRepo.EXPECT().GetAllCommentsByBookID(gomock.Any()).Return(commen, errors.New("needed error"))
		}

		if benchCase.needRepoNoErr {
			var commen []repository.Comment
			mocCommentsRepo.EXPECT().GetAllCommentsByBookID(gomock.Any()).Return(commen, nil)
		}

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK && !benchCase.needError {
			b.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		result := rr.Body.String()
		if result != benchCase.expResponseBody && !benchCase.needError {
			b.Errorf("handler returned unexpected body: \n wanted: %v\n but got %v", benchCase.expResponseBody, result)
		}
	}
}
