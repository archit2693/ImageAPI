package main

import(
	"testing"
	"net/http"
	"net/http/httptest"
)

func init(){
     http.HandleFunc("/api/resize", ImageHandler)
}

func TestCorrectResponse(t *testing.T){
     req := httptest.NewRequest(http.MethodGet, "/api/resize", nil)
     rw := httptest.NewRecorder()
     http.DefaultServeMux.ServeHTTP(rw, req)
     if rw.Code != 200 {
       t.Fatalf("Expected 200 response code, but got: %v\n", rw.Code)
     }
}