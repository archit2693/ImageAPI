package main

import (
       "gopkg.in/gographics/imagick.v2/imagick"
       "net/http"
)

func ResizeImage(writer http.ResponseWriter,final_file string, input_file string, width int, height int , image_type string) {

  imagick.Initialize()
  defer imagick.Terminate()
  magic_wand := imagick.NewMagickWand()
  defer magic_wand.Destroy()
  error := magic_wand.ReadImage(input_file)
  if error != nil {
    ReportError(writer,"Sorry, Invalid image type.", 500)
    return
  }

  error = magic_wand.ResizeImage(uint(width),uint(height), imagick.FILTER_LANCZOS, 1)
  
  if error == nil {
     error = magic_wand.SetImageCompressionQuality(100)
  }
  if error == nil {
     error = magic_wand.WriteImage(final_file)
     DecodeImage(writer, final_file, image_type)
  }
}