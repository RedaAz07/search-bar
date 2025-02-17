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
	var wu sync.Mutex
	if r.Method != http.MethodGet {
		// execute the not found  template
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorMethodnotAll, http.StatusMethodNotAllowed)
		return
	}
	var spliteInput []string
	spliteInput = strings.Split(r.FormValue("search"), "#")
	var InputSearch string
	var typrSearch string

	if len(spliteInput) >= 2 {

		InputSearch = (strings.ToLower(spliteInput[0]))
		typrSearch = (strings.ToLower(spliteInput[1]))
	} else {
		InputSearch = (strings.ToLower(spliteInput[0]))
	}
	var foundId []int
	for _, v := range Artists {
		switch {
		case len(spliteInput) >= 2 && typrSearch == "artist/band name":
			if strings.ToLower(v.Name) == InputSearch {
				foundId = append(foundId, v.Id)
			}
		case len(spliteInput) >= 2 && typrSearch == "first album date":

			if strings.ToLower(v.FirstAlbum) == InputSearch {
				foundId = append(foundId, v.Id)
			}
		case len(spliteInput) >= 2 && typrSearch == "creation date":

			if strconv.Itoa(v.CreationDate) == InputSearch {
				foundId = append(foundId, v.Id)
			}

		case len(spliteInput) >= 2 && typrSearch == "members":
    for _, members := range v.Members {
        if strings.ToLower(members) == InputSearch && strings.ToLower(v.Name) != InputSearch {
            foundId = append(foundId, v.Id)
        }
    }
		//	fmt.Println(foundId)
	case len(spliteInput) >= 2 && typrSearch == "locations":
		var wg sync.WaitGroup
		var mu sync.Mutex // قفل (Mutex) لحماية التعديلات على `foundId`
	
		for _, v := range Artists {
			wg.Add(1) // زيادة العداد لكل Goroutine
			go func(locId string, artistId int) {
				defer wg.Done() // عند انتهاء Goroutine، ننقص العداد
				
				Loc := &tools.Locations{}
				helpers.Fetch_By_Id(locId, Loc)
				
				for _, Location := range Loc.Locations {
					if InputSearch == strings.ToLower(Location) {
						mu.Lock() // قفل لحماية `foundId`
						foundId = append(foundId, artistId)
						mu.Unlock()
						return // خروج من Goroutine بمجرد العثور على نتيجة
					}
				}
			}(v.Locations, v.Id) // تمرير `v.Locations` و `v.Id` لتجنب مشاكل التكرار
		}
	
		wg.Wait() 
			// ! end
		default:

			for _, members := range v.Members {
				if strings.HasPrefix(strings.ToLower(members), strings.ToLower(InputSearch)) {
					foundId = append(foundId, v.Id)
				}
			}

			if strings.HasPrefix(strings.ToLower(v.Name), strings.ToLower(InputSearch)) || strings.HasPrefix(strings.ToLower(v.FirstAlbum), strings.ToLower(InputSearch)) || strings.HasPrefix(strings.ToLower(strconv.Itoa(v.CreationDate)), strings.ToLower(InputSearch)) {
				foundId = append(foundId, v.Id)
			}

			// !  i use  goroutin , mutex
			go func(locId string) {
				Loc := &tools.Locations{}
				helpers.Fetch_By_Id(locId, Loc)
				wu.Lock()
				for _, Location := range Loc.Locations {
					if strings.HasPrefix(strings.ToLower(Location), strings.ToLower(InputSearch)) {
						foundId = append(foundId, v.Id)
					}
				}
				wu.Unlock()
			}(v.Locations)

			// ! end

		}
	}
	//  !  to ensure that the result doasent repated in my map
	var sliceArt []tools.Artists
	var LastSliceId []int
	m := make(map[int]bool)
	for _, n := range foundId {
		if !m[n] {
			m[n] = true
			LastSliceId = append(LastSliceId, n)
		}
	}

	// ! end
	go func() {
		for _, v := range LastSliceId {

			var artists tools.Artists
			url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/artists/%d", v)
			helpers.Fetch_By_Id(url, &artists)
			sliceArt = append(sliceArt, artists)
		}
		}()

	var lresult Result
	if len(sliceArt) == 0 {
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorBadReq, 404)
		return

	} else {
		lresult = Result{
			Artist:        sliceArt,
			SearchElement: SearchArtist,
		}
	}
// error in search by location l
// long time when i search about london-uk


 


	helpers.RenderTemplates(w, "index.html", lresult, 200)
}
