package userService

import "github.com/google/uuid"

type UserService interface {
	CreateUser(email, password string) (User, error)
	GetAllUsers() ([]User, error)
	GetUserByID(id string) (User, error)
	UpdateUser(id, email, password string) (User, error)
	DeleteUser(id string) error
}

type userService struct {
	repo UserRepository
}

func NewUserService(r UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) CreateUser(email, password string) (User, error) {
	user := User{
		ID:       uuid.NewString(),
		Email:    email,
		Password: password,
	}

	if err := s.repo.CreateUser(&user); err != nil {
		return User{}, err
	}

	return user, nil
}

func (s *userService) GetAllUsers() ([]User, error) {
	return s.repo.GetAllUsers()
}

func (s *userService) GetUserByID(id string) (User, error) {
	return s.repo.GetUserByID(id)
}

func (s *userService) UpdateUser(id, email, password string) (User, error) {
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return User{}, err
	}

	user.Email = email
	user.Password = password

	if err := s.repo.UpdateUser(user); err != nil {
		return User{}, err
	}

	return user, nil
}

func (s *userService) DeleteUser(id string) error {
	return s.repo.DeleteUser(id)
}
