package natureremo

import (
	"context"

	"github.com/pkg/errors"
	"github.com/tenntenn/natureremo"
)

// Client wraps Nature Remo API Client
type Client struct {
	client *natureremo.Client
}

// RoomMetrics represents a set of room temperature, humidity and illuminance
type RoomMetrics struct {
	Temperature  float64
	Humidity     float64
	Illumination float64
}

// NewClient creates Nature Remo API Client wrapper
func NewClient(accessToken string) *Client {
	return &Client{
		client: natureremo.NewClient(accessToken),
	}
}

// GetRoomMetrics returns available rooms temperature, humidity and illuminance
func (c *Client) GetRoomMetrics(ctx context.Context) (map[string]*RoomMetrics, error) {
	devices, err := c.client.DeviceService.Devices(ctx)
	if err != nil {
		return map[string]*RoomMetrics{}, errors.Wrap(err, "cannot get devices")
	}

	metrics := map[string]*RoomMetrics{}

	for _, d := range devices {
		metrics[d.ID] = &RoomMetrics{
			Temperature:  d.NewestEvents[natureremo.SensorTypeTemperature].Value,
			Humidity:     d.NewestEvents[natureremo.SensorTypeHumidity].Value,
			Illumination: d.NewestEvents[natureremo.SensortypeIllumination].Value,
		}
	}

	return metrics, nil
}
