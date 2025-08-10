package database

import (
	"database/sql"
	// "time"
)

type EventModel struct {
	DB *sql.DB
}

type Event struct {
	Id int `json:"id"`
	OwnerId int `json:"ownerId" binding:"required"`
	Name string `json:"name" binding:"required, min=3"`
	Description string `json:"description" binding:"required, min=10"`
	Date string `json:"date" binding:"required, datetime=2020-08-10"`
	Location string `json:"" binding:"required, min=3"`
}