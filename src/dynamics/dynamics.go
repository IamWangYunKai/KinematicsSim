package dynamics

import (
	"../model"
	"math"
	"sync"
)

func clip(value, minValue, maxValue float64) float64 {
	return math.Min(math.Max(value, minValue), maxValue)
}

func normalize(theta float64) float64 {
	if theta > math.Pi {
		return theta - 2 * math.Pi
	} else if theta < -math.Pi {
		return theta + 2 * math.Pi
	} else {
		return theta
	}
}

func Step(robotMap map[string]model.Robot, actionMap map[string]model.Action){
	var mutex sync.Mutex
	var wg sync.WaitGroup
	wg.Add(len(robotMap))

	var ids []string
	for id, _ := range robotMap{
		ids = append(ids, id)
	}

	for _, robotId := range ids {
		go func(robotId string) {
			defer wg.Done()
			action := actionMap[robotId]
			mutex.Lock()
			robot := robotMap[robotId]
			mutex.Unlock()

			dv := action.V - robot.V
			dv = clip(dv, -model.MAX_STEP_V, model.MAX_STEP_V)
			dw := action.W - robot.W
			dw = clip(dw, -model.MAX_STEP_W, model.MAX_STEP_W)

			robot.V += dv
			robot.V = clip(robot.V, -model.MAX_V, model.MAX_V)
			robot.W += dw
			robot.W = clip(robot.W, -model.MAX_W, model.MAX_W)

			if math.Abs(robot.W) < 0.001 {
				robot.X = robot.X + robot.V * math.Cos(robot.Theta) * model.DT
				robot.Y = robot.Y + robot.V * math.Sin(robot.Theta) * model.DT
			} else {
				robot.X = robot.X - (robot.V/robot.W) * math.Sin(robot.Theta) + (robot.V/robot.W) * math.Sin(robot.Theta + robot.W * model.DT)
				robot.Y = robot.Y + (robot.V/robot.W) * math.Cos(robot.Theta) - (robot.V/robot.W) * math.Cos(robot.Theta + robot.W * model.DT)
			}
			robot.X = clip(robot.X, -model.MAX_X, model.MAX_X)
			robot.Y = clip(robot.Y, -model.MAX_Y, model.MAX_Y)

			robot.Theta += robot.W * model.DT
			robot.Theta = normalize(robot.Theta)
			mutex.Lock()
			robotMap[robotId] = robot
			mutex.Unlock()
		}(robotId)
	}
	wg.Wait()
}

func ShowInfo(robotMap map[string]model.Robot){
	for robotId, robot := range robotMap {
		println(robotId, robot.X, robot.Y, robot.Theta)
		println(robotId, robot.V, robot.W)
	}
}