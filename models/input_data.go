package models

type DataInput struct {
	PageTypeId uint   `json:"type" binding:"required"`
	Author     uint   `json:"author"`
	Title      string `json:"title" binding:"required"`
	Content    string `json:"content"`
	Status     int    `json:"status"`
}
