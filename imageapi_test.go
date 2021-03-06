package main

import(
	"testing"
	"net/http"
	"net/http/httptest"
	"strings"
	"fmt"
)

func init(){
     http.HandleFunc("/api/resize", ImageHandler)
}

func TestCorrectResponse(t *testing.T){
     request := httptest.NewRequest(http.MethodGet, "/api/resize", nil)
     response_recorder := httptest.NewRecorder()
     http.DefaultServeMux.ServeHTTP(response_recorder, request)
     if response_recorder.Code != 200 {
       t.Fatalf("Expected 200 response code, but got: %v\n", response_recorder.Code)
     }
}

func TestParamtersPresence(t *testing.T){
     request := httptest.NewRequest(http.MethodGet, "/api/resize", nil)
     response_recorder := httptest.NewRecorder()
     http.DefaultServeMux.ServeHTTP(response_recorder, request)
     expected := "Sorry, Invalid Parameters."
     
     actual := response_recorder.Body.String()
      fmt.Println(actual)
     if !strings.Contains(actual, expected){
     	t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func TestFileAbsence(t *testing.T) {
     request := httptest.NewRequest(http.MethodGet, "/api/resize?image_name=test.jpg&width=600&height=500", nil)
     response_recorder := httptest.NewRecorder()
     http.DefaultServeMux.ServeHTTP(response_recorder, request)
     expected := "Sorry, the image you requested is not found."
     actual := response_recorder.Body.String()
     if !strings.Contains(actual, expected){
     	t.Errorf("handler returned unexpected body: got %v want %v ", actual, expected)
     }
}

func TestWidthandHeightType(t *testing.T) {
     request := httptest.NewRequest(http.MethodGet, "/api/resize?image_name=husky.jpeg&width=abc&height=500a", nil)
     response_recorder := httptest.NewRecorder()
     http.DefaultServeMux.ServeHTTP(response_recorder, request)
     expected := "Sorry, there is something wrong with the parameters provided."
     actual := response_recorder.Body.String()
     if !strings.Contains(actual, expected) {
     	t.Errorf("handler returned unexpected body: got %v want %v %v", actual, expected, response_recorder.Body.String())
     }
}

func TestWidthandHeightRange(t *testing.T) {
     request := httptest.NewRequest(http.MethodGet, "/api/resize?image_name=husky.jpeg&width=400&height=5000", nil)
     response_recorder := httptest.NewRecorder()
     http.DefaultServeMux.ServeHTTP(response_recorder, request)
     expected := "Sorry, your requested quite a big image"
     actual := strings.Contains(response_recorder.Body.String(),expected)
     if actual == false {
     	t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
     }
}

func TestFileResponseAndContentType(t *testing.T) {
     request := httptest.NewRequest(http.MethodGet, "/api/resize?image_name=husky.jpeg&width=500&height=400", nil)
     response_recorder := httptest.NewRecorder()
     http.DefaultServeMux.ServeHTTP(response_recorder, request)
     expected := "image/jpeg"
     actual := strings.Contains(response_recorder.Header().Get("Content-Type"),expected)
     if actual == false || len(response_recorder.Body.String()) < 1 {
     	t.Errorf("handler returned unexpected body: got %v want %v .. %v", actual, expected, response_recorder.Header().Get("Content-type"))
     }
}

func TestInvalidFileType(t *testing.T) {
     request := httptest.NewRequest(http.MethodGet, "/api/resize?image_name=test.txt&width=500&height=400", nil)
     response_recorder := httptest.NewRecorder()
     http.DefaultServeMux.ServeHTTP(response_recorder, request)
     expected := "Sorry, Invalid image type."
     actual := response_recorder.Body.String()
     if !strings.Contains(actual, expected) {
     	t.Errorf("handler returned unexpected body: got %v want %v v", actual, expected)
     }
}

func TestPngImages(t *testing.T) {
  request := httptest.NewRequest(http.MethodGet, "/api/resize?image_name=car.png&width=500&height=400", nil)
     response_recorder := httptest.NewRecorder()
     http.DefaultServeMux.ServeHTTP(response_recorder, request)
     expected := "image/png"
     actual := strings.Contains(response_recorder.Header().Get("Content-Type"),expected)
     if actual == false || len(response_recorder.Body.String()) < 1 {
     	t.Errorf("handler returned unexpected body: got %v want %v .. %v", actual, expected,response_recorder.Header().Get("Content-type"))
     }
}

func TestForCachedFile(t *testing.T) {
  request := httptest.NewRequest(http.MethodGet, "/api/resize?image_name=car.png&width=550&height=450", nil)
     response_recorder := httptest.NewRecorder()
     http.DefaultServeMux.ServeHTTP(response_recorder, request)
     expected := "image/png"
     actual := strings.Contains(response_recorder.Header().Get("Content-Type"),expected)
     filename := "images/cached/car.png_final_500_400.png"
     if actual == false || len(response_recorder.Body.String()) < 1 && DoesImageExist(filename) {
     	t.Errorf("handler returned unexpected body: got %v want %v .. %v", actual, expected,response_recorder.Header().Get("Content-type"))
     }
}
