package main

import (
	"./dynamics"
	"./model"
	"fmt"
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
	fmt.Println("Hello World")

	id := strconv.Itoa(rand.Intn(10000))
	addRobot(id, 0.0, 0.0, 0.0)
	for i := 0; i < 100; i++ {
		dynamics.Step(robotMap, actionMap)
	}
	dynamics.ShowInfo(robotMap)
}