package pingservice

import (
	"net/http"

	"github.com/emicklei/go-restful"
	. "github.com/gin-tonic/pkg/api/services/health/pingmodel"
)

// CheckService Healthcheck which checks if the service is alive
func CheckService(request *restful.Request, response *restful.Response) {

	ping := &PingResponse{Alive: true}

	response.WriteHeader(http.StatusOK)
	response.WriteEntity(ping)
}
