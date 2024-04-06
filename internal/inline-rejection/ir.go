package inline_rejection

import (
	"gaia/internal/features/defect_features"
	"gaia/internal/features/defect_type_features"
	"gaia/internal/features/packaging_features"
	"gaia/internal/features/product_features"
	"gaia/internal/inline-rejection/feature_handler"
	"log"
	"net/http"
)

func IRRestInterface(defectFeature *defect_features.DefectFeatureHandler,
	productFeature *product_features.ProductFeatureHandler,
	defectTypeFeature *defect_type_features.DefectTypeFeatureHandler,
	packagingQueryFeature *packaging_features.PackagingQueryHandler) {
	feature_handler.NewDefectRestHandler(defectFeature)
	feature_handler.NewProductRestHandler(productFeature)
	feature_handler.NewDefectTypeRestHandler(defectTypeFeature)
	feature_handler.NewPackagingRestHandler(packagingQueryFeature)
	log.Println("Running at :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panic("Server didn't start", err)
	}
}
