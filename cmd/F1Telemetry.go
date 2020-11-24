package main

import "github.com/jwjames83/f1telemetry-go/pkg/F1Telemetry"

func main() {
	server := F1Telemetry.New()
	server.Start(20777)
}
