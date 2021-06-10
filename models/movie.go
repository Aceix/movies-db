package models

import "github.com/kamva/mgm/v3"

type Movie struct {
	mgm.DefaultModel  `bson:",inline"`
	Title  string  `json:"title" bson:"title"`
	Year   int32  `json:"year" bson:"year"`
	Genres []string  `json:"genres" bson:"genres"`
	ImgUrl string  `json:"imgUrl" bson:"imgUrl"`
}

func CreateMovie() *Movie {
	return &Movie{
		Title:  "",
		Year:   0,
		Genres: []string{},
		ImgUrl: "",
	}
}
