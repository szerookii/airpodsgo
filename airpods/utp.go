package airpods

type UTP uint8

func (utp UTP) OneOrBothInEar() bool {
	return (utp>>0x1)&0x01 == 1
}

func (utp UTP) BothInCase() bool {
	return (utp>>0x2)&0x01 == 1
}

func (utp UTP) BothInEar() bool {
	return (utp>>0x3)&0x01 == 1
}

func (utp UTP) OneOrBothInCase() bool {
	return (utp>>0x4)&0x01 == 1
}
