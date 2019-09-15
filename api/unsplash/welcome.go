package unsplash

import (
	"fmt"
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Wellcome to my Unsplash API Endpoint for Fondo!")
}