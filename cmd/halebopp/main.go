package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/fsufitch/halebopp-hotspot"
)

func main() {
	ctx := context.Background()
	ctx, _ = signal.NotifyContext(ctx, os.Kill, os.Interrupt)

	hb, cleanup, err := initializeDefaultHaleBopp()

	if err != nil {
		cleanup()
		panic(err)
	}

loop:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("exit")
			break loop
		case <-time.After(2 * time.Second):
		}

		fmt.Println("-----")
		voltage, level, err := hb.Battery.Stats()
		if err != nil {
			fmt.Printf("battery error: %v\n", err)
		}

		charging, err := hb.ChargeState()
		if err != nil {
			fmt.Printf("charging error: %v\n", err)
		}

		fmt.Printf("Voltage: %f\n", voltage)
		fmt.Printf("Battery: %f\n", level)
		fmt.Printf("Charging: %v\n", charging == halebopp.ChargeState_Charging)
	}

	cleanup()
}
