package database

import "gorm.io/gorm"

type Secret struct {
	gorm.Model `json:"-"`
	ID         uint   `json:"id" gorm:"primaryKey,autoIncrement"`
	Content    string `json:"content"`
	RemovalKey string `json:"-"`
	UsagesLeft *int   `json:"usagesLeft"`
}
