package entity

import "time"

type Book struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Price       int32     `json:"price"`
	PublishDate time.Time `json:"publish_date"`
}
