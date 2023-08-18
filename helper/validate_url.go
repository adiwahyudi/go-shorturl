package helper

import (
	"net/url"
	"regexp"
)

// Valid URL i.e https://google.com, http://google.com
func IsValidUrl(requestUrl string) bool {
	/*
	   This function checks if the URL is valid.
	   Args:
	     requestUrl: The URL to check.
	   Returns:
	     True if the URL is valid, False otherwise.
	*/

	parsedUrl, err := url.Parse(requestUrl)
	if err != nil {
		return false
	}

	return parsedUrl.Scheme != "" && parsedUrl.Host != "" && isValidHost(parsedUrl.Host)
}

func isValidHost(host string) bool {
	/*
	   This function checks if the host of URL is valid.
	   Args:
	     host: The host to check.
	   Returns:
	     True if the URL is valid, False otherwise.
	*/

	regex := regexp.MustCompile(`^[a-zA-Z0-9-_.]+\.[a-zA-Z]{2,6}$`)
	return regex.MatchString(host)
}
