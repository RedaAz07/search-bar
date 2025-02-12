package handler

import (
	"net/http"
	"os"
	"strings"

	tools "groupie/tools"
)

func StyleFunc(w http.ResponseWriter, r *http.Request) {
	filePath := strings.TrimPrefix(r.URL.Path, "/")
	File, err := os.Stat(filePath)
	//chekc if the file exists and if its a directory 
	if err != nil || File.IsDir() {

	

		w.WriteHeader(http.StatusNotFound)
		tools.Tp.ExecuteTemplate(w, "statusPage.html", tools.ErrorNotFound)
		return
	}
	// server the  style 
	http.StripPrefix("/static", http.FileServer(http.Dir("static"))).ServeHTTP(w, r)
}
