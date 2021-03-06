type loopPacketRaw struct {
	BarTrend        int8
	PacketType      int8
	NextRecord      uint16
	Barometer       uint16
	InsideTemp      int16
	InsideHumidity  uint8
	OutsideTemp     int16
	WindSpeed       uint8
	TenMinWindAvg   uint8
	WindDir         uint16
	ExtraTemp       [7]byte
	SoilTemp        [4]byte
	LeafTemp        [4]byte
	OutsideHumidity uint8
	ExtraHumidities [7]byte
	RainRate        uint16
	// Code Cut For Brevity
}