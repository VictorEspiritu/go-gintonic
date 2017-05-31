package persistence

// Storable interface
type Storable interface {
	GetKey() string
	GetValue() string
}

// Persistence interface
type Persistence interface {
	Save(r Storable) error
	Get(id string) (Storable, error)
}
