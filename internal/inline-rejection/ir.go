package inline_rejection

import (
	"log"
	"mackenzie/internal/features/defect_features"
	"mackenzie/internal/features/defect_type_features"
	"mackenzie/internal/features/product_features"
	"mackenzie/internal/inline-rejection/feature_handler"
	"net/http"
)

func IRRestInterface(defectFeature *defect_features.DefectFeatureHandler, productFeature *product_features.ProductFeatureHandler, defectTypeFeature *defect_type_features.DefectTypeFeatureHandler) {
	feature_handler.NewDefectRestHandler(defectFeature)
	feature_handler.NewProductRestHandler(productFeature)
	feature_handler.NewDefectTypeRestHandler(defectTypeFeature)
	log.Println("Running at :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panic("Server didn't start", err)
	}
}
