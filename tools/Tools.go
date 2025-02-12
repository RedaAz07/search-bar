package tools

import (
	"net/http"
	"text/template"
)

var Tp *template.Template

type ErrorPage struct {
	Code         int
	ErrorMessage string
}

var ErrorBadReq = ErrorPage{
	Code:         http.StatusBadRequest,
	ErrorMessage: "Oops! It looks like there was an issue with your request. Please check your input and try again.",
}
var ErrorNotFound = ErrorPage{
	Code:         http.StatusNotFound,
	ErrorMessage: "Uh-oh! The page you're looking for doesn't exist. It might have been moved or deleted.",
}
var ErrorMethodnotAll = ErrorPage{
	Code:         http.StatusMethodNotAllowed,
	ErrorMessage: "The request method is not supported for this resource. Please check and try again with a valid method.",
}
var ErrorInternalServerErr = ErrorPage{
	Code:         http.StatusInternalServerError,
	ErrorMessage: "Something went wrong on our end. We're working on fixing it—please try again later!",
}

type Artists struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Locations struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type ConcertDates struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Relations struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}
