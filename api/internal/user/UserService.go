package user

import (
	"api/pkg/converters"

	"gorm.io/gorm"
)

type UserService struct {
	userRepository UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		userRepository: *NewUserRepository(),
	}
}

func (userService *UserService) GetUsers(limit int) (users []User, total int64, err error) {
	usersChan := make(chan []User)
	totalChan := make(chan int64)
	errChan := make(chan error, 2)

	go func() {
		var err error
		users, err = userService.userRepository.FindAll(limit)
		usersChan <- users
		if err != nil {
			errChan <- err
		}
	}()

	go func() {
		var err error
		total, err = userService.userRepository.Total()
		totalChan <- total
		if err != nil {
			errChan <- err
		}
	}()

	users = <-usersChan
	total = <-totalChan
	close(errChan)

	for e := range errChan {
		if e != nil {
			err = e
			break
		}
	}

	return
}

func (userService *UserService) GetUserById(id uint) (user User, err error) {
	user, err = userService.userRepository.FindById(id)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = ErrorUserNotFound
		}
		err = Error(err)
		return
	}

	if user.ID == 0 {
		err = Error(ErrorUserNotFound)
	}

	return
}

func (userService *UserService) GetUserByEmail(email string) (user User, err error) {
	user, err = userService.userRepository.FindByEmail(email)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = ErrorUserNotFound
		}
		err = Error(err)
		return
	}

	if user.ID == 0 {
		err = Error(ErrorUserNotFound)
	}

	return
}

func (userService *UserService) CreateUser(data UserFullData) (newUser User, err error) {
	var user User

	user.Email = data.Email
	user.Name = data.Name
	user.Role = UserRole(data.Role)

	newUser, err = userService.userRepository.Create(user)

	if err != nil {
		err = Error(err)
		return
	}

	if newUser.ID == 0 {
		err = Error(ErrorUserNotCreated)
	}

	return
}

func (userService *UserService) UpdateUser(id uint, data UserUpdatableData) (updatedUser User, err error) {
	user, err := userService.userRepository.FindById(id)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = ErrorUserNotFound
		}
		err = Error(err)
		return
	}

	if user.ID == 0 {
		err = Error(ErrorUserNotFound)
		return
	}

	fields, err := converters.StructToMap(data)

	if err != nil {
		err = Error(err)
		return
	}

	updatedUser, err = userService.userRepository.Update(user, fields)

	if err != nil {
		err = Error(err)
	}

	return
}

func (userService *UserService) DeleteUser(id uint) (err error) {
	user, err := userService.userRepository.FindById(id)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = ErrorUserNotFound
		}
		err = Error(err)
		return
	}

	if user.ID == 0 {
		err = Error(ErrorUserNotFound)
		return
	}

	err = userService.userRepository.Delete(id)

	if err != nil {
		err = Error(err)
	}

	return
}

func (userService *UserService) UserExistsWithEmail(email string) (exists bool, err error) {
	user, err := userService.userRepository.FindByEmail(email)
	if err == gorm.ErrRecordNotFound {
		err = nil
	}

	exists = user.ID > 0

	if err != nil {
		err = Error(err)
	}
	return
}

func (userService *UserService) MakeSureUserExists(email, name string) (user User, err error) {
	userExists, err := userService.UserExistsWithEmail(email)

	if userExists {
		user, err = userService.GetUserByEmail(email)
	} else {
		userRole, _ := UserRoleUser.Value()
		user, err = userService.CreateUser(UserFullData{
			Email: email,
			Name:  name,
			Role:  userRole.(string),
		})
	}

	if err != nil {
		err = Error(err)
	}

	return
}
