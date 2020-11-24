package f1packet

import (
	"fmt"
	"math"
	"strconv"
)

type float = float32
type double = float64
type char = byte

const HeaderSize int = 23
const MaxDrivers int = 22

type PacketInterface interface {
	Size() int
}

func GetSize(id uint8) int {
	switch id {
	case IdEvent:
		return EventPacketSize

	case IdCarTelemetry:
		return CarTelemetryPacketSize

	case IdLapData:
		return LapPacketSize

	case IdCarSetups:
		return CarSetupsPacketSize

	case IdCarStatus:
		return CarStatusPacketSize

	case IdLobbyInfo:
		return LobbyPacketSize

	case IdParticipants:
		return ParticipantsPacketSize

	case IdClassification:
		return ClassificationPacketSize

	case IdSession:
		return SessionPacketSize

	default:
		fmt.Println("Unknown packet ID: " + strconv.Itoa(int(id)))
		return math.MaxUint32
	}
}

type Header struct {
	PacketFormat     uint16 // 2020
	GameMajorVersion uint8  // Game major version - "X.00"
	GameMinorVersion uint8  // Game minor version - "1.XX"
	PacketVersion    uint8  // Version of this packet type, all start from 1
	PacketId         uint8  // Identifier for the packet type, see below
	SessionUid       uint64 // Unique identifier for the session
	SessionTime      float  // Session timestamp
	FrameIdentifier  uint32 // Identifier for the frame the data was retrieved on
	PlayerCarIndex   uint8  // Index of player's car in the array

	// ADDED IN BETA 2:
	SecondaryPlayerCarIndex uint8 // Index of secondary player's car in the array (splitscreen), 255 if no second player
}

const (
	IdMotion uint8 = iota
	IdSession
	IdLapData
	IdEvent
	IdParticipants
	IdCarSetups
	IdCarTelemetry
	IdCarStatus
	IdClassification
	IdLobbyInfo
)

const IdResult = IdClassification
