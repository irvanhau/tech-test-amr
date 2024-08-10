package handler

import (
	"TechnicalTest/features/products"
	"TechnicalTest/helpers"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	service products.ProductServiceInterface
}

func NewHandler(s products.ProductServiceInterface) *ProductHandler {
	return &ProductHandler{
		service: s,
	}
}

func (p *ProductHandler) GetProducts() echo.HandlerFunc {
	return func(c echo.Context) error {
		keyword := c.QueryParam("keyword")
		sort := c.QueryParam("sort")
		filter := c.QueryParam("filter")
		offset := c.QueryParam("offset")
		offsetInt, errConv := strconv.Atoi(offset)
		if errConv != nil {
			offsetInt = 0
		}

		res, paginate, err := p.service.GetProducts(sort, keyword, filter, offsetInt)
		if err != nil {
			logrus.Error("Handler : Error Get Products : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helpers.FormatResponse("Get Product Process Failed", nil, nil))
		}

		return c.JSON(http.StatusOK, helpers.FormatResponse("Success", res, paginate))
	}
}

func (p *ProductHandler) CreateProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(InputRequest)

		if err := c.Bind(input); err != nil {
			c.Logger().Error("Handler : Bind Input Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helpers.FormatResponse("Invalid User Input", nil, nil))
		}

		isValid, errValidation := helpers.ValidateJSON(input)
		if !isValid {
			return c.JSON(http.StatusBadRequest, helpers.FormatResponseValidation("Invalid Format Request", errValidation))
		}

		var serviceInput = new(products.Product)
		serviceInput.Name = input.Name
		serviceInput.Description = input.Description
		serviceInput.CategoryID = input.CategoryID

		res, errRes := p.service.CreateProduct(*serviceInput)
		if errRes != nil {
			logrus.Error("Handler : Create Product Error : ", errRes)
			return c.JSON(http.StatusInternalServerError, helpers.FormatResponse("Create Product Process Failed", nil, nil))
		}

		var response = new(InputResponse)
		response.ID = res.ID
		response.Name = res.Name
		response.CategoryID = res.CategoryID
		response.Description = res.Description
		response.CreatedAt = res.CreatedAt

		return c.JSON(http.StatusCreated, helpers.FormatResponse("Success Created", response, nil))
	}
}
