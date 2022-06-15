package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Interface interface {
	AvailableDevices() ([]Device, error)
	PlayTrack() error
}

type spotify struct {
	token string
}

type Config struct {
	Token string
}

type Device struct {
	Id               string `json:"id"`
	IsActive         bool   `json:"is_active"`
	IsPrivacySession bool   `json:"is_private_session"`
	IsRestricted     bool   `json:"is_restricted"`
	Name             string `json:"name"`
	Type             string `json:"type"`
	VolumePercent    int    `json:"volume_percent"`
}

type Devices struct {
	Devices []Device `json:"devices"`
}

const (
	devicesToken  = "BQD-QxlFjUaaJ1JxrorqrOBPuHTFuzCaK-c6euwSJ2uOEig-7ys-DafwDDif8FUsnNQfW_dOw4xUPYpIELKX3BlUik3aJekZ6JxnS6qrTVbs1BdnXfyOT_I1G95xgmZvIO6z3_hM0WOJtqjCbpSps-K_-OIUJDhrexdTKyNKzvPcCfVyhYs9jY9FPw"
	playbackToken = "BQD-QxlFjUaaJ1JxrorqrOBPuHTFuzCaK-c6euwSJ2uOEig-7ys-DafwDDif8FUsnNQfW_dOw4xUPYpIELKX3BlUik3aJekZ6JxnS6qrTVbs1BdnXfyOT_I1G95xgmZvIO6z3_hM0WOJtqjCbpSps-K_-OIUJDhrexdTKyNKzvPcCfVyhYs9jY9FPw"

	devicesURL  = "https://api.spotify.com/v1/me/player/devices"
	playbackURL = "https://api.spotify.com/v1/me/player/play?device_id=4e7cec290b3ae45cd054662aa35b19da2d686cb2"

	playbackBody = "{\"context_uri\": \"spotify:album:3MXU6UoWrf4w4bOvjZTlvY\", \"offset\": {\"position\": 2}, \"position_ms\": 12650}"
)

func New(conf *Config) (Interface, error) {
	return &spotify{}, nil
}

func (s *spotify) AvailableDevices() ([]Device, error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", devicesURL, nil)
	if err != nil {
		//Handle Error
	}

	req.Header = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {"Bearer " + playbackToken},
	}

	res, err := client.Do(req)
	if err != nil {
		//Handle Error
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	fmt.Printf("body: %+v", string(b))

	devices := Devices{}
	err = json.Unmarshal(b, &devices)

	return devices.Devices, err
}

func (s *spotify) PlayTrack() error {
	client := http.Client{}
	req, err := http.NewRequest("PUT", playbackURL, bytes.NewBuffer([]byte(playbackBody)))
	if err != nil {
		//Handle Error
	}

	req.Header = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {"Bearer " + playbackToken},
	}

	_, err = client.Do(req)
	fmt.Printf("err: %+v\n", err)
	return err
}
