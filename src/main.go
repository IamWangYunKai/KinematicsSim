package main

import (
	"./dynamics"
	"./model"
	"encoding/json"
	"fmt"
	"sync"
	"net"
)

var done = make(chan struct{})

var robotMap = make(map[string]model.Robot)
var actionMap = make(map[string]model.Action)

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
			dynamics.Step(robotMap, actionMap)
			dynamics.ShowInfo(robotMap)
		} else {
			fmt.Println("Error message", message)
		}
	}
}

func main() {
	clientAddr := &net.UDPAddr{IP: net.IPv4zero, Port: model.LISTEN_PORT}
	clientListener, err := net.ListenUDP("udp", clientAddr)
	if err != nil {
		fmt.Println(err)
	}
	go read(clientListener)
	<-done
}