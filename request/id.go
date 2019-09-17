package request

type (
	// ID struct
	ID struct {
		ID string `json:"id"`
	}
)

// Rules ...
func (ID) Rules() map[string][]string {
	return map[string][]string{
		"id": []string{"required"},
	}
}

// Messages ...
func (ID) Messages() map[string][]string {
	return nil
}
