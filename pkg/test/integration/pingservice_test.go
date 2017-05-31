/*// +build integration*/

package integration

import (
	. "github.com/gin-tonic/pkg/api/utils"
	restClient "github.com/go-resty/resty"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Pingservice Integration Tests", func() {

	var (
		configService ServiceConfig
	)

	BeforeSuite(func() {
		existingPath := "../../config/service-config.yml"
		config, err := ReadConfig(existingPath)

		if err != nil {
			panic("Pingservice Integration test cannot be configure")
		}

		configService = *config

	})

	Describe("Given endpoint", func() {
		It("should return OK if service exists", func() {
			// When
			resp, err := restClient.R().Get("http://" +
				configService.Server.Host + configService.Server.Port + "/ping")

			// Then
			Expect(err).To(BeNil())
			Expect(resp).ToNot(BeNil())
			Expect(resp.StatusCode()).To(Equal(200))
		})

		It("should return Down if service does not exist", func() {
			// When
			resp, err := restClient.R().Get("http://" +
				configService.Server.Host + configService.Server.Port + "/unknown")

			// Then
			Expect(err).To(BeNil())
			Expect(resp).ToNot(BeNil())
			Expect(resp.StatusCode()).To(Equal(404))
		})
	})
})
