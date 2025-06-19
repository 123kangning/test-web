package entity

import "time"

type User struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Password   string    `json:"password"`
	Status     int32     `json:"status"`
	CreateTime time.Time `json:"create_time"`
}
