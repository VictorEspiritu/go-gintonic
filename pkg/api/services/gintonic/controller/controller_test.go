package controller_test

import (
	"errors"

	"github.com/gin-tonic/pkg/api/persistence/mocks"
	"github.com/gin-tonic/pkg/api/persistence/model"
	"github.com/gin-tonic/pkg/api/services/gintonic/controller"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//////////////////////////////////////////////////////
////// Testing

var _ = Describe("Controller Unit Tests", func() {

	var (
		target controller.Controller
	)

	Describe("Get Tests", func() {
		It("should get a gintonic given found item and no db error", func() {
			// Given
			expectedIdFound := "idFound"
			gintonicExpected := model.Gintonic{Name: expectedIdFound, Brand: "testBrand"}

			persistenceMock := new(mocks.Persistence)
			persistenceMock.On("Get", expectedIdFound).Return(gintonicExpected, nil)

			target = controller.NewController(persistenceMock)

			// When
			gintonicResponse, err := target.FindGintonic(expectedIdFound)

			// Then
			Expect(err).To(BeNil())
			Expect(gintonicExpected).To(Equal(gintonicResponse))
		})

		It("should not get a gintonic given db error", func() {
			// Given
			errorIdFound := "errorIdFound"
			errorExpected := errors.New(errorIdFound)

			persistenceMock := new(mocks.Persistence)
			persistenceMock.On("Get", errorIdFound).Return(nil, errorExpected)

			target = controller.NewController(persistenceMock)

			// When
			gintonicResponse, errorResponse := target.FindGintonic(errorIdFound)

			// Then
			Expect(gintonicResponse).To(Not(BeNil()))
			Expect(errorExpected).To(Equal(errorResponse))
		})
	})

	Describe("Save Tests", func() {
		It("should save item given no db error", func() {
			// Given
			gintonicExpected := model.Gintonic{Name: "testId", Brand: "testBrand"}

			persistenceMock := new(mocks.Persistence)
			persistenceMock.On("Save", gintonicExpected).Return(nil)

			target = controller.NewController(persistenceMock)

			// When
			err := target.SaveGintonic(gintonicExpected)

			// Then
			Expect(err).To(BeNil())
		})

		It("should not save item given db error", func() {
			// Given
			gintonicExpected := model.Gintonic{Name: "testId", Brand: "testBrand"}
			errorExpected := errors.New("error")

			persistenceMock := new(mocks.Persistence)
			persistenceMock.On("Save", gintonicExpected).Return(errors.New("error"))

			target = controller.NewController(persistenceMock)

			// When
			err := target.SaveGintonic(gintonicExpected)

			// Then
			Expect(err).To(Equal(errorExpected))
		})
	})

})
