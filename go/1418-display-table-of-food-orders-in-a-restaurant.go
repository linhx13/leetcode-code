package main

import (
	"sort"
	"strconv"
)

func displayTable(orders [][]string) [][]string {
	m := make(map[string]map[string]int)
	itemSet := make(map[string]bool)
	for _, order := range orders {
		tableId := order[1]
		itemName := order[2]
		itemSet[itemName] = true
		if _, ok := m[tableId]; !ok {
			m[tableId] = make(map[string]int)
		}
		m[tableId][itemName]++
	}
	itemList := []string{}
	for item, _ := range itemSet {
		itemList = append(itemList, item)
	}
	sort.Strings(itemList)
	tableList := []string{}
	for k, _ := range m {
		tableList = append(tableList, k)
	}
	sort.Slice(tableList, func(i, j int) bool {
		a, _ := strconv.Atoi(tableList[i])
		b, _ := strconv.Atoi(tableList[j])
		return a < b
	})
	res := [][]string{}
	itemList = append([]string{"Table"}, itemList...)
	res = append(res, itemList)
	for _, tableId := range tableList {
		row := []string{}
		row = append(row, tableId)
		cm := m[tableId]
		for i := 1; i < len(itemList); i++ {
			row = append(row, strconv.Itoa(cm[itemList[i]]))
		}
		res = append(res, row)
	}
	return res
}
