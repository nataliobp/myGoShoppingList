package identityAccess

type UserRepository interface {
	save(user User) int64
	findAll() []User
	findByEmail(email string) (*User, error)
	findById(id int64) (*User, error)
}
