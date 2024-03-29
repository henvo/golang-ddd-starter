package postgres

import "github.com/henvo/golang-ddd-starter/internal/domain"

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

// Create a user in the database
func (s *UserService) CreateUser(du *domain.User) error {
	return nil
}

// Delete a user from the database
func (s *UserService) DeleteUser(du *domain.User) error {
	return nil
}

// Find a user from the database
func (s *UserService) FindUser(id string) (*domain.User, error) {
	return &domain.User{}, nil
}

func (s *UserService) FindUsers(ids []string) ([]*domain.User, error) {
	return []*domain.User{}, nil
}

func (s *UserService) UpdateUser(du *domain.User) error {
	return nil
}
