package inline_rejection

import "net/http"

func IRRestInterface() {

	http.ListenAndServe(":8080", nil)
}
