package expr

import (
	"math"
)

func toRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func toDegrees(radians float64) float64 {
	return radians * 180 / math.Pi
}
