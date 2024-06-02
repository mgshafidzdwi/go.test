package entity

import "time"

type Book struct {
	ID            uint64    `gorm:"primary_key:auto_increment" json:"id"`
	Title         string    `gorm:"type:varchar(255);not null" json:"title"`
	Author        string    `gorm:"type:varchar(255);not null" json:"author"`
	ISBN          string    `gorm:"type:varchar(13);unique;not null" json:"isbn"`
	PublishedDate time.Time `gorm:"type:date;not null" json:"published_date"`
}
