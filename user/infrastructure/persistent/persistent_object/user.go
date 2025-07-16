package persistent_object

type User struct {
	Base
	email    string `gorm:"size:255;not null"`
	password string `gorm:"size:255;not null"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) ID() int64 {
	return u.Base.id
}

func (u *User) Email() string {
	return u.email
}

func (u *User) Password() string {
	return u.password
}
