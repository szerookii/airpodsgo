package airpods

import "fmt"

// TODO: Add PairingModeData struct and decode function

type PairedModeData struct {
	ProtocolID        uint8
	ProtocolLength    uint8
	PairingMode       PairingMode
	DeviceModel       DeviceModel
	UTP               UTP
	BatteryIndication BatteryIndication
	LidIndication     uint8
	Unknown           uint8
	EncryptedPayload  [0xA]byte
}

func DecodePairedModeData(data []byte) (PairedModeData, error) {
	if len(data) < 0xE {
		return PairedModeData{}, fmt.Errorf("data too short")
	}

	if data[0] != ProtocolID {
		return PairedModeData{}, fmt.Errorf("invalid protocol ID")
	}

	if data[2] != 0x01 {
		return PairedModeData{}, fmt.Errorf("invalid pairing mode")
	}

	if data[1] != 0x19 {
		return PairedModeData{}, fmt.Errorf("invalid protocol length")
	}

	appData := PairedModeData{
		ProtocolID:        data[0],
		ProtocolLength:    data[1],
		PairingMode:       PairingMode(data[2]),
		DeviceModel:       DeviceModel(data[3])<<8 | DeviceModel(data[4]),
		UTP:               UTP(data[5]),
		BatteryIndication: BatteryIndication(data[6])<<8 | BatteryIndication(data[7]),
		LidIndication:     data[8],
		Unknown:           data[9],
	}

	copy(appData.EncryptedPayload[:], data[0xA:0xA+0xA])

	return appData, nil
}
