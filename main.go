package main

import (
	"fmt"
	"slices"
	"unsafe"
)

func main() {

	// Игра со значениями слайса, длинной и емкостью
	/*

		data := make([]int, 3, 6)
		fmt.Println("data elements:", data, len(data), cap(data))
		slice := data[1:3]
		fmt.Println("slice elements:", slice, len(slice), cap(slice))
		//data elements: [0 0 0] 3 6
		//slice elements: [0 0] 2 5
		fmt.Println()

		data[1] = 1
		fmt.Println("data elements:", data, len(data), cap(data))
		fmt.Println("slice elements:", slice, len(slice), cap(slice))
		//data elements: [0 1 0] 3 6
		//slice elements: [1 0] 2 5
		fmt.Println()

		slice = append(slice, 2)
		fmt.Println("slice elements:", slice, len(slice), cap(slice))
		fmt.Println("data elements:", data, len(data), cap(data))
		//slice elements: [1 0 2] 3 5
		//data elements: [0 1 0] 3 6
		fmt.Println()

		data = append(data, 3)
		fmt.Println("slice elements:", slice, len(slice), cap(slice))
		fmt.Println("data elements:", data, len(data), cap(data))
		//slice elements: [1 0 3] 3 5
		//data elements: [0 1 0 3] 4 6
		fmt.Println()

	*/

	// Игра со значениями слайса, длинной и емкостью
	/*
		data1 := make([]int, 4, 6)
		data2 := data1[2:6] // data2 elements: [0 0 0 0] 4 4
		fmt.Println("data1 elements:", data1, len(data1), cap(data1))
		fmt.Println("data2 elements:", data2, len(data2), cap(data2))
		fmt.Println()

		data2 = append(data2, 1) // // data2 elements: [0 0 0 0 1] 4 8 (new *Array)
		fmt.Println("data1 elements:", data1, len(data1), cap(data1))
		fmt.Println("data2 elements:", data2, len(data2), cap(data2))
		fmt.Println()

		data1[1] = 100 // data1 elements: [0 1 0 0] 4 4
		fmt.Println()
		// data1 elements: [0 0 0 0] 4 6
		//data2 elements: [0 0 0 0] 4 4
		//
		//data1 elements: [0 0 0 0] 4 6
		//data2 elements: [0 0 0 0 1] 5 8

	*/

	// Получение массива из слайса копированием
	/*
		slice := make([]int, 3, 6)
		array := [3]int(slice[:3])
		slice[0] = 10

		fmt.Println("slice elements:", slice, len(slice), cap(slice))
		fmt.Println("array elements:", array, len(array), cap(array))
		// slice elements: [10 0 0] 3 6
		//array elements: [0 0 0] 3 3
	*/

	// Получение массива из слайса по указателю
	/*
		slice := make([]int, 3, 6)
		array := (*[3]int)(slice)
		slice[0] = 10
		fmt.Println("slice elements:", slice, len(slice), cap(slice))
		fmt.Println("array elements:", *array, len(array), cap(array))
	*/

	// Удаление последнего элемента из среза
	/*
		data := []int{1, 2}       // len(2), cap(2)
		data = data[:len(data)-1] // len(1), cap(1)
		fmt.Println(data)
		fmt.Println()
		// [1]
	*/

	// Удаление последнего элемента из среза c пакетом "slices"
	/*
		data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		data = slices.Delete(data, 0, 3)
		fmt.Println("data elements:", data, len(data), cap(data))
		fmt.Println()
		// data elements: [3 4 5 6 7 8 9] 7 10
	*/

	// Увеличение длинны до емкости при помощи нарезки
	/*
		data := make([]int, 3, 6) // 0 0 0 0 0 0 / len(3) / cap(6)
		data = data[:6]
		fmt.Println("data elements:", data, len(data), cap(data))
		fmt.Println()
		// data elements: [0 0 0 0 0 0] 6 6
	*/

	// Как релоцировать емкость среза
	///*
	data := make([]int, 5, 100<<20) // 100МБ
	fmt.Println("data elements:", data, "pointer:", unsafe.Pointer(&data), len(data), cap(data))
	// data elements: [0 0 0 0 0] pointer: 0x1400012e000 5 104857600
	fmt.Println()

	temp := data[:]
	fmt.Println("temp1 elements:", temp, "pointer:", unsafe.Pointer(&temp), len(temp), cap(temp))
	// temp1 elements: [0 0 0 0 0] pointer: 0x1400012e030 5 104857600
	fmt.Println()

	temp2 := slices.Clip(temp)
	fmt.Println("temp2 elements:", temp2, "pointer:", unsafe.Pointer(&temp2), len(temp2), cap(temp2))
	fmt.Println()
	// Но по фатку мы не избавились от памяти, так как у нас поинтер
	// temp2 elements: [0 0 0 0 0] pointer: 0x1400000e078 5 5
	//*/
}
