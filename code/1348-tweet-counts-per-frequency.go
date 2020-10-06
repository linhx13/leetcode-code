package main

type TweetCounts struct {
	tweets map[string][]int
}

func Constructor() TweetCounts {
	return TweetCounts{
		tweets: map[string][]int{},
	}
}

func (this *TweetCounts) RecordTweet(tweetName string, time int) {
	if _, ok := this.tweets[tweetName]; !ok {
		this.tweets[tweetName] = []int{time}
	} else {
		this.tweets[tweetName] = append(this.tweets[tweetName], time)
	}
}

func (this *TweetCounts) GetTweetCountsPerFrequency(freq string, tweetName string, startTime int, endTime int) []int {

	if _, ok := this.tweets[tweetName]; !ok {
		return []int{}
	}
	interval := 0
	switch freq {
	case "minute":
		interval = 60
	case "hour":
		interval = 3600
	case "day":
		interval = 86400
	}
	res := make([]int, (endTime-startTime+interval)/interval)
	for _, time := range this.tweets[tweetName] {
		if time < startTime || time > endTime {
			continue
		}
		cur := (time - startTime) / interval
		res[cur]++
	}
	return res
}

func main() {

}
