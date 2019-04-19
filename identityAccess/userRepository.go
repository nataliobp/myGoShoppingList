package identityAccess

type UserRepository interface {
	save(user User)
	findAll() []User
	findByEmail(email string) *User
}
