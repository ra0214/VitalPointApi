package domain

type IUser interface {
	SaveUser(userName string, email string, password string) error
	DeleteUser(id int32) error
	UpdateUser(id int32, userName string, email string, password string) error
	GetAll() ([]User, error)
	GetUserByCredentials(userName string) (*User, error)
}

type User struct {
	ID       int32  `json:"id"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

func NewUser(userName string, email string, password string) *User {
	return &User{
		UserName: userName,
		Email:    email,
		Password: password,
	}
}

func (u *User) SetUserName(userName string) {
	u.UserName = userName
}
