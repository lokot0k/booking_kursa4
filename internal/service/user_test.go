package service_test

import (
	"errors"
	"testing"

	"meeting-room-booking/internal/domain"
	"meeting-room-booking/internal/service"
	"meeting-room-booking/internal/service/mocks"

	"github.com/stretchr/testify/require"
)

func TestUserService_GetAllUsers(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	mockUsers := []domain.User{
		{ID: 1, Name: "user1", Password: "pass1"},
		{ID: 2, Name: "user2", Password: "pass2"},
	}
	mockRepo.On("GetAll").Return(mockUsers, nil)

	svc := service.NewUserService(mockRepo)

	users, err := svc.GetAllUsers()
	require.NoError(t, err)
	require.Equal(t, mockUsers, users)

	mockRepo.AssertExpectations(t)
}

func TestUserService_GetUserByID(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	mockUser := &domain.User{ID: 1, Name: "user1", Password: "pass1"}
	mockRepo.On("GetByID", 1).Return(mockUser, nil)

	svc := service.NewUserService(mockRepo)

	user, err := svc.GetUserByID(1)
	require.NoError(t, err)
	require.Equal(t, mockUser, user)

	mockRepo.AssertExpectations(t)
}

func TestUserService_CreateUser(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	newUser := domain.User{Name: "newuser", Password: "newpass"}
	mockRepo.On("Create", newUser).Return(nil)

	svc := service.NewUserService(mockRepo)

	err := svc.CreateUser(newUser)
	require.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestUserService_Authenticate_Success(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	mockUser := &domain.User{ID: 1, Name: "user1", Password: "pass1"}
	mockRepo.On("GetByUsername", "user1").Return(mockUser, nil)

	svc := service.NewUserService(mockRepo)

	user, err := svc.Authenticate("user1", "pass1")
	require.NoError(t, err)
	require.Equal(t, mockUser, user)

	mockRepo.AssertExpectations(t)
}

func TestUserService_Authenticate_InvalidCredentials(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	mockUser := &domain.User{ID: 1, Name: "user1", Password: "pass1"}
	mockRepo.On("GetByUsername", "user1").Return(mockUser, nil)

	svc := service.NewUserService(mockRepo)

	user, err := svc.Authenticate("user1", "wrongpass")
	require.Error(t, err)
	require.Nil(t, user)
	require.EqualError(t, err, "invalid credentials")

	mockRepo.AssertExpectations(t)
}

func TestUserService_Authenticate_UserNotFound(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	mockRepo.On("GetByUsername", "user1").Return(nil, errors.New("user not found"))

	svc := service.NewUserService(mockRepo)

	user, err := svc.Authenticate("user1", "pass1")
	require.Error(t, err)
	require.Nil(t, user)
	require.EqualError(t, err, "user not found")

	mockRepo.AssertExpectations(t)
}
