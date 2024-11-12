package routes

import (
	"errors"
	"github.com/fromsi/jwt-oauth-sso/internal/configs"
	"github.com/fromsi/jwt-oauth-sso/internal/http/requests"
	"github.com/fromsi/jwt-oauth-sso/internal/http/responses"
	"github.com/fromsi/jwt-oauth-sso/internal/repositories"
	"github.com/fromsi/jwt-oauth-sso/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegisterRoute struct {
	config           configs.Config
	userRepository   repositories.UserRepository
	deviceRepository repositories.DeviceRepository
	userService      services.UserService
	deviceService    services.DeviceService
}

func NewRegisterRoute(
	config configs.Config,
	userRepository repositories.UserRepository,
	deviceRepository repositories.DeviceRepository,
	userService services.UserService,
	deviceService services.DeviceService,
) *RegisterRoute {
	return &RegisterRoute{
		config:           config,
		userRepository:   userRepository,
		deviceRepository: deviceRepository,
		userService:      userService,
		deviceService:    deviceService,
	}
}

func (receiver RegisterRoute) Method() string {
	return "POST"
}

func (receiver RegisterRoute) Pattern() string {
	return "/register"
}

func (receiver RegisterRoute) Handle(context *gin.Context) {
	request, errResponse := requests.NewRegisterRequest(context)

	if errResponse != nil {
		context.JSON(http.StatusBadRequest, errResponse)

		return
	}

	user := receiver.userRepository.GetUserByEmail(request.Body.Email)

	if user != nil {
		err := responses.NewErrorConflictResponse(
			errors.New("user already exists with this email"),
		)

		context.JSON(http.StatusConflict, err)

		return
	}

	userUUID := receiver.userService.GenerateUUID()

	hashedPassword, err := receiver.userService.HashPassword(request.Body.Password)

	if err != nil {
		context.JSON(http.StatusConflict, responses.NewErrorConflictResponse(err))

		return
	}

	err = receiver.userService.CreateUserByUUIDAndEmailAndHashedPassword(
		userUUID,
		request.Body.Email,
		hashedPassword,
	)

	if err != nil {
		context.JSON(http.StatusConflict, responses.NewErrorConflictResponse(err))

		return
	}

	device := receiver.deviceService.GetNewDeviceByUserUUIDAndIpAndUserAgent(
		userUUID,
		request.IP,
		request.UserAgent,
	)

	err = receiver.deviceRepository.CreateDevice(device)

	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			responses.NewErrorInternalServerResponse(err),
		)

		return
	}

	device = receiver.deviceService.GetNewRefreshDetailsByDevice(device)

	err = receiver.deviceRepository.UpdateDevice(device)

	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			responses.NewErrorInternalServerResponse(err),
		)

		return
	}

	response, err := responses.NewSuccessRegisterResponse(receiver.config, device)

	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			responses.NewErrorInternalServerResponse(err),
		)

		return
	}

	context.JSON(http.StatusCreated, response)
}
