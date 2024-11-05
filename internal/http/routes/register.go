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

	userExists := receiver.userRepository.HasUserByEmail(request.Body.Email)

	if userExists {
		context.JSON(http.StatusConflict, responses.NewErrorConflictResponse(errors.New("user already exists with this email")))

		return
	}

	userUUID := receiver.userService.GenerateUUID()

	err := receiver.userService.CreateUserByUUIDAndEmailAndPassword(userUUID, request.Body.Email, request.Body.Password)

	if err != nil {
		context.JSON(http.StatusConflict, responses.NewErrorConflictResponse(err))

		return
	}

	device, err := receiver.deviceService.GetNewDeviceByUserUUIDAndIpAndUserAgent(
		receiver.config,
		userUUID,
		request.IP,
		request.UserAgent,
	)

	if err != nil {
		context.JSON(http.StatusInternalServerError, responses.NewErrorInternalServerResponse(err))

		return
	}

	response, err := responses.NewSuccessRegisterResponse(receiver.config, device)

	if err != nil {
		context.JSON(http.StatusInternalServerError, responses.NewErrorInternalServerResponse(err))

		return
	}

	context.JSON(http.StatusCreated, response)
}
