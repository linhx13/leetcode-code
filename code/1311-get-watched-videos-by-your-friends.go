package main

import "sort"

type Item struct {
	video string
	freq  int
}

func watchedVideosByFriends(watchedVideos [][]string, friends [][]int, id int, level int) []string {
	visited := make([]bool, len(watchedVideos))
	queue := [][]int{}
	queue = append(queue, []int{id, 0})
	targetIds := make(map[int]bool)
	for len(queue) > 0 {
		cur_id, cur_level := queue[0][0], queue[0][1]
		queue = queue[1:len(queue)]
		if visited[cur_id] {
			continue
		}
		visited[cur_id] = true
		if cur_level == level {
			targetIds[cur_id] = true
		} else if cur_level < level {
			for _, i := range friends[cur_id] {
				queue = append(queue, []int{i, cur_level + 1})
			}
		}
	}
	m := make(map[string]int)
	for k, _ := range targetIds {
		for _, v := range watchedVideos[k] {
			m[v]++
		}
	}
	items := []Item{}
	for k, v := range m {
		items = append(items, Item{video: k, freq: v})
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].freq == items[j].freq {
			return items[i].video < items[j].video
		} else {
			return items[i].freq < items[j].freq
		}
	})
	res := []string{}
	for _, item := range items {
		res = append(res, item.video)
	}
	return res
}
