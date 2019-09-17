package response

import "encoding/json"

// List struct
type List struct {
	Object     string      `json:"object"`
	Data       interface{} `json:"data"`
	Page       int         `json:"page"`
	Limit      int         `json:"limit"`
	TotalItems int         `json:"total_items,omitempty"`
	HasMore    bool        `json:"has_more"`
}

// HasMoreItems ...
func (s *List) HasMoreItems() bool {
	if s.Limit > 0 && s.Page > 0 && s.TotalItems > 0 && s.Limit*s.Page < s.TotalItems {
		return true
	}
	return false
}

// MarshalJSON custom implementation
func (s *List) MarshalJSON() ([]byte, error) {
	if s.Page <= 0 {
		s.Page = 1
	}
	if s.Limit <= 0 {
		s.Limit = 50
	}
	if !s.HasMore {
		s.HasMore = s.HasMoreItems()
	}
	type alias List
	return json.Marshal(&struct {
		Object string `json:"object"`
		*alias
	}{
		Object: "List",
		alias:  (*alias)(s),
	})
}
