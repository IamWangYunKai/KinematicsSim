package main

import (
	"./dynamics"
	"./model"
	"math/rand"
	"syscall/js"
	"time"
	"strconv"
	"math"
	//"fmt"
)

const (
	width = 600
	height = 600
	rate = 5//2*model.MAX_X/width
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

func getRandomNum() float32 {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	n := float32(rand.Intn(10000))
	return  n / 10000.0
}

func render() {
	var canvas js.Value = js.
		Global().
		Get("document").
		Call("getElementById", "canvas")

	var context js.Value = canvas.Call("getContext", "2d")

	canvas.Set("height", height)
	canvas.Set("width", width)
	context.Call("clearRect", 0, 0, width, height)

	for _, robot := range robotMap {
		context.Set("strokeStyle", "red")
		context.Set("lineWidth", 5)
		context.Call("beginPath")
		context.Call("arc", rate*robot.X*width+width/2, rate*robot.Y*height+height/2, 20, 0, 2*math.Pi)
		context.Call("stroke")
	}
}

func addEventListener()  {
	done := make(chan struct{})

	var cb js.Func
    cb = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		dynamics.Step(robotMap, actionMap)
		dynamics.ShowInfo(robotMap)
        render()
        return nil
    })
	js.Global().Get("document").
		Call("getElementById", "canvas").
		Call("addEventListener", "click", cb)
	<-done
}

func bootstrapApp() {
	render()
	addEventListener()
}

func main() {
	println("wasm app works")
	for i := 0; i < 3; i++ {
		id := strconv.Itoa(rand.Intn(10000))
		addRobot(id, 0.0, 0.0, 0.0)
	}
	// bootstrap app
	bootstrapApp()
}