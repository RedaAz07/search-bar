package routes

import (
	"fmt"
	"net/http"
	"text/template"

	"groupie/handler"
	"groupie/tools"
)

func Route() {
	var err error
	// parse all the html file from the template folder to variable Tp
	tools.Tp, err = template.ParseGlob("template/*.html")
	if err != nil {
		fmt.Println("Error parsing templates: ", err)
		return
	}
	http.HandleFunc("/search", handler.Search_Func)

	http.HandleFunc("/static/", handler.Style_Func)
	http.HandleFunc("/details/", handler.Detail_Card_Func)
	http.HandleFunc("/", handler.Groupie_Func)
}
