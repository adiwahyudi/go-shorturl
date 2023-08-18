package helper

import (
	"crypto/md5"

	"github.com/jxskiss/base62"
)

func GenerateShortURL(url string) string {
	/*
		This function generate the short URL from original URL.
		Step by step: Original url -> hash to MD5 -> encode to base62.
		Args:
			requestUrl: The URL to check.
		Returns:
			string, the short url.
	*/
	hashMD5 := md5.Sum([]byte(url))
	encodeBase62 := base62.EncodeToString(hashMD5[:])
	return encodeBase62
}
