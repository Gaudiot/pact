package student_repository

import (
	"context"
	"strings"
	"uuid"
)

var _ Repository = &stub{}

type stub struct {
	students map[string]*Student
}

func NewStub() *stub {
	return &stub{
		students: map[string]*Student{},
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

func (s *stub) GetByName(ctx context.Context, name string) ([]*Student, error) {
	students := []*Student{}

	lowerName := strings.ToLower(name)
	for _, student := range s.students {
		lowerStudentName := strings.ToLower(student.Name)
		if lowerName == lowerStudentName {
			students = append(students, student)
		}
	}

	return students, nil
}
