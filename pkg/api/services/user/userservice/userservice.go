package userservice

import (
	"github.com/emicklei/go-restful"
	. "github.com/gin-tonic/pkg/api/services/user/usermodel"
)

func New() *restful.WebService {
	service := new(restful.WebService)
	service.
		Path("/users").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_XML, restful.MIME_JSON)

	service.Route(service.GET("/{user-id}").To(FindUser).Writes(User{}))
	service.Route(service.POST("").To(UpdateUser).Writes(User{}))
	service.Route(service.PUT("/{user-id}").To(CreateUser).Writes(User{}))
	service.Route(service.DELETE("/{user-id}").To(RemoveUser).Writes(User{}))

	return service
}

