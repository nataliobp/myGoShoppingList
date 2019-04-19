package identityAccess

type SignUpService struct {
	userRepository UserRepository
}

func NewSignUpService(repository UserRepository) *SignUpService {
	return &SignUpService{repository}
}

func (s *SignUpService) SignUp(email, password string) {
	user := User{
		Email:    email,
		Password: password,
	}

	s.userRepository.save(user)
}
