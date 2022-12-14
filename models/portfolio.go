package models

type Portfolio struct {
	ID          string
	Url         string
	CreaterUser string
	Name        string
	Text        *[]Text
	Photo       *[]Photo
	Colors      *Colors
	Struct      *[][]Block
}
