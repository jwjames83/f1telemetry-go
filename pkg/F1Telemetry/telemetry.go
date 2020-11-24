package F1Telemetry

import (
	"bytes"
	"encoding/binary"
	"net"
	"strconv"

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

/** Flow --
initialize packets and memory

*/
func (r *Receiver) Start(port int) {
	var err error
	var header f1packet.Header
	con, err := net.ListenPacket("udp4", ":"+strconv.Itoa(port))
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

	for {
		n, _, err := con.ReadFrom(buffer)
		check(err)

		// Abort if we don't even have the header size
		if n < f1packet.HeaderSize {
			continue
		}

		err = binary.Read(bytes.NewReader(buffer[:f1packet.HeaderSize+1]), binary.LittleEndian, &header)
		check(err)

		// Scoot up
		buffer = buffer[f1packet.HeaderSize+1:]

		// GetSize will return math.MaxUint32 if the packet ID is unknown
		if n < f1packet.GetSize(header.PacketId) {
			continue
		}

		switch header.PacketId {
		case f1packet.IdMotion:
			check(binary.Read(bytes.NewReader(buffer), binary.LittleEndian, &motion))

		case f1packet.IdSession:
			check(binary.Read(bytes.NewReader(buffer), binary.LittleEndian, &session))

		case f1packet.IdLapData:
			check(binary.Read(bytes.NewReader(buffer), binary.LittleEndian, &lapData))

		case f1packet.IdCarTelemetry:
			check(binary.Read(bytes.NewReader(buffer), binary.LittleEndian, &telemetry))

		case f1packet.IdParticipants:
			check(binary.Read(bytes.NewReader(buffer), binary.LittleEndian, &participants))

		case f1packet.IdEvent:
			check(binary.Read(bytes.NewReader(buffer), binary.LittleEndian, &event))

		case f1packet.IdCarSetups:
			check(binary.Read(bytes.NewReader(buffer), binary.LittleEndian, &setups))

		case f1packet.IdCarStatus:
			check(binary.Read(bytes.NewReader(buffer), binary.LittleEndian, &carStatus))

		case f1packet.IdClassification:
			check(binary.Read(bytes.NewReader(buffer), binary.LittleEndian, &results))

		case f1packet.IdLobbyInfo:
			check(binary.Read(bytes.NewReader(buffer), binary.LittleEndian, &lobbyInfo))
		}
	}
}

func (r *Receiver) Stop() {

}
