package request

import (
	"fmt"
	"net/url"
)

type (
	// Paginator interface
	Paginator interface {
		GetLimit() uint
		GetCurrentPage() uint
	}

	// PaginationArgs struct
	PaginationArgs struct {
		Limit int `json:"limit"`
		Page  int `json:"page"`
	}
)

// GetLimit ...
func (s *PaginationArgs) GetLimit() int {
	if s.Limit <= 0 {
		s.Limit = 50
	}
	return s.Limit
}

// GetCurrentPage ...
func (s *PaginationArgs) GetCurrentPage() int {
	if s.Page <= 0 {
		s.Page = 1
	}
	return s.Page
}

// GetURLValues ...
func (s *PaginationArgs) GetURLValues() url.Values {
	return url.Values{
		"limit": []string{fmt.Sprintf("%d", s.GetLimit())},
		"page":  []string{fmt.Sprintf("%d", s.GetCurrentPage())},
	}
}
