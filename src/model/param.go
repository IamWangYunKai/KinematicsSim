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
	ROBOT_SIZE float64 = 3
	LASER_NUM int = 1024
	DELTA_ANGLE = 2.0 * math.Pi / float64(LASER_NUM)
	INFINITE float64 = 99999.9
	LISTEN_PORT int = 10006
)