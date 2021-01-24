package main

import (
	"log"
	// "fmt"
	"html/template"
	"net/http"
	"time"
)

func loadTemplate(name string) *template.Template {
	// t, err := template.ParseFiles("./templates/_header.html", "./templates/"+name+".html", "./templates/_footer.html")
	t, err := template.ParseFiles("templates/"+name+".html", "templates/_header.html", "templates/_footer.html")

	if err != nil {
		log.Fatalf("Not found file: %v", err)
	}

	return t
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	t := loadTemplate("index")

	if err := t.Execute(w, struct {
		Title string
		Message string
		Time time.Time
	}{
		Title: "やっぱりテストしたくない",
		Message: "Hello world",
		Time: time.Now(),
	}); err != nil {
		log.Printf("failed to execute template:", err)
	}
}

func singleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		lat := r.Form.Get("lat")
		lng := r.Form.Get("lng")

		t := loadTemplate("single")

		if err := t.Execute(w, struct {
			Title string
			Message string
			Time time.Time
		}{
			Title: "singleのテスト",
			Message: "lat: "+lat+" lng: "+lng,
			Time: time.Now(),
		}); err != nil {
			log.Printf("failed to execute template:", err)
		}
	}
}

func main() {
	http.HandleFunc("/", testHandler)
	http.HandleFunc("/single", singleHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("error starting http server: ", err)
		return
	}
}