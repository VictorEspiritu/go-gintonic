package endpoint

import (
	"github.com/emicklei/go-restful"
	elasticsearch "github.com/gin-tonic/pkg/api/persistence/elasticsearch"
	model "github.com/gin-tonic/pkg/api/persistence/model"
	"github.com/gin-tonic/pkg/api/services/gintonic/controller"
)

// New constructor
func New() *restful.WebService {
	service := new(restful.WebService)
	service.
		Path("/gintonic").
		Consumes(restful.MIME_JSON, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_JSON)

	persistence := elasticsearch.New("Host")
	controller := controller.NewController(persistence)

	// Controller at service level
	client := NewService(&controller)

	service.Route(service.GET("/{gintonic-id}").To(client.FindGintonic).Writes(model.Gintonic{}))
	service.Route(service.PUT("").To(client.SaveGintonic).Writes(model.Gintonic{}))

	return service
}
