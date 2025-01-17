package model

type Todo struct {
	ID   uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Task string `json:"task"`
	Done bool   `json:"done" gorm:"default:false"`
}
