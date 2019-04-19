package identityAccess

type InMemoryUserRepository struct{ memory []User }

func NewInMemoryUserRepository() InMemoryUserRepository {
	return InMemoryUserRepository{[]User{}}
}

func (r *InMemoryUserRepository) save(user User) {
	r.memory = append(r.memory, user)
}

func (r *InMemoryUserRepository) findAll() []User {
	return r.memory
}

func (r *InMemoryUserRepository) findByEmail(email string) *User {
	for _, user := range r.memory {
		if user.Email == email {
			return &user
		}
	}

	return nil
}
