package model

// Gintonic structure
type Gintonic struct {
	Brand string `json:"brand"`
	Name  string `json:"name"`
}

// NewGintonic constructor
func NewGintonic(name string, brand string) Gintonic {
	return Gintonic{Name: name, Brand: brand}
}

// GetKey which will return the ID
func (g Gintonic) GetKey() string {
	return g.Name
}

// GetValue which will return the value
func (g Gintonic) GetValue() string {
	return g.Brand
}
