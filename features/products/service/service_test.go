package service

import (
	"TechnicalTest/features/products"
	"TechnicalTest/features/products/mocks"
	mocks2 "TechnicalTest/helpers/mocks"
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	time2 "time"
)

func TestGetProducts(t *testing.T) {
	data := mocks.NewProductDataInterface(t)
	genUUID := mocks2.NewGenerateUUIDInterface(t)
	time := mocks2.NewTimeInterface(t)
	service := New(data, genUUID, time)
	product := []products.ProductInfo{}
	paginate := map[string]any{}

	t.Run("Success Get", func(t *testing.T) {
		data.On("GetAll", "", "", "", 0).Return(product, paginate, nil).Once()

		res, pagination, err := service.GetProducts("", "", "", 0)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.NotNil(t, pagination)
		data.AssertExpectations(t)
	})

	t.Run("Server Error", func(t *testing.T) {
		data.On("GetAll", "", "", "", 0).Return(nil, nil, errors.New("ERROR Get All Error")).Once()

		res, pagination, err := service.GetProducts("", "", "", 0)

		assert.Error(t, err)
		assert.EqualError(t, err, "ERROR Get All Error")
		assert.Nil(t, pagination)
		assert.Nil(t, res)
	})
}

func TestCreateProduct(t *testing.T) {
	data := mocks.NewProductDataInterface(t)
	genUUID := mocks2.NewGenerateUUIDInterface(t)
	time := mocks2.NewTimeInterface(t)
	service := New(data, genUUID, time)
	product := products.Product{
		ID:          uuid.New(),
		Name:        "asdasd",
		Description: "asdasdaasdasda",
		CategoryID:  "asdasdasdasdasda",
		CreatedAt:   time2.Now(),
	}

	t.Run("Success Create", func(t *testing.T) {
		genUUID.On("GenerateUUID").Return(product.ID).Once()
		time.On("NowTime").Return(product.CreatedAt).Once()
		data.On("Insert", product).Return(&product, nil).Once()

		res, err := service.CreateProduct(product)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, product.ID, res.ID)
		assert.Equal(t, product.Name, res.Name)
		assert.Equal(t, product.Description, res.Description)
		assert.Equal(t, product.CategoryID, res.CategoryID)
		assert.Equal(t, product.CreatedAt, res.CreatedAt)
		data.AssertExpectations(t)
	})

	t.Run("Server Error", func(t *testing.T) {
		genUUID.On("GenerateUUID").Return(product.ID).Once()
		time.On("NowTime").Return(product.CreatedAt).Once()
		data.On("Insert", product).Return(nil, errors.New("ERROR Create Product Error")).Once()

		res, err := service.CreateProduct(product)

		assert.Error(t, err)
		assert.EqualError(t, err, "ERROR Create Product Error")
		assert.Nil(t, res)
	})
}
