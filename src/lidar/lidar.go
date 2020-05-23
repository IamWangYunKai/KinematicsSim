package lidar

import (
	"../model"
	"math"
	"sync"
)

func LaserDetection(robotMap map[string]model.Robot, obstacleMap map[string]model.Obstacle, robotId string) []float64 {
	var mutex sync.Mutex
	var wg sync.WaitGroup
	wg.Add(model.LASER_NUM)
	var distanceArray []float64
	for i:=0; i<model.LASER_NUM; i++{
		go func(i int) {
			defer wg.Done()
			angle :=  float64(i) * model.DELTA_ANGLE - math.Pi
			distance := checkLaser(robotMap, obstacleMap, robotId, angle)
			mutex.Lock()
			distanceArray = append(distanceArray, distance)
			mutex.Unlock()
		}(i)
	}
	wg.Wait()
	return distanceArray
}

func checkLaser(robotMap map[string]model.Robot, obstacleMap map[string]model.Obstacle, robotId string, angle float64) float64 {
	master := robotMap[robotId]
	var distanceArray []float64
	for id, robot := range robotMap {
		if id == robotId{
			continue
		}
		distance := checkCircleIntersection(master.X, master.Y, angle, robot.X, robot.Y, model.ROBOT_SIZE)
		distanceArray = append(distanceArray, distance)
	}
	for _, obstacle := range obstacleMap {
		distance := checkCircleIntersection(master.X, master.Y, angle, obstacle.X, obstacle.Y, obstacle.R)
		distanceArray = append(distanceArray, distance)
	}
	minDist := findMin(distanceArray)
	return minDist
}

func checkCircleIntersection(x1, y1, angle, x2, y2, R float64) float64 {
	// check direct
	if angle >= - math.Phi && angle <= math.Phi {
		if x2 < x1 {
			return model.INFINITE
		}
	} else {
		if x2 > x1 {
			return model.INFINITE
		}
	}
	// calculate distance
	d := math.Abs(math.Tan(angle)*(x2 -x1) - (y2 - y1)) / math.Sqrt(math.Tan(angle)*math.Tan(angle) + 1)
	if d > R {
		return model.INFINITE
	} else {
		l := math.Sqrt(math.Pow(x2-x1,2) + math.Pow(y2-y1,2))
		return l - math.Sqrt(math.Pow(R,2) + math.Pow(d,2))
	}
}

func findMin(array []float64) (minValue float64) {
	if len(array) > 0 {
		minValue = array[0]
	}
	for i := 1; i < len(array); i++ {
		if array[i] < minValue {
			minValue = array[i]
		}
	}
	return minValue
}