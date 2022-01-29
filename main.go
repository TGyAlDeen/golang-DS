package main

import "fmt"

var size = 10
var arr = make([]int,size)

func main(){
	initArr()
	printArr()
	rotateRight()
	printArr()
	rotateLeft()
	printArr()
	fmt.Println(pop(3))
	printArr()
}

func pop(idx int) int{
	val := arr[idx]
	// shift all data to left
	for i := idx+1; i < size; i++ {
		arr[i-1] = arr[i]
	}
	size--
	return val
}

func rotateRight(){
	lastElm := arr[size-1]
	for i := size -2; i >=0 ; i-- {
		arr[i+1] = arr[i]
	}
	arr[0] = lastElm
}

func rotateLeft(){
	fisrElm := arr[0]
	for i := 1; i < size; i++ {
		arr[i-1] = arr[i]
	}
	arr[size-1] = fisrElm

}


 func initArr(){
	 for i := 0; i < 10; i++ {
		 arr[i] =i
	 }
 }


func printArr() {
	fmt.Println(arr)
}