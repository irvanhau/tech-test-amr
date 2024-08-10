package data

import (
	"TechnicalTest/features/products"
	"TechnicalTest/helpers"
	"TechnicalTest/utils/cache"
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"math"
	"strconv"
)

type ProductData struct {
	db    *gorm.DB
	redis cache.Redis
}

func New(db *gorm.DB, r cache.Redis) *ProductData {
	return &ProductData{db: db, redis: r}
}

func (p *ProductData) GetAll(sort, keyword, filter string, offset int) ([]products.ProductInfo, map[string]any, error) {
	var listProducts []products.ProductInfo
	var totalData int64
	var qry = p.db.Table("products").
		Select("products.*", "product_categories.id as category_id, product_categories.name as category_name").
		Joins("LEFT JOIN product_categories ON product_categories.id = products.category_id")

	if keyword != "" {
		qry.Where("products.name LIKE ?", "%"+keyword+"%")
	}

	if filter != "" {
		qry.Where("product_categories.id =?", filter)
	}

	if sort != "" {
		qry.Order("products.created_at " + sort)
	}

	qry.Count(&totalData)

	if offset > 0 {
		qry.Offset(offset)
	} else {
		qry.Offset(0)
	}

	qry.Limit(10)

	pagination := p.paginate(totalData, 10, offset)

	key := "product:limit:10:offset:" + strconv.Itoa(offset) + ":keyword:" + keyword + ":filter:" + filter + ":sort:" + sort
	products, err := p.redis.Get(context.Background(), key)
	if err != nil {
		if err = qry.Scan(&listProducts).Error; err != nil {
			logrus.Error("DATA : Get All Error : ", err.Error())
			return nil, nil, err
		}
		err = p.redis.Set(context.Background(), key, listProducts)
		if err != nil {
			return nil, nil, err
		}
		return listProducts, pagination, nil
	}

	err = json.Unmarshal([]byte(products), &listProducts)
	if err != nil {
		return nil, nil, err
	}

	return listProducts, pagination, nil

}

func (p *ProductData) Insert(newData products.Product) (*products.Product, error) {
	var dbData = new(Product)
	dbData.ID = newData.ID
	dbData.Name = newData.Name
	dbData.CategoryID = newData.CategoryID
	dbData.Description = newData.Description
	dbData.CreatedAt = newData.CreatedAt

	if err := p.db.Create(dbData).Error; err != nil {
		logrus.Error("DATA : Insert Error : ", err.Error())
		return nil, err
	}

	return &newData, nil
}

func fetchData(qry *gorm.DB) {

}

func (p *ProductData) paginate(total int64, limit, offset int) map[string]any {
	var totalPage, currentPage, lastPage, nextPage float64

	totalPage = math.Round(float64(total) / float64(limit))

	if offset == 0 {
		currentPage = 1
	} else {
		currentPage = math.Round(float64(offset/limit + 1))
	}

	if lastPage = currentPage - 1; lastPage < 0 {
		lastPage = 0
	}

	if nextPage = currentPage + 1; nextPage > totalPage {
		nextPage = totalPage
	}

	paginate := helpers.FormatPaginate(totalPage, limit, total, currentPage, lastPage, nextPage)
	return paginate
}
