package domain

type User struct {
	ID       uint
	Email    string
	Password string
}

type UserRepository interface {
	FindByEmail(email string) (User, error)
	Create(user User) error
}
