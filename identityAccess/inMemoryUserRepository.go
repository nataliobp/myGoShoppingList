package identityAccess

import "fmt"

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

func (r *InMemoryUserRepository) findByEmail(email string) (*User, error) {
	for _, user := range r.memory {
		if user.Email == email {
			return &user, nil
		}
	}

	return nil, fmt.Errorf("not found")
}

func (r *InMemoryUserRepository) findById(id int64) (*User, error) {
	for _, user := range r.memory {
		if user.Id == id {
			return &user, nil
		}
	}

	return nil, fmt.Errorf("not found")
}
