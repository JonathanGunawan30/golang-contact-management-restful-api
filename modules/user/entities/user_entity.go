package entities

type User struct {
	Username string `json:"username" gorm:"column:username;primaryKey;size:255;not null"`
	Password string `json:"-" gorm:"column:password;size:255;not null"`
	Name     string `json:"name" gorm:"column:name;size:255;not null"`
	Token    string `json:"token" gorm:"column:token;size:255"`
}

func (User) TableName() string {
	return "users"
}
