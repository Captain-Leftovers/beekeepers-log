package main

import (
	"fmt"
	"net/http"

	"github.com/Captain-Leftovers/gohtmxtemplbeelog/view/layout"
	"github.com/a-h/templ"
)

func main() {

	mainRouter := http.NewServeMux()

	mainRouter.Handle("/", templ.Handler(layout.Base()))

	fmt.Println("Server is running at http://localhost:3000")

	http.ListenAndServe(":3000", mainRouter)

}
