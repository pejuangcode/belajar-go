package service

import (
	"context"
	"fmt"
	"myapp/graph/model"
	"myapp/tools"

	"gorm.io/gorm"
)

func CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	var (
		password = tools.HashAndSalt([]byte(input.Password))

		user = model.User{
			Name:     input.Name,
			Password: password,
			Email:    input.Email,
		}
	)

	err := DB.Model(&user).Create(&user).Error
	if err != nil {
		panic(err)
	}
	fmt.Println(&user)

	return &user, nil

}

func UpdateUser(ctx context.Context, id int, input model.UpdateUser) (string, error) {
	var (
		user = model.User{
			Name:  *input.Name,
			Email: *input.Email,
		}
	)

	err := DB.Model(&user).Where("id = ?", id).Updates(&user).Error
	if err != nil {
		panic(err)
	}

	return "success", nil
}

func DelteUser(ctx context.Context, id int) (string, error) {
	var (
		user model.User
	)

	err := DB.Model(&user).Where("id = ?", id).Delete(&user).Error
	if err != nil {
		panic(err)
	}

	return "success", nil
}

func GetAllUser(ctx context.Context) ([]*model.User, error) {
	var (
		users []*model.User
	)

	err := DB.Model(&users).Find(&users).Error
	if err != nil {
		panic(err)
	}

	return users, nil
}

func GetUserById(ctx context.Context, id int) (*model.User, error) {
	var (
		user model.User
	)

	err := DB.Model(&user).Where("id = ?", id).Take(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	} else if err != nil {
		panic(err)
	}

	return &user, nil
}
