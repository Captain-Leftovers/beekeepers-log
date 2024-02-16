package main

import (
	"fmt"
	"net/http"

	"github.com/Captain-Leftovers/gohtmxtemplbeelog/internal/handler"
	"github.com/Captain-Leftovers/gohtmxtemplbeelog/view/component"
	"github.com/a-h/templ"
)

func main() {

	mainRouter := http.NewServeMux()

	homeHandler := &handler.HomeHandler{}

	mainRouter.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	mainRouter.HandleFunc("/", homeHandler.HandleShowHome)

	mainRouter.Handle("GET /greet", templ.Handler(component.HomeGreeting("Captain Leftovers")))

	fmt.Println("Server is running at http://localhost:3000")

	http.ListenAndServe(":3000", mainRouter)

}
