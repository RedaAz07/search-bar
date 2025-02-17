package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"

	"groupie/helpers"
	tools "groupie/tools"
)

var Artists []tools.Artists
type Result struct {
	Artist        []tools.Artists
	SearchElement map[string]string
}
var  SearchArtist = make(map[string]string)

func Groupie_Func(w http.ResponseWriter, r *http.Request) {
//	var wg sync.WaitGroup
	var wu sync.Mutex

	// check the path
	if r.URL.Path != "/" {
		// execute the not found  template
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorNotFound, http.StatusNotFound)
		return
	}
	// check the methd
	if r.Method != http.MethodGet {
		// execute the not found  template
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorMethodnotAll, http.StatusMethodNotAllowed)
		return
	}
	url := "https://groupietrackers.herokuapp.com/api/artists"
	// get the api data
	res, err := http.Get(url)
	if err != nil {
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorInternalServerErr, http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()
	// decode the jsone data
	err = json.NewDecoder(res.Body).Decode(&Artists)
	if err != nil {
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorInternalServerErr, http.StatusInternalServerError)
		return
	}

	// add  name
	for _, structs := range Artists {
		SearchArtist[structs.Name+"#artist/band name"] = "artist/band name"
		SearchArtist[structs.FirstAlbum+"#first album date"] = "first album date"
		SearchArtist[strconv.Itoa(structs.CreationDate)+"#creation date"] = "creation date"

		for _, v := range structs.Members {
			SearchArtist[v+"#members"] = "members"
		}
// !  i use  goroutin , mutex 
		go func(locId string) {
			Loc := &tools.Locations{}
			helpers.Fetch_By_Id(locId, Loc)
			wu.Lock()
			for _, Location := range Loc.Locations {
				SearchArtist[Location+"#locations"] = "locations"
			}
			wu.Unlock()	
		}(structs.Locations)

	}
	// ! end 
	lresult := Result{
		Artist:        Artists,
		SearchElement: SearchArtist,
	}

	helpers.RenderTemplates(w, "index.html", lresult, 200)
}
