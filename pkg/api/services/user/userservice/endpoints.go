package userservice

import (
	"net/http"

	"github.com/emicklei/go-restful"
	. "github.com/gin-tonic/pkg/api/services/user/usermodel"
)

func FindUser(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("user-id")
	// here you would fetch user from some persistence system
	usr := &User{Id: id, Name: "John Doe"}
	response.WriteEntity(usr)
}

func UpdateUser(request *restful.Request, response *restful.Response) {
	usr := new(User)
	err := request.ReadEntity(&usr)
	// here you would update the user with some persistence system
	if err == nil {
		response.WriteEntity(usr)
	} else {
		response.WriteError(http.StatusInternalServerError, err)
	}
}

func CreateUser(request *restful.Request, response *restful.Response) {
	usr := User{Id: request.PathParameter("user-id")}
	err := request.ReadEntity(&usr)
	// here you would create the user with some persistence system
	if err == nil {
		response.WriteEntity(usr)
	} else {
		response.WriteError(http.StatusInternalServerError, err)
	}
}

func RemoveUser(request *restful.Request, response *restful.Response) {
	// here you would delete the user from some persistence system
}
