package f1packet

/**
Event Packet
This packet gives details of events that happen during the course of a session.

Frequency: When the event occurs
Size: 35 bytes (Packet size updated in Beta 3)
Version: 1
*/

const EventPacketSize int = 35

type EventDetails interface {
	Visit(Visitor)
}

// The handlers for each event type are defined in instances of this struct
type Visitor struct {
	VisitFastestLap     func(lap FastestLap)
	VisitRetirement     func(retirement Retirement)
	VisitTeamMateInPits func(pits TeamMateInPits)
	VisitSpeedTrap      func(trap SpeedTrap)
}

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
	StringCode string       // Event string code, see below
	Details    EventDetails // Event details - should be interpreted differently for each type
}

const (
	SessionStartedStr = "SSTA"
	SessionEndedStr   = "SEND"       // Sent when the session ends
	FastestLapStr     = "FTLP"       // When a driver achieves the fastest lap
	RetirementStr     = "RTMT"       // When a driver retires
	DRSenabledStr     = "DRSE"       // Race control have enabled DRS
	DRSdisabledStr    = "DRSD"       // Race control have disabled DRS
	TeamMateInPitStr  = "TMPT"       // Your team mate has entered the pits
	ChequeredStr      = "CHQF"       // The chequered flag has been waved
	CheckeredStr      = ChequeredStr // Alternative spelling of chequred
	RaceWinnerStr     = "RCWN"       // The race winner is announced
	PenaltyIssuedStr  = "PENA"       // A penalty has been issued – details in event
	SpeedTrapTrStr    = "SPTP"       // Speed trap has been triggered by fastest speed
)
