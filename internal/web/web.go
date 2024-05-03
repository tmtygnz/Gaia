package web

import (
	"gaia/utils"
	"log"
	"net/http"
)

func handleStatic() {
	fs := http.FileServer(http.Dir("./dist"))
	http.Handle("/dist/", http.StripPrefix("/dist/", fs))
}

func WebInterface() {
	log.Println("Running web server at :8081")

	handleStatic()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := utils.RenderPage("index", "./internal/web/views/index.html", w, struct{ Count string }{Count: "Hello World"})
		log.Println(err)
	})

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Println("Web server didn't start", err)
		return
	}
}
