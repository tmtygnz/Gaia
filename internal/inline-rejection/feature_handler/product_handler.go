package feature_handler

import (
	"gaia/internal/entities"
	"gaia/internal/features/product_features"
	"gaia/utils"
	"net/http"
	"strconv"
)

type ProductFeature interface {
	FetchAllProducts() *[]entities.DProduct
	FetchProductById(id int64) *entities.DProduct
}

type ProductRestHandler struct {
	productFeatureHandler *product_features.ProductFeatureHandler
}

func NewProductRestHandler(productFeatHandler *product_features.ProductFeatureHandler) {
	handler := &ProductRestHandler{
		productFeatureHandler: productFeatHandler,
	}
	http.HandleFunc("/product", handler.FetchAllProductHandler)
	http.HandleFunc("/product/by/", handler.FetchProductByIdHandler)
}

func (restHandler *ProductRestHandler) FetchAllProductHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		allProducts := restHandler.productFeatureHandler.FetchAllProducts()

		utils.Send(writer, &allProducts, "application/json")
	}
}

func (restHandler *ProductRestHandler) FetchProductByIdHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		idStr := utils.GetRequestQuery(writer, request, "id")

		id, err := strconv.Atoi(*idStr)
		if err != nil {
			http.Error(writer, "Invalid id type", http.StatusBadRequest)
			return
		}
		product := restHandler.productFeatureHandler.FetchProductById(int64(id))

		utils.Send(writer, &product, "application/json")
	}
}
