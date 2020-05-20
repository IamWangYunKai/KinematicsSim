package model

type Robot struct {
	X float64
	Y float64
	Theta float64
	V float64
	W float64
}

type Action struct {
	V float64
	W float64
}