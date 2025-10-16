package models

type Tell struct {
	ID      int    `json:"id"`
	Tell    string `json:"tell"`
	LevelId int    `json:"level_id"`
}
