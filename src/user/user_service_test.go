package user

import (
	"context"
	"errors"
	"testing"

	"github.com/pact/core/database/repository"
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

		ctx := t.Context()
		mockStudentRepository.On("Create", ctx, &student_repository.CreateStudentParams{
			Name: expected.Name, Email: expected.Email,
		}).Return(expected, nil).Once()

		result, err := sut.CreateUser(ctx, expected.Name, expected.Email)

		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Equal(t, &User{Id: expected.ID, Name: expected.Name, Email: expected.Email}, result)
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

				user, err := sut.CreateUser(t.Context(), expected.Name, expected.Email)

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

		ctx := t.Context()
		mockStudentRepository.On("GetByID", ctx, expected.ID).Return(expected, nil).Once()

		user, err := sut.GetUser(ctx, expected.ID)

		mockStudentRepository.AssertExpectations(t)

		require.NoError(t, err)
		require.NotNil(t, user)
		assert.Equal(t, &User{Id: expected.ID, Name: expected.Name, Email: expected.Email}, user)
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

				user, err := sut.GetUser(t.Context(), "id")

				mockStudentRepository.AssertExpectations(t)

				assert.ErrorIs(t, err, tt.expectedError)
				require.Nil(t, user)
			})
		}
	})
}

func TestListUsers(t *testing.T) {
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

		ctx, cancel := context.WithCancel(t.Context())
		defer cancel()
		params := ListUserParams{Name: "doe", Page: 2, Limit: 10}
		mockStudentRepository.On("GetUsers", ctx,
			&student_repository.GetUsersParams{Name: params.Name},
			&repository.Pagination{Page: params.Page, Limit: params.Limit},
		).Return(expected, nil).Once()

		users, err := sut.ListUsers(ctx, params)

		mockStudentRepository.AssertExpectations(t)

		require.NoError(t, err)
		require.Len(t, users, len(expected))
		for i, student := range expected {
			assert.Equal(t, User{Id: student.ID, Name: student.Name, Email: student.Email}, users[i])
		}
	})
	t.Run("No Users Found", func(t *testing.T) {
		sut, mockStudentRepository := setupTest(t)
		ctx := t.Context()
		params := ListUserParams{Limit: 10}
		mockStudentRepository.On("GetUsers", ctx,
			&student_repository.GetUsersParams{},
			&repository.Pagination{Limit: params.Limit},
		).Return([]*student_repository.Student{}, nil).Once()

		users, err := sut.ListUsers(ctx, params)

		require.NoError(t, err)
		assert.NotNil(t, users)
		assert.Empty(t, users)
		mockStudentRepository.AssertExpectations(t)
	})
	t.Run("User Repository Error", func(t *testing.T) {
		tests := []struct {
			testName      string
			expectedError error
		}{
			{
				"Repository Internal Error",
				errors.New("internal_error"),
			},
		}

		for _, tt := range tests {
			t.Run(tt.testName, func(t *testing.T) {
				sut, mockStudentRepository := setupTest(t)
				mockStudentRepository.On("GetUsers", mock.Anything, mock.Anything, mock.Anything).Return([]*student_repository.Student{}, tt.expectedError).Once()

				user, err := sut.ListUsers(t.Context(), ListUserParams{Limit: 10})

				mockStudentRepository.AssertExpectations(t)

				assert.ErrorIs(t, err, tt.expectedError)
				require.Nil(t, user)
			})
		}
	})
}
