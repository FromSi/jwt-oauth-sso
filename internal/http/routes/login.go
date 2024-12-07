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
	userService                 services.UserService
	deviceService               services.DeviceService
	userRepository              repositories.UserRepository
	deviceRepository            repositories.DeviceRepository
	loginRequest                requests.LoginRequest
	successLoginResponse        responses.SuccessLoginResponse
	errorBadRequestResponse     responses.ErrorBadRequestResponse
	errorConflictResponse       responses.ErrorConflictResponse
	errorInternalServerResponse responses.ErrorInternalServerResponse
}

func NewLoginRoute(
	userService services.UserService,
	deviceService services.DeviceService,
	userRepository repositories.UserRepository,
	deviceRepository repositories.DeviceRepository,
	loginRequest requests.LoginRequest,
	successLoginResponse responses.SuccessLoginResponse,
	errorBadRequestResponse responses.ErrorBadRequestResponse,
	errorConflictResponse responses.ErrorConflictResponse,
	errorInternalServerResponse responses.ErrorInternalServerResponse,
) *LoginRoute {
	return &LoginRoute{
		userService:                 userService,
		deviceService:               deviceService,
		userRepository:              userRepository,
		deviceRepository:            deviceRepository,
		loginRequest:                loginRequest,
		successLoginResponse:        successLoginResponse,
		errorBadRequestResponse:     errorBadRequestResponse,
		errorConflictResponse:       errorConflictResponse,
		errorInternalServerResponse: errorInternalServerResponse,
	}
}

func (receiver LoginRoute) Method() string {
	return "POST"
}

func (receiver LoginRoute) Pattern() string {
	return "/login"
}

func (receiver LoginRoute) Handle(context *gin.Context) {
	request, err := receiver.loginRequest.Make(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, receiver.errorBadRequestResponse.Make(err))

		return
	}

	user := receiver.userRepository.GetUserByEmail(request.GetBody().GetEmail())

	if user == nil {
		context.JSON(
			http.StatusConflict,
			receiver.errorConflictResponse.Make(errors.New("invalid email or password")),
		)

		return
	}

	err = receiver.userService.CheckHashedPasswordAndNativePassword(
		user.GetPassword(),
		request.GetBody().GetPassword(),
	)

	if err != nil {
		context.JSON(
			http.StatusConflict,
			receiver.errorConflictResponse.Make(errors.New("invalid email or password")),
		)

		return
	}

	device := receiver.deviceService.GetOldDeviceByUserUUIDAndIpAndUserAgent(
		user.GetUUID(),
		request.GetIP(),
		request.GetUserAgent(),
	)

	if device == nil {
		device, err = receiver.deviceService.GetNewDeviceByUserUUIDAndIpAndUserAgent(
			user.GetUUID(),
			request.GetIP(),
			request.GetUserAgent(),
		)

		if err != nil {
			context.JSON(
				http.StatusInternalServerError,
				receiver.errorInternalServerResponse.Make(err),
			)

			return
		}

		err = receiver.deviceRepository.CreateDevice(device)

		if err != nil {
			context.JSON(
				http.StatusInternalServerError,
				receiver.errorInternalServerResponse.Make(err),
			)

			return
		}
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

	response, err := receiver.successLoginResponse.Make(device)

	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			receiver.errorInternalServerResponse.Make(err),
		)

		return
	}

	context.JSON(http.StatusOK, response)
}
