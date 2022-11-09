package service

import (
	"context"
	"myapp/graph/model"

	"gorm.io/gorm"
)

func CreateProduct(ctx context.Context, input model.NewProduct) (*model.Product, error) {
	var (
		product = model.Product{
			Code:        input.Code,
			Description: input.Description,
			Quantity:    input.Quantity,
		}
	)

	err := DB.Model(&product).Create(&product).Error
	if err != nil {
		panic(err)
	}

	return &product, nil

}

func UpdateProduct(ctx context.Context, id int, input model.UpdateProduct) (string, error) {
	var (
		product = model.Product{
			Code:        *input.Code,
			Description: *input.Description,
			Quantity:    *input.Quantity,
		}
	)

	err := DB.Model(&product).Where("id = ?", id).Updates(&product).Error
	if err != nil {
		panic(err)
	}

	return "success", nil
}

func DelteProduct(ctx context.Context, id int) (string, error) {
	var (
		product model.Product
	)

	err := DB.Model(&product).Where("id = ?", id).Delete(&product).Error
	if err != nil {
		panic(err)
	}

	return "success", nil
}

func GetAllProduct(ctx context.Context) ([]*model.Product, error) {
	var (
		products []*model.Product
	)

	err := DB.Model(&products).Find(&products).Error
	if err != nil {
		panic(err)
	}

	return products, nil
}

func GetProductById(ctx context.Context, id int) (*model.Product, error) {
	var (
		product model.Product
	)

	err := DB.Model(&product).Where("id = ?", id).Take(&product).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	} else if err != nil {
		panic(err)
	}

	return &product, nil
}

func GetProductByUserId(ctx context.Context) ([]*model.Product, error) {

	var (
		users []*model.User
	)

	var (
		products []*model.Product
	)

	err := DB.Model(&users).Find(&users).Error
	if err != nil {
		panic(err)
	}

	for _, user := range users {
		var product []*model.Product
		DB.Where(&model.Product{UserID: user.ID}).Find(&product)
		products = product
	}

	return products, nil
}
