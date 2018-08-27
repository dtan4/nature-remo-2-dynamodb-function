package natureremo

import (
	"context"

	"github.com/pkg/errors"
	"github.com/tenntenn/natureremo"

	"github.com/dtan4/nature-remo-2-dynamodb-function/types"
)

// Client wraps Nature Remo API Client
type Client struct {
	client *natureremo.Client
}

// NewClient creates Nature Remo API Client wrapper
func NewClient(accessToken string) *Client {
	return &Client{
		client: natureremo.NewClient(accessToken),
	}
}

// GetRoomMetrics returns available rooms temperature, humidity and illuminance
func (c *Client) GetRoomMetrics(ctx context.Context) (map[string]*types.RoomMetrics, error) {
	devices, err := c.client.DeviceService.Devices(ctx)
	if err != nil {
		return map[string]*types.RoomMetrics{}, errors.Wrap(err, "cannot get devices")
	}

	metrics := map[string]*types.RoomMetrics{}

	for _, d := range devices {
		metrics[d.ID] = &types.RoomMetrics{
			Temperature:  d.NewestEvents[natureremo.SensorTypeTemperature].Value,
			Humidity:     d.NewestEvents[natureremo.SensorTypeHumidity].Value,
			Illumination: d.NewestEvents[natureremo.SensortypeIllumination].Value,
			// Nature Remo mini cannot monitor humidity and illumination
			CreatedAt: d.NewestEvents[natureremo.SensorTypeTemperature].CreatedAt,
		}
	}

	return metrics, nil
}
