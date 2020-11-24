package f1

import "github.com/jwjames83/f1telemetry-go/internal/pkg/f1/team"

type driverInfo struct {
	Name string
	Team team.Info
}

type Drivers struct {
	Total int
	Info  []driverInfo
}

func InitDrivers() *Drivers {
	rv := new(Drivers)
	// TODO: Read f1-drivers.csv and populate Drivers struct
	// rv.Total =
	// rv.Info = make([]driverInfo, rv.Total)
	return rv
}
