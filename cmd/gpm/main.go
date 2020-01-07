package main

import (
	"fmt"
	"time"

	gopor "github.com/krsanky/go-portmidi1"
	"github.com/rakyll/portmidi"
)

func main() {
	var dev_id portmidi.DeviceID = 7

	fmt.Printf("go portmidi 1\n")
	gopor.GP1()
	portmidi.Initialize()
	out, err := portmidi.NewOutputStream(dev_id, 1024, 0)
	if err != nil {
		panic(err)
	}

	// note on events to play C major chord
	out.WriteShort(0x90, 60, 100)
	out.WriteShort(0x90, 64, 100)
	out.WriteShort(0x90, 67, 100)

	// notes will be sustained for 2 seconds
	time.Sleep(2 * time.Second)

	// note off events
	out.WriteShort(0x80, 60, 100)
	out.WriteShort(0x80, 64, 100)
	out.WriteShort(0x80, 67, 100)

	out.Close()

	portmidi.Terminate()
}
