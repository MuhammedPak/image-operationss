package main

import (
	"image"
	"image/draw"
	"image/jpeg"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/nfnt/resize"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", connection)
	http.ListenAndServe(":4000", r)
}
func connection(w http.ResponseWriter, r *http.Request) {

	image_resize, err := ioutil.ReadDir("D:/image-api/photos")
	for _, item := range image_resize {
		if exist_for_resizing(item.Name()) == true {
			resizing(item.Name())
		}
	}

	imagelist, err := ioutil.ReadDir("D:/image-api/resizing") //Dosyaları okuyacağımız klasörün adı
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range imagelist {
		if exist(item.Name()) == true {
			waterimagesss(item.Name())
		}
	}

}

func waterimagesss(image_name string) {
	image1, err := os.Open("D:/image-api/resizing/" + image_name) //Gelen fotoğraf adına göre dosyadan çekiyoruz
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}

	first, err := jpeg.Decode(image1)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}
	defer image1.Close()

	image2, err := os.Open("D:/new.jpg") //Ekleyeceğimiz logonun adresi
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}
	second, err := jpeg.Decode(image2)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}
	defer image2.Close()

	offset := image.Pt(150, 150)
	b := first.Bounds()
	image3 := image.NewRGBA(b)
	draw.Draw(image3, b, first, image.ZP, draw.Src)
	draw.Draw(image3, second.Bounds().Add(offset), second, image.ZP, draw.Over)

	third, err := os.Create("D:/image-api/waterimages/" + image_name) //Logo yerleştirdiğimiz fotoğrafların yazılacağı klasör

	if err != nil {
		log.Fatalf("failed to create: %s", err)
	}

	jpeg.Encode(third, image3, &jpeg.Options{jpeg.DefaultQuality})
	defer third.Close()
}
func resizing(photo_name string) {
	file, err := os.Open("D:/image-api/photos/" + photo_name)
	if err != nil {
		log.Fatal(err)
	}

	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(500, 500, img, resize.Lanczos3)

	out, err := os.Create("D:/image-api/resizing/" + photo_name)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)
}
func exist(image string) bool {

	path := "D:/image-api/waterimages/" + image
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return true
	} else {
		return false
	}
}
func exist_for_resizing(image string) bool {

	path := "D:/image-api/resizing" + image
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return true
	} else {
		return false
	}
}
