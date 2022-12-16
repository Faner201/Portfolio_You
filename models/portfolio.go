package models

type Portfolio struct {
	ID          string     `json:"id"`
	Url         string     `json:"url"`
	CreaterUser string     `json:"createrUser"`
	Name        string     `json:"name"`
	Text        *[]Text    `json:"texts"`
	Photo       *[]Photo   `json:"images"`
	Colors      *Colors    `json:"colors"`
	Struct      *[][]Block `json:"structure"`
}
