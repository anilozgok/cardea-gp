package entity

type Recipe struct {
	ID           string   `json:"id"`
	Ingredients  []string `json:"ingredients"`
	PictureLink  string   `json:"picture_link"`
	Instructions string   `json:"instructions"`
	Title        string   `json:"title"`
}
