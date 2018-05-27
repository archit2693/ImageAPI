package main

import(
	"net/http"
	"log"
)

func ImageHandler(writer http.ResponseWriter, request *http.Request){
    image_name, name_status := request.URL.Query()["image_name"]
    image_width, width_status := request.URL.Query()["width"]
    image_height, height_status := request.URL.Query()["height"]
    if !name_status || !width_status || !height_status || len(image_name) < 1 || len(image_width) < 1 || len(image_height) < 1 {
       ReportError(writer, "Sorry, Invaild Parameters.", 500)
    }
}

func main(){
   http.HandleFunc("/api/resize", ImageHandler)
   log.Fatal(http.ListenAndServe(":8000", nil))
   
}