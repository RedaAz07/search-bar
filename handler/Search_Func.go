package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"groupie/helpers"
	"groupie/tools"
)

func Search_Func(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	var wu sync.Mutex
	if r.Method != http.MethodPost {
		// execute the not found  template
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorMethodnotAll, http.StatusMethodNotAllowed)
		return
	}
	var spliteInput []string
	spliteInput = strings.Split(r.FormValue("search"), "#")

	InputSearch := (strings.ToLower(spliteInput[0]))
	typrSearch := (strings.ToLower(spliteInput[1]))
	var foundId []int
	for _, v := range Artists {
		switch {
		case typrSearch == "artist/band name":
			if strings.ToLower(v.Name) == InputSearch {
				foundId = append(foundId, v.Id)
			}
		case typrSearch == "first album date":

			if strings.ToLower(v.FirstAlbum) == InputSearch {
				foundId = append(foundId, v.Id)
			}
		case typrSearch == "creation date":

			if strconv.Itoa(v.CreationDate) == InputSearch {
				foundId = append(foundId, v.Id)
			}

		case typrSearch == "members":

			for _, members := range v.Members {
				if strings.ToLower(members) == InputSearch {
					foundId = append(foundId, v.Id)
				}
			}
		case typrSearch == "locations":

			// !  i use  goroutin , mutex
			wg.Add(1)
			go func(locId string) {
				defer wg.Done()
				Loc := &tools.Locations{}
				helpers.Fetch_By_Id(locId, Loc)
				wu.Lock()
				for _, Location := range Loc.Locations {
					if InputSearch == Location {
						foundId = append(foundId, v.Id)
					}
				}
				wu.Unlock()
			}(v.Locations)

		}
	}
	wg.Wait()
	// ! end

	var sliceArt []tools.Artists
	for _, v := range foundId {
		var artists tools.Artists

		url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/artists/%d", v)
		helpers.Fetch_By_Id(url, &artists)
		sliceArt = append(sliceArt, artists)
	}

	lresult := Result{
		Artist:        sliceArt,
		SearchElement: SearchArtist,
	}

	helpers.RenderTemplates(w, "index.html", lresult, 200)
}
