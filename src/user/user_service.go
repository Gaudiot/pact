package user

import (
	"context"

	student_repository "github.com/pact/core/database/repository/student"
)

type UserService struct {
	studentRepository student_repository.Repository
}

func NewUserService(studentRepository student_repository.Repository) *UserService {
	return &UserService{
		studentRepository: studentRepository,
	}
}

func (us *UserService) CreateUser(name string, email string) (*struct{ Id, Name, Email string }, error) {
	ctx := context.Background()

	student, err := us.studentRepository.Create(ctx, &student_repository.CreateStudentParams{
		Name:  name,
		Email: email,
	})
	if err != nil {
		return nil, err
	}

	return &struct {
		Id    string
		Name  string
		Email string
	}{
		Id:    student.ID,
		Name:  student.Name,
		Email: student.Email,
	}, nil
}

func (us *UserService) GetUser(id string) (*struct{ Id, Name string }, error) {
	ctx := context.Background()

	user, err := us.studentRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &struct {
		Id   string
		Name string
	}{
		Id:   user.ID,
		Name: user.Name,
	}, nil
}

func (us *UserService) GetByName(name string) ([]struct{ Id, Name string }, error) {
	return nil, nil
}
