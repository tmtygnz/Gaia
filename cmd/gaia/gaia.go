package main

import (
	"gaia/internal/features/area_features"
	defect_features "gaia/internal/features/defect_features"
	"gaia/internal/features/defect_type_features"
	"gaia/internal/features/packaging_features"
	product_features "gaia/internal/features/product_features"
	inline_rejection "gaia/internal/inline-rejection"
	"gaia/provider"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
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
	defectQueryFeatureHandler := defect_features.NewDefectQueryFeature(databaseInstance)
	defectCommandFeatureHandler := defect_features.NewDefectCommandFeature(databaseInstance)
	productQueryFeatureHandler := product_features.NewProductQueryFeature(databaseInstance)
	defectTypeQueryFeatureHandler := defect_type_features.NewDefectTypeQueryFeature(databaseInstance)
	packagingQueryFeatureHandler := packaging_features.NewPackagingTypeQuery(databaseInstance)
	areaQueryFeatureHandler := area_features.NewAreaQueryFeature(databaseInstance)

	inline_rejection.IRRestInterface(defectQueryFeatureHandler, defectCommandFeatureHandler, productQueryFeatureHandler, defectTypeQueryFeatureHandler,
		packagingQueryFeatureHandler, areaQueryFeatureHandler)
}
