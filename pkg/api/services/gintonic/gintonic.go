package gintonic

import "github.com/gin-tonic/pkg/api/persistence/model"

type Gintonic interface {
	// SaveGintonic will save one gintonic
	SaveGintonic(gintonic model.Gintonic) error
	// FindGintonic will find one gintonic within the datamodel
	FindGintonic(id string) (model.Gintonic, error)
}
