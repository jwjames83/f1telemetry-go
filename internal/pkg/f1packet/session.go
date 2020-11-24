package f1packet

/**
Session Packet
The session packet includes details about the current session in progress.

Frequency: 2 per second
Size: 251 bytes (Packet size updated in Beta 3)
Version: 1
*/

const SessionPacketSize = 251

type marshalZone struct {
	ZoneStart float // Fraction (0..1) of way through the lap the marshal zone starts
	ZoneFlag  int8  // -1 = invalid/unknown, 0 = none, 1 = green, 2 = blue, 3 = yellow, 4 = red
}

type weatherForecastSample struct {
	SessionType uint8 // 0 = unknown, 1 = P1, 2 = P2, 3 = P3, 4 = Short P, 5 = Q1 6 = Q2, 7 = Q3, 8 = Short Q, 9 = OSQ, 10 = R, 11 = R2 12 = Time Trial
	TimeOffset  uint8 // Time in minutes the forecast is for
	Weather     uint8 // Weather - 0 = clear, 1 = light cloud, 2 = overcast 3 = light rain, 4 = heavy rain, 5 = storm
	TrackTemp   int8  // Track temp. in degrees celsius
	AirTemp     int8  // Air temp. in degrees celsius
}

type Session struct {
	Weather                   uint8                     // Weather - 0 = clear, 1 = light cloud, 2 = overcast 3 = light rain, 4 = heavy rain, 5 = storm
	TrackTemp                 int8                      // Track temp. in degrees celsius
	AirTemp                   int8                      // Air temp. in degrees celsius
	TotalLaps                 uint8                     // Total number of Laps in this race
	TrackLength               uint16                    // Track length in metres
	Sessiontype               uint8                     // 0 = unknown, 1 = P1, 2 = P2, 3 = P3, 4 = Short P  5 = Q1, 6 = Q2, 7 = Q3, 8 = Short Q, 9 = OSQ 10 = R, 11 = R2, 12 = Time Trial
	Trackid                   int8                      // -1 for unknown, 0-21 for tracks, see appendix
	Formula                   uint8                     // Formula, 0 = F1 Modern, 1 = F1 Classic, 2 = F2, 3 = F1 Generic
	SessionTimeleft           uint16                    // Time left in session in seconds
	SessionDuration           uint16                    // Session duration in seconds
	PitSpeedLimit             uint8                     // Pit speed limit in kilometres per hour
	GamePaused                uint8                     // Whether the game is paused
	IsSpectating              uint8                     // Whether the player is spectating
	SpectatorCarIndex         uint8                     // Index of the car being spectated
	SLIProNativeSupport       uint8                     // SLI Pro support, 0 = inactive, 1 = active
	NumMarshalZones           uint8                     // Number of marshal zones to follow
	MarshalZones              [21]marshalZone           // List of marshal zones – max 21
	SafetyCarStatus           uint8                     // 0 = no safety car, 1 = full safety car 2 = virtual safety car
	NetworkGame               uint8                     // 0 = offline, 1 = online
	NumWeatherForecastSamples uint8                     // Number of weather samples to follow
	WeatherForecastSamples    [20]weatherForecastSample // Array of weather forecast samples
}
