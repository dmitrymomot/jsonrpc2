package response

import "encoding/json"

// Bool response struct
type Bool struct {
	Value bool `json:"-"`
}

// MarshalJSON custom implementation
func (s *Bool) MarshalJSON() ([]byte, error) {
	val := s.Value
	return json.Marshal(&val)
}
