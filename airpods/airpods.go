package airpods

const (
	AppleCompanyID = uint16(0x004C)
	ProtocolID     = uint8(0x07)
)

type PairingMode uint8

const (
	PairingModePairing PairingMode = 0x00
	PairingModePaired  PairingMode = 0x01
)
