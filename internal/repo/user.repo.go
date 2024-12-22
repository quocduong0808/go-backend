package repo

import "strings"

type IUserRepository interface {
	EmailExist(email string) bool
}

type userRepository struct {
}

// EmailExist implements IUserRepository.
func (us *userRepository) EmailExist(email string) bool {
	return strings.HasSuffix(email, "registed")
}

func NewUserRepository() IUserRepository {
	return &userRepository{}
}
