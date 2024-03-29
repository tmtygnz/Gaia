package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	defect_features "mackenzie/internal/features/defect_features"
	"mackenzie/internal/features/defect_type_features"
	product_features "mackenzie/internal/features/product_features"
	inline_rejection "mackenzie/internal/inline-rejection"
	"mackenzie/provider"
)

/*
contains code that's needed for the main logic to run
*/
func startup() {
	log.Println("Loading .env")
	err := godotenv.Load(".env")
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	startup()
	databaseInstance := provider.NewDatabase()
	defectFeatureHandler := defect_features.NewDefectFeatureHandler(databaseInstance)
	productFeatureHandler := product_features.NewProductFeatureHandler(databaseInstance)
	defectTypeFeatureHandler := defect_type_features.NewDefectTypeFeatureHandler(databaseInstance)

	inline_rejection.IRRestInterface(defectFeatureHandler, productFeatureHandler, defectTypeFeatureHandler)
}
