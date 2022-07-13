package algorithm

import "fmt"

func main() {
	array:=[]int{1,4,3,8,10}
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-1-i; j++ {
			if array [j]> array [j+1] {
				//j j+1 ,j+1 j
				array[j+1],array[j]=array[j],array[j+1]
			}
		}
	}
	fmt.Println(array)
}
