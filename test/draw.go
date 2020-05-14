package main

// 需要导入 syscall/js 包以调用 js API
import (
	"math/rand"
	"syscall/js"
	"time"
)

const (
	width = 400
	height = 400
)

// 生成 0 - 1 的随机数
func getRandomNum() float32 {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	n := float32(rand.Intn(10000))
	return  n / 10000.0
}

// 使用 canvas 绘制随机图
func draw() {
	var canvas js.Value = js.
		Global().
		Get("document").
		Call("getElementById", "canvas")

	var context js.Value = canvas.Call("getContext", "2d")

	// reset
	canvas.Set("height", height)
	canvas.Set("width", width)
	context.Call("clearRect", 0, 0, width, height)

        // 随机绘制 50 条直线
	for i := 0; i < 50; i ++ {
		context.Call("beginPath")
		context.Call("moveTo", getRandomNum() * width, getRandomNum() * height)
		context.Call("lineTo", getRandomNum() * width, getRandomNum() * height)
		context.Call("stroke")
	}
}

// 主程序入口
func main() {
	println("wasm app works")
	// bootstrap app
	draw()
}
