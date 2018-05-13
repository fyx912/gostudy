package main

import "fmt"

func main(){
	arrays := [] int{3,8,9,10,21,2,5,11}
	fmt.Println(SelectSort(arrays))
}
/*简单选择排序*/
func SelectSort(a []int) []int{
	n := len(a)
	for i := 0; i < n-1; i++ {
		min := i
		for k := i+1; k < n; k++ {
			if (a[k] < a[min]){
				min = k
			}
		}
		if i!= min {
			temp := a[i]
			a[i] = a[min]
			a[min] = temp
		}
	}
	return a
}
