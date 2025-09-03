package entities

type Contact struct {
	ID        int     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	FirstName string  `json:"first_name" gorm:"column:first_name;size:100;not null"`
	LastName  *string `json:"last_name,omitempty" gorm:"column:last_name;size:100"`
	Email     *string `json:"email,omitempty" gorm:"column:email;size:100"`
	Phone     *string `json:"phone,omitempty" gorm:"column:phone;size:20"`
	Username  string  `json:"username" gorm:"column:username;size:255;not null"`
}

func (Contact) TableName() string {
	return "contacts"
}
