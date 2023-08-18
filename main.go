package main

import (
	"encoding/json"
	"fmt"
	"go-shorturl/helper"
	"go-shorturl/model"
	"io"
	"log"
	"net/http"
)

var db map[string]string

func main() {
	db = map[string]string{}

	http.HandleFunc("/short-url", shortenUrl)
	http.HandleFunc("/long-url", getLongUrl)

	fmt.Println("Starting server..")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func shortenUrl(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Calling Shorten URL /short-url")
	request := model.Url{}

	// Get Request from Body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Validate Request
	if request.Url == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	valid := helper.IsValidUrl(request.Url)
	if !valid {
		http.Error(w, "Invalid URL format", http.StatusBadRequest)
		return
	}

	// Check if the LongUrl is Exist
	shortUrl, longUrlExist := checkLongUrl(request.Url)

	// If not exist,
	if !longUrlExist {
		// generate short_url of the longUrl then get the first 7 characters
		hashUrl := helper.GenerateShortURL(request.Url)
		shortUrl = hashUrl[:7]

		// Check if short_url exist on database
		if checkShortUrl(shortUrl) {

			// If exist, then get next 7 character from hashUrl. Until last 7 characters in hashUrl
			for i := 0; i < len(hashUrl)-7; i++ {
				shortUrl = hashUrl[i+1 : i+8]

				// Not exist on database? the end the loop
				if !checkShortUrl(shortUrl) {
					break
				}
			}
		}
		shortUrl = "http://localhost:8080/" + shortUrl
	}

	db[request.Url] = shortUrl

	response := model.Url{
		Url:      request.Url,
		ShortUrl: shortUrl,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

}

func getLongUrl(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Calling Get Long URL /long-url")

	request := model.Url{}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if request.ShortUrl == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Search by ShortURL
	url := findByShortUrl(request.ShortUrl)

	if url == "" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	response := model.Url{
		Url:      url,
		ShortUrl: request.ShortUrl,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

}

func checkShortUrl(shortUrl string) bool {
	/*
		This function check if short URL exists on database.
		Args:
			shortUrl: The short URL to check
		Returns:
			boolean, true if exist, false otherwise.
	*/
	found := false
	for _, val := range db {
		if val == shortUrl {
			found = true
			break
		}
	}

	return found
}

func findByShortUrl(shortUrl string) string {
	/*
		This function to find the original URL of a short URL.
		Args:
			shortUrl: The short URL to find the original URL for.
		Returns:
			string, The original URL, if not found will return empty string.
	*/
	found := ""
	for key, val := range db {
		if val == shortUrl {
			found = key
			break
		}
	}

	return found
}

func checkLongUrl(url string) (string, bool) {
	/*
		This function check if long URL exists on database.
		Args:
			url: The long URL to check.
		Returns:
			string and boolean, The short URL of long URL and if exists.
	*/
	val, ok := db[url]

	return val, ok
}
