package unsplash

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Download(w http.ResponseWriter, r *http.Request) {
	/*
	 * Getting the photos' id
	 */
	id := r.URL.Query().Get("id")

	/*
	 * Build the basic URL
	 */
	url := 	"https://api.unsplash.com" +
		fmt.Sprintf("/photos/%s/download", id) +
		"?client_id=51531311dfa090ab81321cd2655e73c59b3d952b5966ed42e861fa7d50da47e8"

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
