package inline_rejection

import (
	"log"
	defect_features "mckenzie/internal/features/defect-features"
	product_features "mckenzie/internal/features/product-features"
	feature_handler "mckenzie/internal/inline-rejection/feature-handler"
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
