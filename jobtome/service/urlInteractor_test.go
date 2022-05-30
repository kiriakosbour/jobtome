package service

import (
	"github.com/golang/mock/gomock"
	"testing"
	"urlShortener/mock_domain"
)

var mockRepo *mock_domain.MockUrlRepoInt
var service *UrlInteractor

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockRepo = mock_domain.NewMockUrlRepoInt(ctrl)
	service = UrlInteractorInit(mockRepo)
	return func() {
		service = nil
		defer ctrl.Finish()
	}
}
func Test_should_return_a_validation_error_response_when_the_request_is_nil(t *testing.T) {
	starting := setup(t)
	defer starting()
	//Act
	service = UrlInteractorInit(mockRepo)
	result := service.CreateTheUrlShortingService("")
	//Assert
	if result != "" {
		t.Error("failed while testing the url interactor")
	}
}
func Test_should_return_a_validation_error_response_when_the_retrieve_receives_empty_string(t *testing.T) {
	//Arrange
	setupFunc := setup(t)
	defer setupFunc()
	service = UrlInteractorInit(mockRepo)
	result := service.RetrieveTheUrlShortingService("")
	if result != "" {
		t.Error("failed while testing the url interactor")
	}
}
