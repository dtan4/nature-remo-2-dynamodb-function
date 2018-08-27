package types

import (
	"time"
)

// RoomMetrics represents a set of room temperature, humidity and illuminance
type RoomMetrics struct {
	Temperature  float64
	Humidity     float64
	Illumination float64
	CreatedAt    time.Time
}
