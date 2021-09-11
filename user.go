package drive

type User struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type UserService interface {
	User(id int64) (*User, error)
	Users() ([]*User, error)
	CreateUser(u *User) error
	DeleteUser(id int64) error
}
