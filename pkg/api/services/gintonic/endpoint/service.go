package endpoint

import (
	"net/http"

	"github.com/emicklei/go-restful"
	model "github.com/gin-tonic/pkg/api/persistence/model"
	"github.com/gin-tonic/pkg/api/services/gintonic"
)

// Service type
type Service struct {
	controller gintonic.Gintonic
}

// NewService constructor
func NewService(newGintonicController gintonic.Gintonic) Service {
	return Service{controller: newGintonicController}
}

// FindGintonic service
func (s *Service) FindGintonic(request *restful.Request, response *restful.Response) {
	storedValue, err := s.controller.FindGintonic(request.PathParameter("gintonic-id"))

	if err != nil {
		response.WriteHeader(http.StatusNoContent)
	} else {
		response.WriteHeader(http.StatusOK)
		response.WriteEntity(storedValue)
	}

}

// SaveGintonic service
func (s *Service) SaveGintonic(request *restful.Request, response *restful.Response) {

	gintonic := new(model.Gintonic)
	readEntityErr := request.ReadEntity(&gintonic)

	if readEntityErr != nil {
		response.WriteHeader(http.StatusInternalServerError)
	} else {
		err := s.controller.SaveGintonic(*gintonic)
		if err != nil {
			response.WriteHeader(http.StatusNoContent)
		} else {
			response.WriteHeader(http.StatusOK)
			response.WriteEntity(gintonic)
		}
	}
}
