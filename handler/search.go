package handler

import (
	"groupie/helpers"
	"groupie/tools"
	"net/http"
	"strconv"
	"strings"
)

func Search(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorMethodnotAll, http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorBadReq, http.StatusBadRequest)
		return
	}

	searchvalue := r.FormValue("search")

	for i := 0; i < len(searchvalue); i++ {
		if i-1 >= 0 && searchvalue[i] == '-' && searchvalue[i-1] == ' ' {
			searchvalue = searchvalue[:i]
			break
		}


	}
	searchvalue=strings.TrimSpace(strings.ToLower(searchvalue))
	

	var artistsData *[]tools.Artists
	err = helpers.Fetch("https://groupietrackers.herokuapp.com/api/artists", &artistsData)
	if err != nil {
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorInternalServerErr, http.StatusInternalServerError)
		return
	}
	var locations *tools.Index

	errLoc := helpers.Fetch("https://groupietrackers.herokuapp.com/api/locations", &locations)
	if errLoc != nil {
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorInternalServerErr, http.StatusInternalServerError)
	}
	var loc []string
	for _, v := range locations.Index {

		loc = append(loc, v.Locations...)

	}

	var FilterArt []tools.Artists

	for r, artist := range *artistsData {

		x := true
		str1 := strings.ToLower(artist.Name)
		if strings.Contains(str1, searchvalue) {

			FilterArt = append(FilterArt, artist)
			x = false

		}
		if x && searchvalue == artist.FirstAlbum {
			FilterArt = append(FilterArt, artist)
			x = false
		}
		if x && searchvalue == strconv.Itoa(artist.CreationDate) {
			FilterArt = append(FilterArt, artist)
			x = false
		}
		for _, j := range artist.Members {
			str2 := strings.ToLower(j)
			if x && strings.Contains(str2, searchvalue) {
				FilterArt = append(FilterArt, artist)
				x = false
			}
		}
		for _, a := range locations.Index {
			if a.ID == r+1 {
				for _, j := range a.Locations {
					str3 := strings.ToLower(j)

					if x && strings.Contains(str3, searchvalue) {
						FilterArt = append(FilterArt, artist)
						x = false
					}
				}
			}
		}

	}

	data := tools.Data{}
	Handle_data(artistsData, &data)
	data.Artists = &FilterArt

	helpers.RenderTemplates(w, "searchPage.html", data, 200)
}
