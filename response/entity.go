package response

import (
	"encoding/json"
)

// Entity struct
type Entity struct {
	Object string      `json:"-"`
	Entity interface{} `json:"-"`
}

// MarshalJSON custom implementation
func (s *Entity) MarshalJSON() ([]byte, error) {
	type alias interface{}
	return json.Marshal(&struct {
		Object string `json:"object"`
		alias
	}{
		Object: s.Object,
		alias:  (alias)(s.Entity),
	})
}
