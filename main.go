package main

import (
	"fmt"
	"net/http"

	"github.com/MGM103/Mensaria/components"
)

var PORT = 8080

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		component := components.Base("Mensaria", components.Dashboard("Marcus"))
		component.Render(r.Context(), w)
	})
	mux.HandleFunc("/add-asset", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			component := components.Popup()
			component.Render(r.Context(), w)
		}
	})
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Server running on port:", PORT)
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), mux)
}
