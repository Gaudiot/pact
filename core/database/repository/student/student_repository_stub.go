package student_repository

import (
	"context"
	"strings"
	"uuid"

	"github.com/pact/core/database/repository"
)

var _ Repository = &stub{}

type stub struct {
	students map[string]*Student
}

func NewStub() *stub {
	return &stub{
		students: map[string]*Student{
			"1": {
				ID:    "1",
				Name:  "Victor",
				Email: "victor@gaudiot.com",
			},
			"2": {
				ID:    "2",
				Name:  "Mylena",
				Email: "mylena@oliveira.com.br",
			},
			"3": {
				ID:    "3",
				Name:  "John Doe",
				Email: "john@doe.com",
			},
			"4": {
				ID:    "4",
				Name:  "Victor",
				Email: "victor@fake.com",
			},
		},
	}
}

func (s *stub) Create(ctx context.Context, params *CreateStudentParams) (*Student, error) {
	newId := uuid.NewV4().String()
	student := &Student{
		ID:    newId,
		Name:  params.Name,
		Email: params.Email,
	}
	s.students[newId] = student
	return student, nil
}

func (s *stub) GetByID(ctx context.Context, id string) (*Student, error) {
	student, ok := s.students[id]
	if !ok {
		return nil, ErrStudentNotFound
	}

	return student, nil
}

func (s *stub) GetUsers(ctx context.Context, params *GetUsersParams, pagination *repository.Pagination) ([]*Student, error) {
	filteredStudents := []*Student{}

	for _, student := range s.students {
		if params.Name != "" && !strings.HasPrefix(strings.ToLower(student.Name), strings.ToLower(params.Name)) {
			continue
		}

		filteredStudents = append(filteredStudents, student)
	}

	start := pagination.Page * pagination.Limit
	if start >= len(filteredStudents) {
		return nil, nil
	}
	end := min(start+pagination.Limit, len(filteredStudents))
	students := filteredStudents[start:end]

	return students, nil
}
