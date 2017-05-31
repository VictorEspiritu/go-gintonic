package endpoint_test

import (
	"errors"

	"github.com/emicklei/go-restful"
	service "github.com/gin-tonic/pkg/api/services/gintonic/endpoint"

	"net/http"

	"github.com/gin-tonic/pkg/api/persistence/model"
	"github.com/gin-tonic/pkg/api/services/gintonic/mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Endpoint", func() {

	var (
		target service.Service
	)

	Describe("Test Find Gintonic", func() {
		It("should find gin tonic given no error", func() {
			// Given
			expectedId := "testId"

			controllerMock := new(mocks.Gintonic)
			responseWriterMock := new(mocks.ResponseWriter)
			target = service.NewService(controllerMock)

			gintonicFound := model.NewGintonic(expectedId, "testValue")
			controllerMock.On("FindGintonic", "").Return(gintonicFound, nil)

			request := restful.NewRequest(nil)
			response := restful.NewResponse(responseWriterMock)

			restful.DefaultResponseContentType("application/json")

			// When
			target.FindGintonic(request, response)

			// Then
			Expect(response.StatusCode()).To(Equal(http.StatusOK))
		})

		It("should not find gin tonic given error", func() {
			// Given
			controllerMock := new(mocks.Gintonic)
			responseWriterMock := new(mocks.ResponseWriter)
			target = service.NewService(controllerMock)

			controllerMock.On("FindGintonic", "").
				Return(nil, errors.New("error"))

			request := restful.NewRequest(nil)
			response := restful.NewResponse(responseWriterMock)

			restful.DefaultResponseContentType("application/json")

			// When
			target.FindGintonic(request, response)

			// Then
			Expect(response.StatusCode()).To(Equal(http.StatusNoContent))
		})
	})
})
