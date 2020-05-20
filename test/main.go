package main

import (
	"./dynamics"
	"./model"
	"math/rand"
	"strconv"
)

var robotMap = make(map[string]model.Robot)
var actionMap = make(map[string]model.Action)

func addRobot(id string, x float64, y float64, theta float64){
	robotMap[id] = model.Robot{
		X: x,
		Y: y,
		Theta: theta,
		V: 0.0,
		W: 0.0,
	}
	actionMap[id] = model.Action{
		V: 10.0,
		W: 0.0,
	}
}

func main() {

	for i := 0; i < 5; i++ {
		id := strconv.Itoa(rand.Intn(10000))
		addRobot(id, 0.0, 0.0, 0.0)
	}

	for i := 0; i < 50; i++ {
		dynamics.Step(robotMap, actionMap)
	}
	dynamics.ShowInfo(robotMap)
}