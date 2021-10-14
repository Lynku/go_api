package models

import (
	"gorm.io/gorm"
)

//Global settings
type Settings struct {
	gorm.Model
	Name  string
	Value string
}

type Users struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Status   int
}

type Tokens struct {
	gorm.Model
	UserId uint
	Token  string
}

type PageType struct {
	ID   uint `gorm:"primarykey"`
	Name string
}

type PageTypeOptions struct {
	ID         uint `gorm:"primarykey"`
	PageTypeId uint
	Type       string
	Name       string
}

type Page struct {
	gorm.Model
	PageTypeId uint
	Author     uint
	Title      string
	Content    string
	Status     int
}

//ACF
type PageCf struct {
	ID                uint64 `gorm:"primarykey"`
	PageId            uint
	PageTypeOptionsId uint
	Value             string
}

type Media struct {
	gorm.Model
	Path string
}
