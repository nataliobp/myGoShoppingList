package identityAccess

type FindUserService struct {
	userRepository UserRepository
}

func NewFindUserService(repository UserRepository) *FindUserService {
	return &FindUserService{repository}
}

func (s *FindUserService) FindAll() []User {
	return s.userRepository.findAll()
}

func (s *FindUserService) FindByEmail(email string) *User {
	return s.userRepository.findByEmail(email)
}
