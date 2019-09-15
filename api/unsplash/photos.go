package unsplash

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Photos(w http.ResponseWriter, r *http.Request) {
	/*
	 * Build the basic URL
	 */
	url := 	"https://api.unsplash.com" +
			"/photos" +
			"?client_id=51531311dfa090ab81321cd2655e73c59b3d952b5966ed42e861fa7d50da47e8"

	/*
	 * Adding params from request
	 */
	page := r.URL.Query().Get("page")
	if page != "" {
		url = url + "&page=" + page
	}

	perPage := r.URL.Query().Get("per_page")
	if perPage != "" {
		url = url + "&per_page=" + perPage
	}

	orderBy := r.URL.Query().Get("order_by")
	if orderBy != "" {
		url = url + "&order_by=" + orderBy
	}

	/*
	 * Make the request to the Unsplash API
	 */
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	/*
	 * Try to get the response from the Unsplash API
	 */
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	/*
	 * Give the response to the client in Json
	 */
	w.Header().Set("Content-Type", "application/json")
	_, _ = fmt.Fprintf(w, string(body))
}
