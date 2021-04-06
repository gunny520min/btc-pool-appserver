package model

type Banner struct {
	Id     string `json:"id"`
	Index  int    `json:"index"`
	ImgUrl string `json:"imgUrl"`
	Link   string `json:"link"`
}
