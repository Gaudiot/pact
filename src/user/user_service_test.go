package user

import (
	"errors"
	"testing"

	student_repository "github.com/pact/core/database/repository/student"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func setupTest(t *testing.T) (
	*UserService,
	*student_repository.MockRepository,
) {
	t.Helper()

	mockStudentRepository := new(student_repository.MockRepository)

	sut := NewUserService(mockStudentRepository)

	return sut, mockStudentRepository
}

func TestCreateUser(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		sut, mockStudentRepository := setupTest(t)
		expected := &student_repository.Student{
			ID:    "id",
			Name:  "john doe",
			Email: "john@doe.com",
		}

		mockStudentRepository.On("Create", mock.Anything, mock.Anything).Return(expected, nil).Once()

		result, err := sut.CreateUser(expected.Name, expected.Email)

		require.ErrorIs(t, err, nil)
		require.Equal(t, result.Name, expected.Name)
		require.Equal(t, result.Email, expected.Email)
		mockStudentRepository.AssertExpectations(t)
	})
	t.Run("User Repository Error", func(t *testing.T) {
		tests := []struct {
			testName      string
			expectedError error
		}{
			{
				"Email Already Registered",
				student_repository.ErrEmailAlreadyRegistered,
			},
			{
				"Repository Internal Error",
				errors.New("internal_error"),
			},
		}

		for _, tt := range tests {
			t.Run(tt.testName, func(t *testing.T) {
				sut, mockStudentRepository := setupTest(t)
				expected := &student_repository.CreateStudentParams{
					Name:  "john doe",
					Email: "john@doe.com",
				}
				mockStudentRepository.On("Create", mock.Anything, expected).Return(nil, tt.expectedError).Once()

				user, err := sut.CreateUser(expected.Name, expected.Email)

				require.ErrorIs(t, err, tt.expectedError)
				require.Nil(t, user)

				mockStudentRepository.AssertExpectations(t)
			})
		}
	})
}

func TestGetUser(t *testing.T) {
	t.Run("Success", func(t *testing.T) {

	})
	t.Run("User Not Found", func(t *testing.T) {

	})
	t.Run("User Repository Error", func(t *testing.T) {

	})
}

func TestGetByName(t *testing.T) {
	t.Run("Success", func(t *testing.T) {

	})
	t.Run("User Repository Internal Error", func(t *testing.T) {

	})
}
