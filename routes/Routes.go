package routes

import (
	"net/http"

	"groupie/handler"
)

func Route() {
	http.HandleFunc("/static/", handler.StyleFunc)
	http.HandleFunc("/details/", handler.DetailsFunc)
	http.HandleFunc("/", handler.GroupieFunc)
}
