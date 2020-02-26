package models

type Movie struct {
	OriginalName string  `json:"original_name"`
	GenreIDs     []int64 `json:"genre_ids"`
	Name         string  `json:"name"`
	Popularity   float64 `json:"popularity"`
}
