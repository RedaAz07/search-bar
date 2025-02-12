package helpers

import (
	"bytes"
	"net/http"

	"groupie/tools"
)

func RenderTemplates(w http.ResponseWriter, temp string, post interface{} , status int) {
	var buf bytes.Buffer

	// exucut the template with buffer to chekc if there is an error in  our template 
	err := tools.Tp.ExecuteTemplate(&buf, temp, post)
	if err != nil {
		errore := tools.ErrorPage{
			Code:         http.StatusInternalServerError,
			ErrorMessage: "Something went wrong on our end. Please try again later.",
		}
		w.WriteHeader(http.StatusInternalServerError)
		tools.Tp.ExecuteTemplate(w, "statusPage.html", errore)
		return
	}
	w.WriteHeader(status)
	w.Write(buf.Bytes())
}
