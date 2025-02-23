package utils

import (
	"os"
	"strings"
)

// IsDifferentDomain checks if the given URL belongs to a different domain
// compared to the domain specified in the environment variable "DOMAIN".
func IsDifferentDomain(url string) bool {
	domain := os.Getenv("DOMAIN")

	if url == domain {
		return false
	}

	cleanURL := strings.TrimPrefix(url, "http://")
	cleanURL = strings.TrimPrefix(url, "https://")
	cleanURL = strings.TrimPrefix(url, "www.")
	cleanURL = strings.Split(url, "/")[0]

	return cleanURL != domain
}

// EnsureHttpPrefix ensures that a URL starts with "http://" or "https://",
// defaulting to "http://" if neither is present.
func EnsureHttpPrefix(url string) string {
	if !strings.HasPrefix(url, "http://") || !strings.HasPrefix(url, "https://") {
		return "http://" + url
	}
	return url
}
