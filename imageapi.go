package main

import(
	"net/http"
	"log"
)

func ImageHandler(w http.ResponseWriter, r *http.Request){
     
}

func main(){
   http.HandleFunc("/api/resize", ImageHandler)
   log.Fatal(http.ListenAndServe(":8000", nil))
   
}