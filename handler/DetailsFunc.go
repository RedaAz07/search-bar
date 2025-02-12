package handler

import (
	"net/http"
	"strconv"

	"groupie/helpers"
	"groupie/tools"
)

func DetailsFunc(w http.ResponseWriter, r *http.Request) {


//check the method 
	if r.Method != http.MethodGet {

		// execute the not found  template
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorMethodnotAll, http.StatusMethodNotAllowed)
		return
	}


	type fetchingData struct {
		Artist    *tools.Artists
		Locations *tools.Locations
		Dates     *tools.ConcertDates
		Relations *tools.Relations
	}
// get the id from url 
	id := r.URL.Query().Get("id")
	// to int 
	Id, err := strconv.Atoi(id)
	if err != nil {

	
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorBadReq, http.StatusBadRequest)
		return
	}

	var artistFound *tools.Artists
// get the user 
for i := range Artists {
    if Id == Artists[i].Id {
        artistFound = &Artists[i]
        break
    }
}
//  to see if the user exists 
	if artistFound == nil {

	

		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorNotFound, http.StatusNotFound)
		return
	}

	var locations tools.Locations
	var dates tools.ConcertDates
	var relations tools.Relations
// fetch the location and get the result  in the location variavle 
	errr := helpers.Fetch_By_Id(artistFound.Locations, &locations)
	if errr != nil {

		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorInternalServerErr, 500)
		return
	}
// fetch the dates and get the result  in the dates variavle 

	errr = helpers.Fetch_By_Id(artistFound.ConcertDates, &dates)
	if errr != nil {
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorInternalServerErr, 500)
		return
	}
// fetch the relations and get the result  in the relations variavle 

	errr = helpers.Fetch_By_Id(artistFound.Relations, &relations)
	if errr != nil {
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorInternalServerErr, 500)
		return
	}
// set all the that that we found into the fetching variable 
	fetching := fetchingData{
		Artist:    artistFound,
		Locations: &locations,
		Dates:     &dates,
		Relations: &relations,
	}

	helpers.RenderTemplates(w, "details.html", fetching, 200)
}


