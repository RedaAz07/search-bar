package main

import (
	"fmt"
	"net/http"
	"text/template"

	"groupie/routes"
	tools "groupie/tools"
)

func main() {
	var err error
	// parse all the html file from the template folder to variable Tp
	tools.Tp, err = template.ParseGlob("template/*.html")
	if err != nil {
		fmt.Println("Error parsing templates: ", err)
		return
	}

	// Register handlers
	routes.Route()
	//  run   the serve 
	fmt.Println("Server running at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
