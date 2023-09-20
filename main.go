package main

import (
	"List/list"
	"fmt"
)

func main() {
	firstList := list.NewList()

	fmt.Printf("Create List. Len: %d\n", firstList.Len())

	firstList.Add(5)
	idx, _ := firstList.GetByValue(5)
	fmt.Printf("Add 5. Len: %d, Index of 5: %d\n", firstList.Len(), idx)

	firstList.Add(10)
	idx, _ = firstList.GetByValue(10)
	fmt.Printf("Add 10. Len: %d, Index of 10: %d\n", firstList.Len(), idx)

	firstList.Add(20)
	idx, _ = firstList.GetByValue(20)
	fmt.Printf("Add 20. Len: %d, Index of 20: %d\n", firstList.Len(), idx)

	firstList.Add(30)
	idx, _ = firstList.GetByValue(30)
	fmt.Printf("Add 30. Len: %d, Index of 30: %d\n", firstList.Len(), idx)

	firstList.Add(5)
	firstList.Add(6)
	firstList.Add(5)
	fmt.Printf("Add 5, 6, 5. Len: %d\n", firstList.Len())
	firstList.Print()

	firstList.RemoveByIndex(0)
	idx, _ = firstList.GetByValue(5)
	fmt.Printf("Remove 0 element in list. Len: %d, Index of second 5: %d\n", firstList.Len(), idx)

	firstList.RemoveByValue(30)
	value, _ := firstList.GetByIndex(3)
	fmt.Printf("Remove value 30 from list. Len: %d, 3 element have value: %d\n", firstList.Len(), value)
	firstList.Print()

	firstList.RemoveAllByValue(5)
	value, _ = firstList.GetByIndex(2)
	fmt.Printf("Remove all 5 value from list. Len: %d, 2 element have value: %d\n", firstList.Len(), value)

	firstList.Clear()
	fmt.Printf("Clear list. Len:%d\n", firstList.Len())

	firstList.Add(5)
	firstList.Add(5)
	firstList.Add(5)
	firstList.Add(5)
	firstList.Add(6)
	firstList.Add(5)
	ids, _ := firstList.GetAllByValue(5)

	fmt.Print("Add 5, 5, 5, 5, 6, 5 in list. All 5 values have indexes: ")
	for i := 0; i < len(ids); i++ {
		fmt.Printf("%d ", ids[i])
	}
	fmt.Println()

	values, _ := firstList.GetAll()
	fmt.Printf("All lists value: ")
	for i := 0; i < len(values); i++ {
		fmt.Printf("%d ", values[i])
	}
	fmt.Println()
}
