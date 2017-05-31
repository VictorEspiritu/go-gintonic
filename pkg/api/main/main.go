package main

import (
	"log"

	"net/http"

	"github.com/emicklei/go-restful"
	gintonic "github.com/gin-tonic/pkg/api/services/gintonic/endpoint"
	"github.com/gin-tonic/pkg/api/services/health/pingservice"
	"github.com/gin-tonic/pkg/api/services/user/userservice"
	utils "github.com/gin-tonic/pkg/api/utils"
)

func main() {
	config, err := utils.ReadConfig("pkg/config/service-config.yml")

	if err != nil {
		log.Fatalf("Fatal error while bootstrapping the service, %v", err)
		panic(err)
	}

	restful.Add(gintonic.New())
	restful.Add(userservice.New())
	restful.Add(pingservice.New())

	log.Printf("Deploying Microservice within port: %v", config.Server.Port)

	deployErr := http.ListenAndServe(config.Server.Port, nil)

	if deployErr != nil {
		log.Fatal(deployErr)
	}

}
