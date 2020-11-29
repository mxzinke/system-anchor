package main

import (
	"github.com/avct/uasurfer"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/:direction", handler)

	log.Println("Server is starting at port 80!")

	if err := http.ListenAndServe(":80", router); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Println("Redirecting for:", p.ByName("direction"))

	userAgent := uasurfer.Parse(r.UserAgent())
	log.Println(r.UserAgent(), "=>", userAgent.OS.Name.String())

	switch userAgent.OS.Name {
	case uasurfer.OSiOS:
		http.Redirect(w, r, "https://apps.apple.com/us/app/typy/id1538771852", 301)

	case uasurfer.OSAndroid:
		http.Redirect(w, r, "https://play.google.com/store/apps/details?id=de.typy", 301)

	default:
		log.Println("No possible direction found!")
		if _, err := w.Write([]byte("No possible direction found!")); err != nil {
			log.Println(err.Error())
		}
	}

}