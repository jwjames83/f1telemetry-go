package f1packet

/**
This is a list of participants in the race. If the vehicle is controlled by AI,
then the name will be the driver name. If this is a multiplayer game, the names
will be the Steam Id on PC, or the LAN name if appropriate.

N.B. on Xbox One, the names will always be the driver name, on PS4 the name will
be the LAN name if playing a LAN game, otherwise it will be the driver name.

The array should be indexed by vehicle index.

Frequency: Every 5 seconds
Size: 1213 bytes (Packet size updated in Beta 3)
Version: 1
*/

const ParticipantsPacketSize int = 1213

type participantDetails struct {
	AiControlled  uint8    // Whether the vehicle is AI (1) or Human (0) controlled
	DriverId      uint8    // Driver id - see appendix
	TeamId        uint8    // Team id - see appendix
	RaceNumber    uint8    // Race number of the car
	Nationality   uint8    // Nationality of the driver
	Name          [48]byte // Name of participant in UTF-8 format – null terminated (Will be truncated with … (U+2026) if too long)
	YourTelemetry uint8    // The player's UDP setting, 0 = restricted, 1 = public
}

type Participants struct {
	NumActiveCars uint8 // Number of active cars in the data – should match number of cars on HUD
	PlayerInfo    [MaxDrivers]participantDetails
}
