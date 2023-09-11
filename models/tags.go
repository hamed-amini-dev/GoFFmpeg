package models

// Tags ...
type ITags interface {
	GetEncoder() string
}

// Tags ...
type Tags struct {
	Encoder string `json:"ENCODER"`
}

// GetEncoder ...
func (t Tags) GetEncoder() string {
	return t.Encoder
}
