package airpods

import "fmt"

type BatteryLevel uint8

func (b BatteryLevel) String() string {
	switch b {
	case 0x00:
		return "0%"
	case 0x01:
		return "10%"
	case 0x02:
		return "20%"
	case 0x03:
		return "30%"
	case 0x04:
		return "40%"
	case 0x05:
		return "50%"
	case 0x06:
		return "60%"
	case 0x07:
		return "70%"
	case 0x08:
		return "80%"
	case 0x09:
		return "90%"
	case 0x0A:
		return "100%"
	case 0x0F:
		return "Disconnected"
	case 0x10:
		return "Unknown"
	case 0x11:
		return "Not Supported"
	default:
		return fmt.Sprintf("Unknown battery level: %02x", uint8(b))
	}
}

type BatteryIndication uint16

func (b BatteryIndication) CaseBatteryLevel() BatteryLevel {
	return BatteryLevel(b & 0x0F)
}

func (b BatteryIndication) LeftLevel() BatteryLevel {
	return BatteryLevel((b >> 0xC) & 0x0F)
}

func (b BatteryIndication) RightLevel() BatteryLevel {
	return BatteryLevel((b >> 0x8) & 0x0F)
}

func (b BatteryIndication) LeftCharging() bool {
	return (b>>0x4)&0x01 == 1
}

func (b BatteryIndication) RightCharging() bool {
	return (b>>0x5)&0x01 == 1
}

func (b BatteryIndication) CaseCharging() bool {
	return (b>>0x6)&0x01 == 1
}

func (b BatteryIndication) Unknown() bool {
	return (b>>0x7)&0x01 == 1
}
