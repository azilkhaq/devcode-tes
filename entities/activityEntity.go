package entities

type Activity struct {
	Id    int    `gorm:"primary_key;auto_increment;" json:"id"`
	Title string `json:"title"`
	Email string `json:"email"`
	Base
}
