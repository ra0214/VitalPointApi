package domain

type IUser interface {
	SaveUser(userName string, email string, password string, role string) error
	DeleteUser(id int32) error
	UpdateUser(id int32, userName string, email string, password string, role string) error
	GetAll() ([]User, error)
	GetUserByCredentials(userName string) (*User, error)
}

type User struct {
	ID       int32  `json:"id"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
	Role     string `json:"role"` // Nuevo campo
}

func NewUser(userName string, email string, password string, role string) *User {
	return &User{
		UserName: userName,
		Email:    email,
		Password: password,
		Role:     role,
	}
}

func (u *User) SetUserName(userName string) {
	u.UserName = userName
}
