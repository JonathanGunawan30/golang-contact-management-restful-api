package entities

type Address struct {
	ID         int     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Street     *string `json:"street,omitempty" gorm:"column:street;size:255"`
	City       *string `json:"city,omitempty" gorm:"column:city;size:100"`
	Province   *string `json:"province,omitempty" gorm:"column:province;size:100"`
	Country    string  `json:"country" gorm:"column:country;size:100;not null"`
	PostalCode string  `json:"postal_code" gorm:"column:postal_code;size:10;not null"`
	ContactID  int     `json:"contact_id" gorm:"column:contact_id;not null"`
}

func (Address) TableName() string {
	return "addresses"
}
