package pingmodel

// PingResponse response which returns if the service is up.
type PingResponse struct {
	Alive       bool   `json:"alive"`
	Description string `json:"description,omitempty"`
}
