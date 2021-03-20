package main

import "fmt"

func merge_sort(left, right [] int) (result [] int){
	result = make([] int, len(left)+len(right))
    i := 0

    for len(left) > 0 && len(right) >0 {
    	if left[0] < right[0]{
    		result[i] = left[0]
    		left = left[1:]
		}else{
			result[i] =  right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++{
		result[i] = left[j]
		i++
	}
    for j:=0; j < len(right); j++{
    	result[i] =  right[j]
    	i++
	}

	return result
}


func merge(arr [] int) [] int {

	n := len(arr)
	//fmt.Println("arr:", arr)

	if n == 1 {
		return arr
	}else{
		mid := int(n / 2)
		right := merge(arr[mid:])
		left := merge(arr[0:mid])
		fmt.Println(right)
		fmt.Println(left)
		return merge_sort(left, right)
	}

}

func main(){

	 var arr = [] int {5,1,6,3,4,7}

	 fmt.Println(merge(arr))

}
