package utils_test

import (
	. "github.com/gin-tonic/pkg/api/utils"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Utils unit tests", func() {

	Describe("Read config file", func() {
		Context("With specified config file path", func() {
			It("should exist and can be marshalled", func() {
				// Given
				existingPath := "../../config/service-config.yml"

				// When
				config, err := ReadConfig(existingPath)

				// Then
				Expect(err).To(BeNil())
				Expect(config.Server).To(Not(BeNil()))
				Expect(config.Server.Port).To(Not(BeNil()))
				Expect(config.Server.Host).To(Not(BeNil()))
			})

			It("should exist and can be not marshalled", func() {
				// Given
				existingPath := "../../config/service-config-wrong.yml"

				// When
				config, err := ReadConfig(existingPath)

				// Then
				Expect(err).To(Not(BeNil()))
				Expect(config).To(BeNil())
			})

			It("should not exist", func() {
				// Given
				existingPath := "../../config/service-config-non-exist.yml"

				// When
				config, err := ReadConfig(existingPath)

				// Then
				Expect(err).To(Not(BeNil()))
				Expect(config).To(BeNil())
			})

		})

		Context("With non specified config file", func() {
			It("should not exist", func() {
				// When
				config, err := ReadConfig("")

				// Then
				Expect(err).To(Not(BeNil()))
				Expect(config).To(BeNil())
			})
		})
	})
})
