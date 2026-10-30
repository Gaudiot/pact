package user

import (
	"context"

	"github.com/pact/core/database/repository"
	student_repository "github.com/pact/core/database/repository/student"
)

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserService struct {
	studentRepository student_repository.Repository
}

func NewUserService(studentRepository student_repository.Repository) *UserService {
	return &UserService{
		studentRepository: studentRepository,
	}
}

func (us *UserService) GetUser(ctx context.Context, id string) (*User, error) {
	user, err := us.studentRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &User{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

type ListUserParams struct {
	Name  string
	Page  int
	Limit int
}

func (us *UserService) ListUsers(ctx context.Context, params ListUserParams) ([]User, error) {
	repoUsers, err := us.studentRepository.GetUsers(
		ctx,
		&student_repository.GetUsersParams{Name: params.Name},
		&repository.Pagination{Page: params.Page, Limit: params.Limit},
	)
	if err != nil {
		return nil, err
	}

	users := []User{}
	for _, user := range repoUsers {
		users = append(users, User{Id: user.ID, Name: user.Name, Email: user.Email})
	}

	return users, nil
}

func (us *UserService) CreateUser(ctx context.Context, name string, email string) (*User, error) {
	user, err := us.studentRepository.Create(ctx, &student_repository.CreateStudentParams{
		Name:  name,
		Email: email,
	})
	if err != nil {
		return nil, err
	}

	return &User{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
