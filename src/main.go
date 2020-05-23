package main

import (
	"./dynamics"
	"./model"
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"sync"
)

var done = make(chan struct{})

var robotMap = make(map[string]model.Robot)
var actionMap = make(map[string]model.Action)

var obstacleMap = make(map[string]model.Obstacle)

type Message struct {
	Mtype string
	Pri int
	Id string
	Data string
}

func addRobot(id string, x float64, y float64, theta float64){
	robotMap[id] = model.Robot{
		X: x,
		Y: y,
		Theta: theta,
		V: 0.0,
		W: 0.0,
	}
	actionMap[id] = model.Action{
		V: 0.0,
		W: 0.0,
	}
}

func addObstacle(id string, x float64, y float64, r float64){
	obstacleMap[id] = model.Obstacle{
		X: x,
		Y: y,
		R: r,
	}
}

func addRandomObstacles(num int){
	var mutex sync.Mutex
	var wg sync.WaitGroup
	wg.Add(num)
	for i:=0; i<num; i++{
		go func() {
			defer wg.Done()
			id := string(rand.Intn(999999))
			x := 2 * model.MAX_X * rand.Float64() - model.MAX_X
			y := 2 * model.MAX_Y * rand.Float64() - model.MAX_Y
			r := 5 * rand.Float64()
			mutex.Lock()
			addObstacle(id, x, y, r)
			mutex.Unlock()
		}()
	}
	wg.Wait()
}

func read(socket *net.UDPConn){
	var mutex sync.Mutex
	for {
		data := make([]byte, 65535)
		n, remoteAddr, err := socket.ReadFromUDP(data)
		if err != nil {
			fmt.Printf("error during read: %s", err)
			continue
		}
		var message Message
		if err := json.Unmarshal(data[:n], &message); err != nil {
			fmt.Println("Json Unmarshal ERROR: ", err, data[:n])
			continue
		}
		if message.Mtype == "register" {
			fmt.Printf("Register <%s> %s\n", message.Id, message.Data)
			feedback := Message{
				Mtype: "register",
				Pri: 5,
				Id: "000000",
				Data: remoteAddr.String(),
			}
			feedbackStr, err := json.Marshal(feedback)
			if err != nil {
				fmt.Println(err)
				continue
			}
			socket.WriteToUDP(feedbackStr, remoteAddr)
			mutex.Lock()
			addRobot(message.Id, 0.0, 0.0, 0.0)
			mutex.Unlock()
		} else if message.Mtype == "cmd" {
			if _, ok := robotMap[message.Id]; ok {
				var action model.Action
				if err := json.Unmarshal([]byte(message.Data), &action); err != nil {
					fmt.Println("Json Unmarshal ERROR in message.Data: ", err, data[:n])
					continue
				}
				mutex.Lock()
				actionMap[message.Id] = action
				mutex.Unlock()
			} else {
				fmt.Println("No robot ID:", message.Id)
			}
		} else if message.Mtype == "step" {
			dynamics.Step(robotMap, actionMap, obstacleMap)
			dynamics.ShowInfo(robotMap)
		} else {
			fmt.Println("Error message", message)
		}
	}
}

func main() {
	addRandomObstacles(10)
	clientAddr := &net.UDPAddr{IP: net.IPv4zero, Port: model.LISTEN_PORT}
	clientListener, err := net.ListenUDP("udp", clientAddr)
	if err != nil {
		fmt.Println(err)
	}
	go read(clientListener)
	<-done
}