package persistent_object

type User struct {
	Base
	Code     string `gorm:"index;unique;size:255;not null"`
	Email    string `gorm:"index;unique;size:255;not null"`
	Password string `gorm:"size:255;not null"`
}

func NewUser(code string, email string, password string) *User {
	return &User{
		Code:     code,
		Email:    email,
		Password: password,
	}
}

func (u *User) TableName() string {
	return "users"
}
