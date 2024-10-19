package services

type QueryDeviceService interface {
	GenerateUUID() string
}

type MutableDeviceService interface {
}

type DeviceService interface {
	QueryDeviceService
	MutableDeviceService
}
