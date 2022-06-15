package main

import (
	"fmt"
	"time"
)

const (
	silverStreak = "Tesla Silver Streak"
)

func main() {
	s, _ := New(&Config{})

	minuteTicker := time.NewTicker(time.Second * 10)
	secondTicker := time.NewTicker(time.Second * 2)
	carIsPlaying := false

	for {
		select {
		case <-secondTicker.C:
			if carIsPlaying {
				fmt.Println("playing track")
				s.PlayTrack()
			}
		case <-minuteTicker.C:
			devices, err := s.AvailableDevices()
			if err != nil {
				fmt.Println(err)
			}

			carIsPlaying = false

			for _, device := range devices {
				if device.IsActive && device.Name == silverStreak {
					fmt.Println("silver streak is online")
					carIsPlaying = true
				}
			}

			fmt.Printf("car is playing: %+v\n", carIsPlaying)
		}
	}
}
