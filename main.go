package main

import (
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/gobuffalo/packr"
)

func main() {
	box := packr.NewBox("./templates")
	t, err := template.New("hello").Parse(string(box.String("hello.txt")))
	if err != nil {
		log.Fatalln(err)
	}
	t.Execute(os.Stdout, "World")
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello Mars!")
	}))
	http.ListenAndServe(":80", nil)
}
