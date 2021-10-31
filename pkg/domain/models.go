package domain

import (
	"encoding/json"
	"time"
)

type Book struct {
	tableName   struct{}    `pg:"books"`
	Id          int64       `json:"id" pg:"id"`
	Isbn        json.Number `json:"isbn," pg:"isbn"`
	Name        string      `json:"name" pg:"name"`
	Author      int64       `json:"author" pg:"author"`
	Year        int64       `json:"year" pg:"year"`
	CreatedDate time.Time   `json:"createddate" pg:"createddate"`
	CreatedBy   string      `json:"createdby" pg:"createdby"`
	UpdatedDate time.Time   `json:"updateddate" pg:"updateddate"`
	UpdatedBy   string      `json:"updatedby" pg:"updatedby"`
	IsActive    bool        `json:"isactive" pg:"isactive"`
}

type Author struct {
	tableName   struct{}  `pg:"author"`
	Id          int64     `json:"id" pg:"id"`
	Name        string    `json:"name" pg:"name"`
	Surname     string    `json:"surname" pg:"surname"`
	CreatedDate time.Time `json:"createddate" pg:"createddate"`
	CreatedBy   string    `json:"createdby" pg:"createdby"`
	UpdatedDate time.Time `json:"updateddate" pg:"updateddate"`
	UpdatedBy   string    `json:"updatedby" pg:"updatedby"`
	IsActive    bool      `json:"isactive" pg:"isactive"`
}

type Members struct {
	tableName   struct{}  `pg:"members"`
	Id          int64     `json:"id" pg:"id"`
	Name        string    `json:"name" pg:"name"`
	Surname     string    `json:"surname" pg:"surname"`
	DateofBrith time.Time `json:"dateofbrith" pg:"dateofbrith"`
	IsPremium   bool      `json:"ispremium" pg:"ispremium"`
	CreatedDate time.Time `json:"createddate" pg:"createddate"`
	CreatedBy   string    `json:"createdby" pg:"createdby"`
	UpdatedDate time.Time `json:"updateddate" pg:"updateddate"`
	UpdatedBy   string    `json:"updatedby" pg:"updatedby"`
	IsActive    bool      `json:"isactive" pg:"isactive"`
}

type History struct {
	tableName   struct{}  `pg:"history"`
	Id          int64     `json:"id" pg:"id"`
	HistoryType int       `json:"historytype" pg:"historytype"`
	CreatedDate time.Time `json:"createddate" pg:"createddate"`
	CreatedBy   string    `json:"createdby" pg:"createdby"`
	UpdatedDate time.Time `json:"updateddate" pg:"updateddate"`
	UpdatedBy   string    `json:"updatedby" pg:"updatedby"`
	IsActive    bool      `json:"isactive" pg:"isactive"`
}
