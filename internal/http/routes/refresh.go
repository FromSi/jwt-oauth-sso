package routes

import (
	"errors"
	"github.com/fromsi/jwt-oauth-sso/internal/http/requests"
	"github.com/fromsi/jwt-oauth-sso/internal/http/responses"
	"github.com/fromsi/jwt-oauth-sso/internal/repositories"
	"github.com/fromsi/jwt-oauth-sso/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RefreshRoute struct {
	deviceService               services.DeviceService
	deviceRepository            repositories.DeviceRepository
	refreshRequest              requests.RefreshRequest
	successRefreshResponse      responses.SuccessRefreshResponse
	errorBadRequestResponse     responses.ErrorBadRequestResponse
	errorConflictResponse       responses.ErrorConflictResponse
	errorInternalServerResponse responses.ErrorInternalServerResponse
}

func NewRefreshRoute(
	deviceService services.DeviceService,
	deviceRepository repositories.DeviceRepository,
	refreshRequest requests.RefreshRequest,
	successRefreshResponse responses.SuccessRefreshResponse,
	errorBadRequestResponse responses.ErrorBadRequestResponse,
	errorConflictResponse responses.ErrorConflictResponse,
	errorInternalServerResponse responses.ErrorInternalServerResponse,
) *RefreshRoute {
	return &RefreshRoute{
		deviceService:               deviceService,
		deviceRepository:            deviceRepository,
		refreshRequest:              refreshRequest,
		successRefreshResponse:      successRefreshResponse,
		errorBadRequestResponse:     errorBadRequestResponse,
		errorConflictResponse:       errorConflictResponse,
		errorInternalServerResponse: errorInternalServerResponse,
	}
}

func (receiver RefreshRoute) Method() string {
	return "POST"
}

func (receiver RefreshRoute) Pattern() string {
	return "/refresh"
}

func (receiver RefreshRoute) Handle(context *gin.Context) {
	request, err := receiver.refreshRequest.Make(context)

	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			receiver.errorBadRequestResponse.Make(err),
		)

		return
	}

	device := receiver.deviceRepository.GetDeviceByRefreshToken(
		request.GetBody().GetRefreshToken(),
	)

	if device == nil {
		context.JSON(
			http.StatusConflict,
			receiver.errorConflictResponse.Make(errors.New("invalid refresh token")),
		)

		return
	}

	device, err = receiver.deviceService.GetNewRefreshDetailsByDevice(device)

	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			receiver.errorInternalServerResponse.Make(err),
		)

		return
	}

	err = receiver.deviceRepository.UpdateDevice(device)

	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			receiver.errorInternalServerResponse.Make(err),
		)

		return
	}

	response, err := receiver.successRefreshResponse.Make(device)

	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			receiver.errorInternalServerResponse.Make(err),
		)

		return
	}

	context.JSON(http.StatusOK, response)
}
