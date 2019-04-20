package identityAccess

type InMemoryUserRepository struct{ memory []User }

func NewInMemoryUserRepository() InMemoryUserRepository {
	return InMemoryUserRepository{[]User{}}
}

func (r *InMemoryUserRepository) save(user User) int64 {
	r.memory = append(r.memory, user)
	return int64(len(r.memory))
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

func (r *InMemoryUserRepository) findById(id int64) *User {
	for _, user := range r.memory {
		if user.Id == id {
			return &user
		}
	}

	return nil
}
