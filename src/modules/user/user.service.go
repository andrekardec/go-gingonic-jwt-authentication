package user

import (
	"context"
	"errors"
	"fmt"
)

type ServiceInterface interface {
	create(ctx context.Context, request CreateUserRequestDto) error
	exists(email string) error
}

type Service struct {
	userRepository RepositoryInterface
}

func NewUserService(userRepository RepositoryInterface) ServiceInterface {
	return &Service{
		userRepository,
	}
}

func (s *Service) exists(email string) error {
	user, _ := s.userRepository.getByEmail(email)

	if user.Id != "" {
		return errors.New("this email has already been taken")
	}

	return nil
}

func (s *Service) create(ctx context.Context, dto CreateUserRequestDto) error {

	err := s.exists(dto.Email)
	if err != nil {
		return err
	}

	if dto.Password != dto.PasswordConfirmation {
		return errors.New(fmt.Sprintf("password do not match"))
	}

	createUser := &CreateUserDto{
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Email:     dto.Email,
		Password:  dto.Password,
	}

	err = s.userRepository.create(ctx, *createUser)
	if err != nil {
		return errors.New("internal server error")
	}

	return nil
}
