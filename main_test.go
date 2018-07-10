package main

import (
	"testing"
	"github.com/metalscreame/GoToBoox/src/services"
	"net/http"
	"bytes"
	"net/http/httptest"
	"github.com/gin-gonic/gin"
	"time"
)

func TestRouter(t *testing.T)  {
	services.Start()

	time.Sleep(time.Second*10)
	req, _ := http.NewRequest("GET", "/serverStatus", bytes.NewReader([]byte("")))

	rr := httptest.NewRecorder()
	router := gin.New()
	router.ServeHTTP(rr, req)

	result := rr.Body.String()
	status := rr.Code;
	if result != `{"status":"alive"}` && status != http.StatusOK {
		t.Errorf("handler returned unexpected body: \n wanted: %v\n but got %v",`{"status":"alive"}`, result)
	}
}