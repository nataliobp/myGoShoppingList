package identityAccess

type UserRepository interface {
	save(user User) int64
	findAll() []User
	findByEmail(email string) *User
	findById(id int64) *User
}
