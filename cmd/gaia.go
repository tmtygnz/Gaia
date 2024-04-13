package main

import (
	"gaia/internal/features/area_features"
	defect_features "gaia/internal/features/defect_features"
	"gaia/internal/features/defect_type_features"
	"gaia/internal/features/packaging_features"
	product_features "gaia/internal/features/product_features"
	inline_rejection "gaia/internal/inline-rejection"
	"gaia/provider"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
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
	defectFeatureHandler := defect_features.NewDefectQueryFeatureHandler(databaseInstance)
	defectCommandFeatureHandler := defect_features.NewDefectCommandFeatureHandler(databaseInstance)
	productFeatureHandler := product_features.NewProductFeatureHandler(databaseInstance)
	defectTypeFeatureHandler := defect_type_features.NewDefectTypeFeatureHandler(databaseInstance)
	packagingQueryFeatureHandler := packaging_features.NewPackagingTypeQueriesHandler(databaseInstance)
	areaQueryFeatureHandler := area_features.NewAreaFeatureHandler(databaseInstance)

	inline_rejection.IRRestInterface(defectFeatureHandler, defectCommandFeatureHandler, productFeatureHandler, defectTypeFeatureHandler,
		packagingQueryFeatureHandler, areaQueryFeatureHandler)
}
