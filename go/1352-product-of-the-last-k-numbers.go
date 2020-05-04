package main

import "fmt"

type ProductOfNumbers struct {
	nums []int
}

func Constructor() ProductOfNumbers {
	p := ProductOfNumbers{}
	p.nums = append(p.nums, 1)
	return p
}

func (this *ProductOfNumbers) Add(num int) {
	if num != 0 {
		this.nums = append(this.nums, this.nums[len(this.nums)-1]*num)
	} else {
		this.nums = []int{1}
	}
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	n := len(this.nums)
	if k < n {
		return this.nums[n-1] / this.nums[n-k-1]
	} else {
		return 0
	}
}

func main() {
	ops := []string{"ProductOfNumbers", "add", "add", "add", "add", "add", "getProduct", "getProduct", "getProduct", "add", "getProduct"}
	params := []int{0, 3, 0, 2, 5, 4, 2, 3, 4, 8, 2}
	var obj ProductOfNumbers
	for i := 0; i < len(ops); i++ {
		switch ops[i] {
		case "ProductOfNumbers":
			obj = Constructor()
		case "add":
			obj.Add(params[i])
		case "getProduct":
			fmt.Println(obj.GetProduct(params[i]))
		}
	}

}
