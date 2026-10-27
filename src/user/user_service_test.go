package user

import (
	"errors"
	"testing"

	student_repository "github.com/pact/core/database/repository/student"
	"github.com/stretchr/testify/assert"
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

		require.NoError(t, err)
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
		sut, mockStudentRepository := setupTest(t)
		expected := &student_repository.Student{
			ID:    "id",
			Name:  "john doe",
			Email: "john@doe.com",
		}

		mockStudentRepository.On("GetByID", mock.Anything, mock.Anything).Return(expected, nil).Once()

		user, err := sut.GetUser("id")

		mockStudentRepository.AssertExpectations(t)

		assert.NoError(t, err)
		require.NotNil(t, user)
		assert.Equal(t, user.Id, expected.ID)
		assert.Equal(t, user.Name, expected.Name)
	})
	t.Run("User Repository Error", func(t *testing.T) {
		tests := []struct {
			testName      string
			expectedError error
		}{
			{
				"User Not Found",
				student_repository.ErrStudentNotFound,
			},
			{
				"Repository Internal Error",
				errors.New("internal_error"),
			},
		}

		for _, tt := range tests {
			t.Run(tt.testName, func(t *testing.T) {
				sut, mockStudentRepository := setupTest(t)
				mockStudentRepository.On("GetByID", mock.Anything, mock.Anything).Return(nil, tt.expectedError).Once()

				user, err := sut.GetUser("id")

				mockStudentRepository.AssertExpectations(t)

				assert.ErrorIs(t, err, tt.expectedError)
				require.Nil(t, user)
			})
		}
	})
}

func TestGetByName(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		sut, mockStudentRepository := setupTest(t)
		expected := []*student_repository.Student{
			{
				ID:    "id",
				Name:  "john doe",
				Email: "john@doe.com",
			},
			{
				ID:    "id2",
				Name:  "jane doe",
				Email: "jane@doe.com",
			},
			{
				ID:    "id3",
				Name:  "lorem ipsum",
				Email: "lorem@ipsum.com",
			},
		}

		mockStudentRepository.On("GetByName", mock.Anything, mock.Anything).Return(expected, nil).Once()

		users, err := sut.GetByName("name")

		mockStudentRepository.AssertExpectations(t)

		assert.NoError(t, err)
		require.Len(t, users, len(expected))
		assert.Equal(t, users[2].Id, expected[2].ID)
		assert.Equal(t, users[2].Name, expected[2].Name)
	})
	t.Run("User Repository Error", func(t *testing.T) {
		tests := []struct {
			testName      string
			expectedError error
		}{
			{
				"User Not Found",
				student_repository.ErrStudentNotFound,
			},
			{
				"Repository Internal Error",
				errors.New("internal_error"),
			},
		}

		for _, tt := range tests {
			t.Run(tt.testName, func(t *testing.T) {
				sut, mockStudentRepository := setupTest(t)
				mockStudentRepository.On("GetByName", mock.Anything, mock.Anything).Return([]*student_repository.Student{}, tt.expectedError).Once()

				user, err := sut.GetByName("name")

				mockStudentRepository.AssertExpectations(t)

				assert.ErrorIs(t, err, tt.expectedError)
				require.Nil(t, user)
			})
		}
	})
}
