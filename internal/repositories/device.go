package repositories

type QueryDeviceRepository interface {
	GetDevicesByUserUUID(string) []Device
}

type MutableDeviceRepository interface {
	CreateDevice(Device) error
	UpdateDevice(Device) error
	DeleteDeviceByUUID(string) error
	DeleteAllDevicesByUserUUID(string) error
}

type DeviceRepository interface {
	QueryDeviceRepository
	MutableDeviceRepository
}

type QueryDevice interface {
	GetUUID() string
	GetUserUUID() string
	GetAgent() string
	GetIp() string
	GetExpiredAt() int
	GetCreatedAt() int
	GetUpdatedAt() int
}

type MutableDevice interface {
	SetUUID(string)
	SetUserUUID(string)
	SetAgent(string)
	SetIp(string)
	SetExpiredAt(int)
	SetCreatedAt(int)
	SetUpdatedAt(int)
}

type Device interface {
	QueryDevice
	MutableDevice
}
