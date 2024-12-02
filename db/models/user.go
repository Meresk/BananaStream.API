package models

type User struct {
	ID       uint   `gorm:"primary_key;auto_increment"`
	Login    string `gorm:"type:nvarchar(255);unique;not null"`
	Password string `gorm:"size:255;not null"`
}
