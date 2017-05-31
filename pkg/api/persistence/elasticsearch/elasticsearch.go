package elasticsearch

import datastore "github.com/gin-tonic/pkg/api/persistence"

// Client type of client with its vars
// TODO - We need to use pointers
type Client struct {
	storagemap map[string]string
	host       string
}

// New constructor
func New(host string) datastore.Persistence {

	newClient := new(Client)

	newClient.storagemap = make(map[string]string)
	newClient.host = host

	return newClient
}

// Save method which will save the data
func (es *Client) Save(r datastore.Storable) error {
	es.storagemap[r.GetKey()] = r.GetValue()
	return nil
}

// Get function which will return some data
func (es *Client) Get(id string) (datastore.Storable, error) {
	return nil, nil
}
