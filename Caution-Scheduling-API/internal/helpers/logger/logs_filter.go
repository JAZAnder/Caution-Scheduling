
package handlers

import (
    "net/http"
    "strings"
)

// LogsFilters defines the possible fields that can be used to filter logs.
type LogsFilters struct {
    Level       string
    Category    string
    SubCategory string
    User        string
}

// ParseFilters parses query parameters from an HTTP request to populate LogsFilters.
func ParseFilters(r *http.Request) LogsFilters {
    return LogsFilters{
        Level:       strings.TrimSpace(r.URL.Query().Get("level")),
        Category:    strings.TrimSpace(r.URL.Query().Get("category")),
        SubCategory: strings.TrimSpace(r.URL.Query().Get("subCategory")),
        User:        strings.TrimSpace(r.URL.Query().Get("user")),
    }
}
