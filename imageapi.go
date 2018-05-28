package main

import(
	"net/http"
	"log"
	"os"
	"strconv"	
 	"bytes"
	"image"
	"image/jpeg"
	"image/png"
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

    requested_image_name := image_name[0]
    image_name[0] = "images/"+image_name[0]
    if !DoesImageExist(image_name[0]) {
     	ReportError(writer, "Sorry, the image you requested is not found.", 404)
	return
    }
    width, width_conversion_status := strconv.Atoi(image_width[0])
    height, height_conversion_status := strconv.Atoi(image_height[0])
     if height_conversion_status != nil || width_conversion_status != nil {
     	ReportError(writer, "Sorry, there is something wrong with the parameters provided.", 500)
	return
     }
     if width > 1000 || height > 1000 {
     	ReportError(writer, "Sorry, your requested quite a big image.", 500)
	return
     }
     image_file, _ := os.Open(image_name[0])
     defer image_file.Close()
     image_type := getImageFormat(image_file)
     final_file := "images/cached/" + requested_image_name + "_final_" + image_width[0] + "_" + image_height[0] + "." + image_type
   
     if DoesImageExist(final_file) {
     	log.Println("File Cached Already.")
    	DecodeImage(writer, final_file, image_type)
    	return
     } else {
       log.Println("File Not Cached")
     }
     ResizeImage(writer, final_file, image_name[0], width, height, image_type)
}

func DecodeImage(writer http.ResponseWriter, final_file string, image_type string) {
  final_image_file, _ := os.Open(final_file)
  defer final_image_file.Close()
  src, _, _ := image.Decode(final_image_file)
  WriteImage(writer, &src, image_type)
}


func getImageFormat(file *os.File) (string) {
  bytes := make([]byte, 4)
  file.ReadAt(bytes, 0)
  if bytes[0] == 0x89 && bytes[1] == 0x50 && bytes[2] == 0x4E && bytes[3] == 0x47 { return "png" }
  if bytes[0] == 0xFF && bytes[1] == 0xD8 { return "jpg" }
  return ""
}


func WriteImage(writer http.ResponseWriter, image_to_process *image.Image, image_type string) {
  buffer := new(bytes.Buffer)
  switch image_type {
    case "jpg" :
     	  jpeg.Encode(buffer, *image_to_process, nil)
     	  writer.Header().Set("Content-Type", "image/jpeg")
    case "png" :
     	  png.Encode(buffer, *image_to_process)
       	  writer.Header().Set("Content-Type", "image/png")
        }
   writer.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
   writer.Write(buffer.Bytes())
}

func main(){
   http.HandleFunc("/api/resize", ImageHandler)
   log.Fatal(http.ListenAndServe(":8000", nil))
   
}