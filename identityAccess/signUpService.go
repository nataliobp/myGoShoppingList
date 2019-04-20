package identityAccess

type SignUpService struct {
	userRepository UserRepository
}

func NewSignUpService(repository UserRepository) *SignUpService {
	return &SignUpService{repository}
}

func (s *SignUpService) SignUp(email, password string) int64 {
	user := User{
		Email:    email,
		Password: password,
	}

	return s.userRepository.save(user)
}
