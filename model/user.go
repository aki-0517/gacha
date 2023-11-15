package model

type User struct{
	ID string
	Name string
	Token string `gorm:"primary_key"`
}