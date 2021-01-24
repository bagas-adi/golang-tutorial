package main

import (
	"fmt"
	"strconv"
)

var data uint8 = 255 // cannot exceed 255

func CompareNumber(num1 int, num2 int) (bool, string) {
	//
	if num1 > num2 {
		return false, strconv.Itoa(num1) + " is greater than " + strconv.Itoa(num2)
	} else if num1 < num2 {
		return false, strconv.Itoa(num2) + " is greater than " + strconv.Itoa(num1)
	}
	return true, strconv.Itoa(num1) + " is equivalent to " + strconv.Itoa(num2)
}
func cobaSwitch(num int) {
	switch num {
	case 1:
		fmt.Println("satu")
	case 2:
		fmt.Println("dua")
	default:
		fmt.Println("kosong cuy")
	}
}
func loopFor(from, to int) {
	for x := from; x < to; x++ {
		fmt.Println("Loop ke " + strconv.Itoa(x))
	}
}
func loopFor2(from, to int) {
	var x int = from
	for x < to {
		fmt.Println("Loop ke " + strconv.Itoa(x))
		x++
	}
}

func createArray() {
	// var x [5]int = [5]int{1, 2, 3, 4, 5}
	x := [5]int{1, 2, 3, 4, 5}
	fmt.Println(x)
}
func createSlice() {
	var _slice []int = []int{} // kosongan
	// _slice := []int{1, 2, 3, 4, 5}
	_slice = append(_slice, 6, 7, 8, 9, 10)
	fmt.Println(_slice)
}
func accessingSlice() {
	var _slice []int // = []int{} // kosongan
	// _slice := []int{1, 2, 3, 4, 5}
	_slice = append(_slice, 6, 7, 8, 9, 10)
	_slice2 := _slice[0:3]
	fmt.Println(_slice2)
	fmt.Println(_slice2[:cap(_slice2)]) // memory cap ngikut dari variable _slice
	fmt.Println(_slice2[:len(_slice2)]) // klo length itu actual length dari variable itu
}

func editSlice() {
	// var _slice []int // = []int{} // kosongan
	_slice := []int{1, 2, 3, 4, 5}
	// _slice = append(_slice, 6, 7, 8, 9, 10)
	_slice2 := _slice[0:3]
	fmt.Println(_slice)
	_slice2[0] = 7
	// padahal yg diubah var _slice2, tapi _slice ikut ke ubah.
	// krn slice berisi pointer jadinya harus dicopy, bukan cuma di referensikan
	fmt.Println(_slice)
}

func editSliceSolution() {
	// var _slice []int // = []int{} // kosongan
	_slice := []int{1, 2, 3, 4, 5}
	// _slice = append(_slice, 6, 7, 8, 9, 10)
	// var _slice2 []int = []int{} // gagal, menyebabkan fatal runtime panic
	_slice2 := make([]int, 3)
	copy(_slice2, _slice[0:3])
	fmt.Println(_slice)
	_slice2[0] = 7
	// padahal yg diubah var _slice2, tapi _slice ikut ke ubah.
	// krn slice berisi pointer jadinya harus dicopy, bukan cuma di referensikan
	fmt.Println(_slice)
	fmt.Println(_slice2)
}

func getTwoElement(_slice []int, firstIndex, secondIndex int) []int {
	// will cause a memory leak, krn returnnya masih bisa ngakses data sumber nya
	// return _slice[first_index:second_index]
	_sliceNew := make([]int, secondIndex-firstIndex)
	copy(_sliceNew, _slice[firstIndex:secondIndex])
	return _sliceNew
}
func main() {
	fmt.Printf("hello, world\n")
	// fmt.Println(CompareNumber(2, 1))
	// fmt.Println(data)
	// cobaSwitch(3)
	// loopFor2(0, 10)
	createSlice()
}
