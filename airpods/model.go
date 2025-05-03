package airpods

import "fmt"

type DeviceModel uint16

const (
	ModelAirPods1        DeviceModel = 0x0220
	ModelAirPods2        DeviceModel = 0x0f20
	ModelAirPods3        DeviceModel = 0x1320
	ModelAirPods4        DeviceModel = 0x1920
	ModelAirPods4ANC     DeviceModel = 0x1b20
	ModelAirPodsPro      DeviceModel = 0x0e20
	ModelAirPodsPro2     DeviceModel = 0x1420
	ModelAirPodsPro2USBC DeviceModel = 0x2420
	ModelAirPodsMax      DeviceModel = 0x0a20
	ModelAirPodsMaxUSBC  DeviceModel = 0x1f20
)

func (dm DeviceModel) String() string {
	switch dm {
	case ModelAirPods1:
		return "AirPods 1"
	case ModelAirPods2:
		return "AirPods 2"
	case ModelAirPods3:
		return "AirPods 3"
	case ModelAirPods4:
		return "AirPods 4"
	case ModelAirPods4ANC:
		return "AirPods 4 (ANC)"
	case ModelAirPodsPro:
		return "AirPods Pro"
	case ModelAirPodsPro2:
		return "AirPods Pro 2"
	case ModelAirPodsPro2USBC:
		return "AirPods Pro 2 (USB-C)"
	case ModelAirPodsMax:
		return "AirPods Max"
	case ModelAirPodsMaxUSBC:
		return "AirPods Max (USB-C)"
	default:
		return fmt.Sprintf("Unknown DeviceModel: %04x", uint16(dm))
	}
}
