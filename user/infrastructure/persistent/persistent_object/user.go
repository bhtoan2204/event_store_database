package persistent_object

type User struct {
	Base
	Email    string `gorm:"index;unique;size:255;not null"`
	Password string `gorm:"size:255;not null"`
}

func NewUser(email string, password string) *User {
	return &User{
		Email:    email,
		Password: password,
	}
}

func (u *User) TableName() string {
	return "users"
}
