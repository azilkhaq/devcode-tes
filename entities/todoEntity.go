package entities

type Todo struct {
	Id              int    `gorm:"primary_key;auto_increment;" json:"id"`
	ActivityGroupId string `json:"activity_group_id"`
	Title           string `json:"title"`
	IsActive        string `json:"is_active"`
	Priority        string `json:"priority"`
	Base
}
