package F1Telemetry

import (
	"bytes"
	"encoding/binary"
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
	// var event f1packet.Event
	var setups f1packet.CarSetups
	var carStatus f1packet.CarStatus
	var results f1packet.Results
	var lobbyInfo f1packet.LobbyInfo

	var frame uint32 = 0
	var reader = bytes.NewReader(buffer)

	for {
		// Reset the reader
		_, err = reader.Seek(0, 0)
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
		// case f1packet.IdEvent:
		// 	check(binary.Read(bytes.NewReader(buffer), binary.LittleEndian, &event))

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
