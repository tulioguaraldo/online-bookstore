package model

import "time"

type Book struct {
	Id          uint64    `gorm:"primaryKey;autoIncrement;column:Id;"`
	Name        string    `gorm:"column:BookName;"`
	Author      string    `gorm:"column:BookAuthor;"`
	Publisher   string    `gorm:"column:BookPublisher;"`
	Price       float64   `gorm:"column:BookPrice;"`
	WroteAt     time.Time `gorm:"column:WroteAt;"`
	PublishedAt time.Time `gorm:"column:PublishedAt;"`
}
