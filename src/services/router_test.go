package services

import (
	"testing"
	"time"
	"net/http"
	"bytes"
	"net/http/httptest"
)

func TestStart(t *testing.T)  {
	TestCaseFlag=true
	go Start()

	time.Sleep(time.Second*3)
	req, _ := http.NewRequest("GET", "http://localhost:8080/serverStatus", bytes.NewReader([]byte("")))

	rr := httptest.NewRecorder()
	Router.ServeHTTP(rr, req)

	result := rr.Body.String()
	status := rr.Code
	if result != `{"status":"alive"}` && status != http.StatusOK {
		t.Errorf("handler returned unexpected body: \n wanted: %v\n but got %v",`{"status":"alive"}`, result)
	}
	Shutdown <- 1
}


