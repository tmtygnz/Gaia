package inline_rejection

import (
	"gaia/internal/features/area_features"
	"gaia/internal/features/defect_features"
	"gaia/internal/features/defect_type_features"
	"gaia/internal/features/packaging_features"
	"gaia/internal/features/product_features"
	"gaia/internal/inline-rejection/feature_handler"
	"log"
	"net/http"
)

// TODO: Move these parameters in a interface or struct
func IRRestInterface(defectFeature *defect_features.DefectQueryFeatureHandler,
	defectCommandFeature *defect_features.DefectCommandFeatureHandler,
	productFeature *product_features.ProductQueryFeatureImpl,
	defectTypeFeature *defect_type_features.DefectTypeQueryFeatureImpl,
	packagingQueryFeature *packaging_features.PackagingQueryFeatureImpl,
	areaQueryFeature *area_features.AreaQueryFeatureImpl) {

	feature_handler.NewDefectRestHandler(defectFeature, defectCommandFeature)
	feature_handler.NewProductRestHandler(productFeature)
	feature_handler.NewDefectTypeRestHandler(defectTypeFeature)
	feature_handler.NewPackagingRestHandler(packagingQueryFeature)
	feature_handler.NewAreaRestHandler(areaQueryFeature)

	log.Println("Running at :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		// Just return if it errors out
		// TODO: Save logs to disk if possible
		log.Println("Server didn't start", err)
		return
	}
}
