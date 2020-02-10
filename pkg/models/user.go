package models

type User struct {
	Model
	Name string `gorm:"column:name;unique;not null" json:"name"`
}

// Set User's table name to be `users`
func (User) TableName() string {
	return "users"
}
