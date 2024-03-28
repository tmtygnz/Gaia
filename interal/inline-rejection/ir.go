package inline_rejection

import (
	defect_features "mckenzie/interal/features/defect-features"
	feature_handler "mckenzie/interal/inline-rejection/feature-handler"
	"net/http"
)

func IRRestInterface(defectFeature *defect_features.DefectFeatureHandler) {
	feature_handler.NewDefectRestHandler(defectFeature)
	http.ListenAndServe(":8080", nil)
}
