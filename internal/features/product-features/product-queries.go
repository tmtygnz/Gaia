package product_features

import (
	"github.com/go-jet/jet/v2/postgres"
	"log"
	"mckenzie/internal/entities"
	"mckenzie/internal/jet/postgres/public/table"
	"mckenzie/provider"
)

type ProductFeatureHandler struct {
	db provider.IDBProvider
}

func NewProductFeatureHandler(db provider.IDBProvider) *ProductFeatureHandler {
	return &ProductFeatureHandler{db: db}
}

/*
FetchAllProducts fetches all products in the database
*/
func (handler *ProductFeatureHandler) FetchAllProducts() *[]entities.DProduct {
	var tfo = new([]entities.DProduct)
	stmt := table.Products.SELECT(table.Products.AllColumns).FROM(table.Products)
	if err := handler.db.Query(stmt, tfo); err != nil {
		log.Panicln("An error occurred when fetching all products", err)
		return nil
	}
	return tfo
}

/*
FetchProductById fetches product with the same id as the parameter
*/
func (handler *ProductFeatureHandler) FetchProductById(id int64) *entities.DProduct {
	tfo := new(entities.DProduct)
	stmt := table.Products.SELECT(table.Products.AllColumns).FROM(table.Products).
		WHERE(table.Products.ID.EQ(postgres.Int64(id)))
	if err := handler.db.Query(stmt, tfo); err != nil {
		log.Panicln("An error occurred when fetching product by id", err)
		return nil
	}
	return tfo
}
