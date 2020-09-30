package main

type SnapshotArray struct {
	snapId int
	snaps  map[int]map[int]int
}

func Constructor(length int) SnapshotArray {
	return SnapshotArray{snaps: make(map[int]map[int]int)}
}

func (this *SnapshotArray) Set(index int, val int) {
	if _, ok := this.snaps[index]; !ok {
		this.snaps[index] = make(map[int]int)
	}
	this.snaps[index][this.snapId] = val
}

func (this *SnapshotArray) Snap() int {
	this.snapId++
	return this.snapId - 1
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	data := this.snaps[index]
	for ; snap_id >= 0; snap_id-- {
		if _, ok := data[snap_id]; ok {
			break
		}
	}
	if snap_id < 0 {
		return 0
	} else {
		return data[snap_id]
	}
}
