package models

type Drink struct {
	ID      int    `json:"id"`
	Drink   string `json:"drink"`
	LevelId int    `json:"level_id"`
}
