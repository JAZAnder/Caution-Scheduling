package logger

import (
	"net/http"
	"strings"
)

// ParseFilters parses query parameters from an HTTP request to populate LogsFilters.
func ParseFilters(r *http.Request) LogsFilters {
	return LogsFilters{
		Level:       strings.TrimSpace(r.URL.Query().Get("level")),
		Category:    strings.TrimSpace(r.URL.Query().Get("category")),
		SubCategory: strings.TrimSpace(r.URL.Query().Get("subCategory")),
		User:        strings.TrimSpace(r.URL.Query().Get("user")),
	}
}
