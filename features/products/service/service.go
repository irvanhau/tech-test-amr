package service

import (
	"TechnicalTest/features/products"
	"TechnicalTest/helpers"
	"TechnicalTest/helpers/generate_uuid"
	"errors"
	"github.com/sirupsen/logrus"
)

type ProductService struct {
	data    products.ProductDataInterface
	genUuid generate_uuid.GenerateUUIDInterface
	time    helpers.TimeInterface
}

func New(data products.ProductDataInterface, gen generate_uuid.GenerateUUIDInterface, t helpers.TimeInterface) *ProductService {
	return &ProductService{
		data:    data,
		genUuid: gen,
		time:    t,
	}
}

func (p *ProductService) GetProducts(sort, keyword, filter string, offset int) ([]products.ProductInfo, map[string]any, error) {
	res, paginate, err := p.data.GetAll(sort, keyword, filter, offset)

	if err != nil {
		logrus.Error("Service : Get Products Error : ", err.Error())
		return nil, nil, errors.New("ERROR Get All Error")
	}

	return res, paginate, nil
}

func (p *ProductService) CreateProduct(newData products.Product) (*products.Product, error) {
	newData.ID = p.genUuid.GenerateUUID()
	newData.CreatedAt = p.time.NowTime()

	res, err := p.data.Insert(newData)

	if err != nil {
		logrus.Error("Service : Create Product Error : ", err.Error())
		return nil, errors.New("ERROR Create Product Error")
	}

	return res, nil
}
