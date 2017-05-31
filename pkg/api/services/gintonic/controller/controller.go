package controller

import (
	datastore "github.com/gin-tonic/pkg/api/persistence"
	model "github.com/gin-tonic/pkg/api/persistence/model"
)

// Controller class which will allow DI (Dependency Injection)
type Controller struct {
	persistence datastore.Persistence
}

// NewController constructor
func NewController(newPersistence datastore.Persistence) Controller {
	return Controller{persistence: newPersistence}
}

// SaveGintonic saving
func (c *Controller) SaveGintonic(gintonic model.Gintonic) error {
	err := c.persistence.Save(gintonic)

	return err
}

// FinGintonic find
func (c *Controller) FindGintonic(id string) (model.Gintonic, error) {
	var response model.Gintonic
	gintonic, err := c.persistence.Get(id)

	if err != nil {
		// TODO - Do some logging
		return response, err
	}

	return gintonic.(model.Gintonic), nil
}
