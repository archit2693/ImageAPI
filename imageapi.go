package main

import(
	"net/http"
	"log"
	"os"
)

func DoesImageExist(image_name string) bool {
     _, error := os.Stat(image_name)
     if !os.IsNotExist(error) {
       return  true
     } else {
       return false
     }
}

func ImageHandler(writer http.ResponseWriter, request *http.Request){
    image_name, name_status := request.URL.Query()["image_name"]
    image_width, width_status := request.URL.Query()["width"]
    image_height, height_status := request.URL.Query()["height"]
    if !name_status || !width_status || !height_status || len(image_name) < 1 || len(image_width) < 1 || len(image_height) < 1 {
       ReportError(writer, "Sorry, Invalid Parameters.", 500)
       return
    }

   // requested_image_name = image_name[0]
    image_name[0] = "images/"+image_name[0]
    if !DoesImageExist(image_name[0]) {
     	ReportError(writer, "Sorry, the image you requested is not found.", 404)
	return
    }
    

}

func main(){
   http.HandleFunc("/api/resize", ImageHandler)
   log.Fatal(http.ListenAndServe(":8000", nil))
   
}