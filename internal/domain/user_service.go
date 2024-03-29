package domain

type UserService interface {
	CreateUser(u *User) error
	FindUser(id string) (*User, error)
	FindUsers(ids []string) ([]*User, error)
	DeleteUser(*User) error
	UpdateUser(*User) error
}
