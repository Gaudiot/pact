package utils

import (
	"fmt"
	"net/http"
	"strconv"
)

// Parse the request to get the pagination data of the request.
//
// It returns (page, limit, error). If not provided page defaults to 0 and limit to 10.
func ParsePagination(r *http.Request) (int, int, error) {
	page := 0
	limit := 10

	if value := r.URL.Query().Get("page"); value != "" {
		parsed, err := strconv.Atoi(value)
		if err != nil || parsed < 0 {
			return 0, 0, fmt.Errorf("invalid page")
		}
		page = parsed
	}

	if value := r.URL.Query().Get("limit"); value != "" {
		parsed, err := strconv.Atoi(value)
		if err != nil || parsed <= 0 {
			return 0, 0, fmt.Errorf("invalid limit")
		}
		limit = parsed
	}

	return page, limit, nil
}
