package pingservice

import (
	"github.com/emicklei/go-restful"
	. "github.com/gin-tonic/pkg/api/services/health/pingmodel"
)

// New Pingservice constructor
func New() *restful.WebService {
	service := new(restful.WebService)
	service.
		Path("/ping").
		Consumes(restful.MIME_JSON, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_JSON)

	service.Route(service.GET("").To(CheckService).Writes(PingResponse{}))

	return service
}
