package F1Telemetry

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"

	"github.com/jwjames83/f1telemetry-go/internal/pkg/f1packet"
)

type Receiver struct {
	listening bool
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func New() *Receiver {
	rv := new(Receiver)
	rv.listening = false

	return rv
}

func memsetRepeat(a []byte, v byte) {
	if len(a) == 0 {
		return
	}
	a[0] = v
	for bp := 1; bp < len(a); bp *= 2 {
		copy(a[bp:], a[:bp])
	}
}
/** Flow --
initialize packets and memory

*/
func (r *Receiver) Start(port int) {
	var err error
	var header f1packet.Header
	con, err := net.ListenUDP("udp4", &net.UDPAddr{
		IP:   []byte{127, 0, 0, 1},
		Port: port,
		Zone: "",
	})
	check(err)

	var buffer []byte
	buffer = make([]byte, 2048)
	r.listening = true

	var telemetry f1packet.CarTelemetry
	var participants f1packet.Participants
	var lapData f1packet.LapData
	var motion f1packet.Motion
	var session f1packet.Session
	var event f1packet.Event
	var setups f1packet.CarSetups
	var carStatus f1packet.CarStatus
	var results f1packet.Results
	var lobbyInfo f1packet.LobbyInfo

	var frame uint32 = 0
	var reader = bytes.NewReader(buffer)

	for {
		// Reset the reader
		_, err = reader.Seek(0, 0)
		memsetRepeat (buffer, 0)
		check(err)

		// Read
		n, _, err := con.ReadFromUDP(buffer)
		check(err)

		// Abort if we don't even have the header size
		if n < f1packet.HeaderSize {
			continue
		}

		// Populate the header
		check(binary.Read(reader, binary.LittleEndian, &header))

		// Check if we've got a new frame
		if header.FrameIdentifier > frame {
			// TODO: send stream data via gRPC
			frame = header.FrameIdentifier
		}

		// Check the size
		if n < f1packet.GetSize(header.PacketId) {
			continue
		}

		// Bump the reader past the header
		_, err = reader.Seek(int64(f1packet.HeaderSize+1), 0)

		// Read based on packet type
		switch header.PacketId {
		case f1packet.IdMotion:
			check(binary.Read(reader, binary.LittleEndian, &motion))

		case f1packet.IdSession:
			check(binary.Read(reader, binary.LittleEndian, &session))

		case f1packet.IdLapData:
			check(binary.Read(reader, binary.LittleEndian, &lapData))

		case f1packet.IdCarTelemetry:
			check(binary.Read(reader, binary.LittleEndian, &telemetry))

		case f1packet.IdParticipants:
			check(binary.Read(reader, binary.LittleEndian, &participants))

		// TODO: Properly parse the event info
		case f1packet.IdEvent:
			readIt := false
			check(binary.Read(reader, binary.LittleEndian, &event.StringCode))

			switch string(buffer[f1packet.HeaderSize + 1:f1packet.HeaderSize + 1 + f1packet.EventStringCodeLen]) {
			case f1packet.EventStrSessionStarted:
				frame = header.FrameIdentifier
				fmt.Println("Session started")
			case f1packet.EventStrSessionEnded:
				fmt.Println("Session ended")
			case f1packet.EventStrDRSenabled:
				fmt.Println("DRS enabled")
			case f1packet.EventStrDRSdisabled:
				fmt.Println("DRS disabled")
			case f1packet.EventStrChequered:
				fmt.Println("Checkered flag")

			case f1packet.EventStrFastestLap:
				event.Details = new(f1packet.FastestLap)
				readIt = true
			case f1packet.EventStrRetirement:
				event.Details = new(f1packet.Retirement)
				readIt = true
			case f1packet.EventStrTeamMateInPit:
				event.Details = new(f1packet.TeamMateInPits)
				readIt = true
			case f1packet.EventStrRaceWinner:
				event.Details = new(f1packet.RaceWinner)
				readIt = true
			case f1packet.EventStrPenaltyIssued:
				event.Details = new(f1packet.Penalty)
				readIt = true
			case f1packet.EventStrSpeedTrap:
				event.Details = new(f1packet.SpeedTrap)
				readIt = true
			}

			if readIt {
				check(binary.Read(reader, binary.LittleEndian, event.Details))
				switch d := event.Details.(type) {
				case *f1packet.SpeedTrap:
					name := string(participants.PlayerInfo[d.VehicleIdx].Name[:])
					fmt.Printf("SpeedTrap: %s, %f\n", name, d.Speed)
				case *f1packet.FastestLap:
					name := string(participants.PlayerInfo[d.VehicleIdx].Name[:])
					fmt.Printf("FastestLap: %s %f\n", name, d.LapTime)
				case *f1packet.RaceWinner:
					name := string(participants.PlayerInfo[d.VehicleIdx].Name[:])
					fmt.Printf("RaceWinner: %s\n", name)
				case *f1packet.Penalty:
					name := string(participants.PlayerInfo[d.VehicleIdx].Name[:])
					fmt.Printf("Penalty: %s lap[%d]\n", name, d.LapNum)
				}
			}

		case f1packet.IdCarSetups:
			check(binary.Read(reader, binary.LittleEndian, &setups))

		case f1packet.IdCarStatus:
			check(binary.Read(reader, binary.LittleEndian, &carStatus))

		case f1packet.IdClassification:
			check(binary.Read(reader, binary.LittleEndian, &results))

		case f1packet.IdLobbyInfo:
			check(binary.Read(reader, binary.LittleEndian, &lobbyInfo))
		}
	}
}

func (r *Receiver) Stop() {

}
