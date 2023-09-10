package enum

type DesiredPlatform uint8

const (
	ShipmentIntel DesiredPlatform = iota
	FreightIntel
	Others
)

func (d DesiredPlatform) String() string {
	switch d {
	case ShipmentIntel:
		return "shipmentIntel"
	case FreightIntel:
		return "freightIntel"
	default:
		return "others"
	}
}
