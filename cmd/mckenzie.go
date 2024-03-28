package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	defect_features "mckenzie/internal/features/defect-features"
	product_features "mckenzie/internal/features/product-features"
	inline_rejection "mckenzie/internal/inline-rejection"
	"mckenzie/provider"
)

/*
contains code that's needed for the main logic to run
*/
func startup() {
	log.Println("Loading .env")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	startup()
	databaseInstance := provider.NewDatabase()
	defectFeatureHandler := defect_features.NewDefectFeatureHandler(databaseInstance)
	productFeatureHandler := product_features.NewProductFeatureHandler(databaseInstance)

	inline_rejection.IRRestInterface(defectFeatureHandler, productFeatureHandler)
}
