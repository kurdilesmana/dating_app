package service

import (
	"github.com/kurdilesmana/dating_app/internal/domain"
	"github.com/kurdilesmana/dating_app/internal/repository"

	"github.com/pkg/errors"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (s *UserService) RegisterUser(user *domain.User) error {
	// Implement user registration logic using UserRepository
	// For example:
	// Check if the username or email is already in use
	// Hash the user's password before storing it in the database
	// Save the user to the database using UserRepository

	// For simplicity, let's assume UserRepository has a Create method
	if err := s.UserRepository.Create(user); err != nil {
		return errors.Wrap(err, "failed to register user")
	}

	return nil
}

func (s *UserService) AuthenticateUser(username, password string) (bool, error) {
	// Implement user authentication logic using UserRepository
	// For example:
	// Check if the user with the given username exists
	// Verify the password against the stored hash in the database

	// For simplicity, let's assume UserRepository has a FindByUsername method
	_, err := s.UserRepository.FindByUsername(username)
	if err != nil {
		return false, errors.Wrap(err, "failed to authenticate user")
	}

	// Compare the hashed password (you'll need to use a password hashing library)
	// For simplicity, let's assume there's a hypothetical PasswordMatches function
	//if !PasswordMatches(password, user.Password) {
	//    return false, nil
	//}

	// Password matches
	return true, nil
}
