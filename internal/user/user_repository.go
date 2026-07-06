package user

type userRepository interface {
	GetUser(id string) (*User, error)
	CreateUser(user *User) error
	UpdateUser(user *User) error
	DeleteUser(id string) error
}
