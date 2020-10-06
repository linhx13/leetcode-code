package main

type CheckInInfo struct {
	station string
	time    int
}

type UndergroundSystem struct {
	CheckInMap map[int]CheckInInfo
	TotalTime  map[string][]int // route (start + "_" + end) -> time_list
}

func Constructor() UndergroundSystem {
	var us UndergroundSystem
	us.CheckInMap = make(map[int]CheckInInfo)
	us.TotalTime = make(map[string][]int)
	return us
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	this.CheckInMap[id] = CheckInInfo{stationName, t}
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	cii := this.CheckInMap[id]
	route := cii.station + "_" + stationName
	totalTime := t - cii.time
	this.TotalTime[route] = append(this.TotalTime[route], totalTime)
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	route := startStation + "_" + endStation
	totalTime := 0
	for _, t := range this.TotalTime[route] {
		totalTime += t
	}
	return float64(totalTime) / float64(len(this.TotalTime[route]))
}

func main() {

}
