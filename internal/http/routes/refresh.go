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
	deviceService    services.DeviceService
	deviceRepository repositories.DeviceRepository
}

func NewRefreshRoute(
	deviceService services.DeviceService,
	deviceRepository repositories.DeviceRepository,
) *RefreshRoute {
	return &RefreshRoute{
		deviceService:    deviceService,
		deviceRepository: deviceRepository,
	}
}

func (receiver RefreshRoute) Method() string {
	return "POST"
}

func (receiver RefreshRoute) Pattern() string {
	return "/refresh"
}

func (receiver RefreshRoute) Handle(context *gin.Context) {
	request, errResponse := requests.NewRefreshRequest(context)

	if errResponse != nil {
		context.JSON(http.StatusBadRequest, errResponse)

		return
	}

	device := receiver.deviceRepository.GetDeviceByRefreshToken(
		request.Body.RefreshToken,
	)

	if device == nil {
		context.JSON(
			http.StatusConflict,
			responses.NewErrorConflictResponse(errors.New("invalid refresh token")),
		)

		return
	}

	device, err := receiver.deviceService.GetNewRefreshDetailsByDevice(device)

	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			responses.NewErrorInternalServerResponse(err),
		)

		return
	}

	err = receiver.deviceRepository.UpdateDevice(device)

	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			responses.NewErrorInternalServerResponse(err),
		)

		return
	}

	response, err := responses.NewSuccessRefreshResponse(device)

	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			responses.NewErrorInternalServerResponse(err),
		)

		return
	}

	context.JSON(http.StatusOK, response)
}
