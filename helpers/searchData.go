package helpers

import (
	"fmt"
	"groupie/tools"
	"strconv"
	"sync"
)

func SearchData(allArtists *[]tools.Artists, wg *sync.WaitGroup, data *tools.Data) {

	defer wg.Done()
	err := Fetch("https://groupietrackers.herokuapp.com/api/locations", &Locations)
	if err != nil {
		fmt.Println("Error fetching API:", err)
		return
	}

	locationsSet := make(map[string]bool)

	for _, entry := range Locations.Index {
		for _, loc := range entry.Locations {
			locationsSet[loc] = true
		}
	}



	var SearchArtist = make(map[string]string)

	for _, structs := range *allArtists {
		SearchArtist[structs.Name] = "artist/band name"
		SearchArtist[structs.FirstAlbum] = "first album date"
		SearchArtist[strconv.Itoa(structs.CreationDate)] = "creation date"

		for _, v := range structs.Members {
			SearchArtist[v] = "members"
		}
		// locations

		for v := range locationsSet {
			SearchArtist[v] = "locations"
		}
	}

	data.SearchElement = SearchArtist
}
