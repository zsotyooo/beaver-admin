package services

import (
	UserModel "api/internal/modules/user/models"
	UserRepository "api/internal/modules/user/repositories"
	UserRequest "api/internal/modules/user/requests"
	UserResponse "api/internal/modules/user/responses"
	"errors"

	"api/packages/converters"

	"gorm.io/gorm"
)

type UserService struct {
	userRepository UserRepository.UserRepositoryInterface
}

func New() *UserService {
	return &UserService{
		userRepository: UserRepository.New(),
	}
}

func (userService *UserService) GetUsers(limit int) (UserResponse.UsersResponse, error) {
	users, err := userService.userRepository.List(limit)

	if err != nil {
		return UserResponse.UsersResponse{}, err
	}

	return UserResponse.ConvertModelsToResponse(users), nil
}

func (userService *UserService) FindUser(id uint) (UserResponse.UserResponse, error) {
	var response UserResponse.UserResponse

	user, err := userService.userRepository.Find(id)

	if err != nil {
		return response, err
	}

	if user.ID == 0 {
		return response, errors.New("User not found!")
	}

	return UserResponse.ConvertModelToResponse(user), nil
}

func (userService *UserService) FindUserByEmail(email string) (UserResponse.UserResponse, error) {
	var response UserResponse.UserResponse

	user, err := userService.userRepository.FindByEmail(email)

	if err != nil {
		return response, err
	}

	if user.ID == 0 {
		return response, errors.New("User not found!")
	}

	return UserResponse.ConvertModelToResponse(user), nil
}

func (userService *UserService) CreateUser(payload UserRequest.UserCreatePayload) (UserResponse.UserResponse, error) {
	var user UserModel.User
	var response UserResponse.UserResponse

	user.Email = payload.Email
	user.Name = payload.Name
	user.Role = UserModel.UserRole(payload.Role)

	newUser, err := userService.userRepository.Create(user)

	if err != nil {
		return response, err
	}

	if newUser.ID == 0 {
		return response, errors.New("Error in creating the user!")
	}

	return UserResponse.ConvertModelToResponse(newUser), nil
}

func (userService *UserService) UpdateUser(id uint, payload UserRequest.UserUpdatePayload) (UserResponse.UserResponse, error) {
	var response UserResponse.UserResponse
	user, err := userService.userRepository.Find(id)

	if err != nil {
		return response, err
	}

	if user.ID == 0 {
		return response, errors.New("User not found!")
	}

	fields, err := converters.StructToMap(payload)

	if err != nil {
		return response, err
	}

	updatedUser, err := userService.userRepository.Update(user, fields)

	if err != nil {
		return response, err
	}

	return UserResponse.ConvertModelToResponse(updatedUser), nil
}

func (userService *UserService) DeleteUser(id uint) error {
	user, err := userService.userRepository.Find(id)

	if err != nil {
		return err
	}

	if user.ID == 0 {
		return errors.New("User not found!")
	}

	userService.userRepository.Delete(id)

	return nil
}

func (userService *UserService) UserExists(email string) (exists bool, err error) {
	user, err := userService.userRepository.FindByEmail(email)
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	exists = user.ID > 0
	return
}
