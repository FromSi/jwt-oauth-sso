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

type LoginRoute struct {
	userService      services.UserService
	deviceService    services.DeviceService
	userRepository   repositories.UserRepository
	deviceRepository repositories.DeviceRepository
}

func NewLoginRoute(
	userService services.UserService,
	deviceService services.DeviceService,
	userRepository repositories.UserRepository,
	deviceRepository repositories.DeviceRepository,
) *LoginRoute {
	return &LoginRoute{
		userService:      userService,
		deviceService:    deviceService,
		userRepository:   userRepository,
		deviceRepository: deviceRepository,
	}
}

func (receiver LoginRoute) Method() string {
	return "POST"
}

func (receiver LoginRoute) Pattern() string {
	return "/login"
}

func (receiver LoginRoute) Handle(context *gin.Context) {
	request, errResponse := requests.NewLoginRequest(context)

	if errResponse != nil {
		context.JSON(http.StatusBadRequest, errResponse)

		return
	}

	user := receiver.userRepository.GetUserByEmail(request.Body.Email)

	if user == nil {
		context.JSON(
			http.StatusConflict,
			responses.NewErrorConflictResponse(errors.New("invalid email or password")),
		)

		return
	}

	err := receiver.userService.CheckHashedPasswordAndNativePassword(
		user.GetPassword(),
		request.Body.Password,
	)

	if err != nil {
		context.JSON(
			http.StatusConflict,
			responses.NewErrorConflictResponse(errors.New("invalid email or password")),
		)

		return
	}

	device := receiver.deviceService.GetOldDeviceByUserUUIDAndIpAndUserAgent(
		user.GetUUID(),
		request.IP,
		request.UserAgent,
	)

	if device == nil {
		device, err = receiver.deviceService.GetNewDeviceByUserUUIDAndIpAndUserAgent(
			user.GetUUID(),
			request.IP,
			request.UserAgent,
		)

		if err != nil {
			context.JSON(
				http.StatusInternalServerError,
				responses.NewErrorInternalServerResponse(err),
			)

			return
		}

		err = receiver.deviceRepository.CreateDevice(device)

		if err != nil {
			context.JSON(
				http.StatusInternalServerError,
				responses.NewErrorInternalServerResponse(err),
			)

			return
		}
	}

	device, err = receiver.deviceService.GetNewRefreshDetailsByDevice(device)

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

	response, err := responses.NewSuccessLoginResponse(device)

	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			responses.NewErrorInternalServerResponse(err),
		)

		return
	}

	context.JSON(http.StatusOK, response)
}
