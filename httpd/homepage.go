package httpd

import (
	"fmt"
	"net/http"
)

func HomePage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Home page shit")
	}
}
