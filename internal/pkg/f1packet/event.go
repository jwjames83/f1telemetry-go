package f1packet

/**
Event Packet
This packet gives details of events that happen during the course of a session.

Frequency: When the event occurs
Size: 35 bytes (Packet size updated in Beta 3)
Version: 1
*/

const EventPacketSize int = 35
const EventStringCodeLen int = 4

type FastestLap struct {
	VehicleIdx uint8 // Vehicle index of car achieving fastest lap
	LapTime    float // Lap time is in seconds
}

type Retirement struct {
	VehicleIdx uint8 // Vehicle index of car retiring
}

type TeamMateInPits struct {
	VehicleIdx uint8 // Vehicle index of team mate
}

type RaceWinner struct {
	VehicleIdx uint8 // Vehicle index of the race winner
}

type Penalty struct {
	PenaltyType      uint8 // Penalty type – see Appendices
	InfringementType uint8 // Infringement type – see Appendices
	VehicleIdx       uint8 // Vehicle index of the car the penalty is applied to
	OtherVehicleIdx  uint8 // Vehicle index of the other car involved
	Time             uint8 // Time gained, or time spent doing action in seconds
	LapNum           uint8 // Lap the penalty occurred on
	PlacesGained     uint8 // Number of places gained by this
}

type SpeedTrap struct {
	VehicleIdx uint8 // Vehicle index of the vehicle triggering speed trap
	Speed      float // Top speed achieved in kilometres per hour
}

type Event struct {
	StringCode [4]byte     // Event string code, see below
	Details    interface{} // Event details - should be interpreted differently for each type
}

const (
	EventStrSessionStarted = "SSTA"
	EventStrSessionEnded   = "SEND"            // Sent when the session ends
	EventStrFastestLap     = "FTLP"            // When a driver achieves the fastest lap
	EventStrRetirement     = "RTMT"            // When a driver retires
	EventStrDRSenabled     = "DRSE"            // Race control have enabled DRS
	EventStrDRSdisabled    = "DRSD"            // Race control have disabled DRS
	EventStrTeamMateInPit  = "TMPT"            // Your team mate has entered the pits
	EventStrChequered      = "CHQF"            // The chequered flag has been waved
	EventStrCheckered      = EventStrChequered // Alternative spelling of chequred
	EventStrRaceWinner     = "RCWN"            // The race winner is announced
	EventStrPenaltyIssued  = "PENA"            // A penalty has been issued – details in event
	EventStrSpeedTrap      = "SPTP"            // Speed trap has been triggered by fastest speed
)
