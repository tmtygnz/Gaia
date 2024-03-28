package feature_handler

import (
	"encoding/json"
	"log"
	"mckenzie/interal/entities"
	product_features "mckenzie/interal/features/product-features"
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
	http.HandleFunc("/prodct", handler.FetchAllProductHandler)
	http.HandleFunc("/product/by/{id}", handler.FetchProductByIdHandler)
}

func (restHandler *ProductRestHandler) FetchAllProductHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		allProducts := restHandler.productFeatureHandler.FetchAllProducts()

		allProductsBytes, err := json.Marshal(allProducts)
		if err != nil {
			log.Println("Error marshaling all products", err)
			writer.WriteHeader(http.StatusInternalServerError)
		}

		writer.Write(allProductsBytes)
	}
}

func (restHandler *ProductRestHandler) FetchProductByIdHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		idStr := request.URL.Query().Get("id")
		if idStr == "" {
			writer.WriteHeader(http.StatusUnprocessableEntity)
			writer.Write([]byte("Id parameter is missing"))
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
		}
		product := restHandler.productFeatureHandler.FetchProductById(int64(id))

		productBytes, err := json.Marshal(product)
		if err != nil {
			log.Println("Can't marshal product", err)
			writer.WriteHeader(http.StatusInternalServerError)
		}

		writer.Write(productBytes)
	}
}
