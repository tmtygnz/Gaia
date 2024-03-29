package inline_rejection

import (
	"log"
	defect_features "mackenzie/internal/features/defect_features"
	product_features "mackenzie/internal/features/product_features"
	feature_handler "mackenzie/internal/inline-rejection/feature_handler"
	"net/http"
)

func IRRestInterface(defectFeature *defect_features.DefectFeatureHandler, productFeature *product_features.ProductFeatureHandler) {
	feature_handler.NewDefectRestHandler(defectFeature)
	feature_handler.NewProductRestHandler(productFeature)
	log.Println("Running at :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panicln("Server didn't start", err)
	}
}
