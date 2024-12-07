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

type RegisterRoute struct {
	userRepository              repositories.UserRepository
	deviceRepository            repositories.DeviceRepository
	userService                 services.UserService
	deviceService               services.DeviceService
	registerRequest             requests.RegisterRequest
	successRegisterResponse     responses.SuccessRegisterResponse
	errorBadRequestResponse     responses.ErrorBadRequestResponse
	errorConflictResponse       responses.ErrorConflictResponse
	errorInternalServerResponse responses.ErrorInternalServerResponse
}

func NewRegisterRoute(
	userRepository repositories.UserRepository,
	deviceRepository repositories.DeviceRepository,
	userService services.UserService,
	deviceService services.DeviceService,
	registerRequest requests.RegisterRequest,
	successRegisterResponse responses.SuccessRegisterResponse,
	errorBadRequestResponse responses.ErrorBadRequestResponse,
	errorConflictResponse responses.ErrorConflictResponse,
	errorInternalServerResponse responses.ErrorInternalServerResponse,
) *RegisterRoute {
	return &RegisterRoute{
		userRepository:              userRepository,
		deviceRepository:            deviceRepository,
		userService:                 userService,
		deviceService:               deviceService,
		registerRequest:             registerRequest,
		successRegisterResponse:     successRegisterResponse,
		errorBadRequestResponse:     errorBadRequestResponse,
		errorConflictResponse:       errorConflictResponse,
		errorInternalServerResponse: errorInternalServerResponse,
	}
}

func (receiver RegisterRoute) Method() string {
	return "POST"
}

func (receiver RegisterRoute) Pattern() string {
	return "/register"
}

func (receiver RegisterRoute) Handle(context *gin.Context) {
	request, err := receiver.registerRequest.Make(context)

	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			receiver.errorBadRequestResponse.Make(err),
		)

		return
	}

	user := receiver.userRepository.GetUserByEmail(request.GetBody().GetEmail())

	if user != nil {
		context.JSON(
			http.StatusConflict,
			errors.New("user already exists with this email"),
		)

		return
	}

	userUUID := receiver.userService.GenerateUUID()

	hashedPassword, err := receiver.userService.HashPassword(request.GetBody().GetPassword())

	if err != nil {
		context.JSON(
			http.StatusConflict,
			receiver.errorConflictResponse.Make(err),
		)

		return
	}

	err = receiver.userService.CreateUserByUUIDAndEmailAndHashedPassword(
		userUUID,
		request.GetBody().GetEmail(),
		hashedPassword,
	)

	if err != nil {
		context.JSON(
			http.StatusConflict,
			receiver.errorConflictResponse.Make(err),
		)

		return
	}

	device, err := receiver.deviceService.GetNewDeviceByUserUUIDAndIpAndUserAgent(
		userUUID,
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

	response, err := receiver.successRegisterResponse.Make(device)

	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			receiver.errorInternalServerResponse.Make(err),
		)

		return
	}

	context.JSON(http.StatusCreated, response)
}
