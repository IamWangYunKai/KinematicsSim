package model

import "math"

const (
	MAX_X float64 = 100.
	MAX_Y float64 = 100.
	MAX_V float64 = 4.0
	MAX_W float64 = 2*math.Pi
	MAX_ACC float64 = 4.0
	MAX_ALPHA float64 = 2*math.Pi
	DT float64 = 1.0/100.
	MAX_STEP_V = MAX_ACC*DT
	MAX_STEP_W = MAX_ALPHA*DT
)