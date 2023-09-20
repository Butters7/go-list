package list

import (
	"fmt"
	"testing"
)

func TestAdding(t *testing.T) {
	fList := NewList()

	var i int64
	for i = 0; i < 100; i++ {
		fList.Add(i)
	}

	if fList.len != int64(100) {
		t.Fatal("Test \"Adding\" failed")
	}

	fmt.Println("Test \"Adding\" OK")
}

func TestRemoveByIndexAndGetByValue(t *testing.T) {
	fList := NewList()

	var i int64
	for i = 0; i < 100; i++ {
		fList.Add(i)
	}

	fList.RemoveByIndex(50)

	idx, _ := fList.GetByValue(51)
	if idx != int64(50) {
		t.Fatal("Test \"RemoveByIndexAndGetByValue\" failed")
	}

	fmt.Println("Test \"RemoveByIndexAndGetByValue\" OK")
}

func TestRemoveByValueAndGeyByIndex(t *testing.T) {
	fList := NewList()

	var i int64
	for i = 0; i < 100; i++ {
		fList.Add(i)
	}

	fList.RemoveByValue(50)

	value, _ := fList.GetByIndex(50)
	if value != int64(51) {
		t.Fatal("Test \"RemoveByValueAndGeyByIndex\" failed")
	}

	fmt.Println("Test \"RemoveByValueAndGeyByIndex\" OK")
}

func TestRemoveAllByValueAndGetAllByValue(t *testing.T) {
	fList := NewList()

	var fValue int64 = 50
	var sValue int64 = 25

	for i := 0; i < 5; i++ {
		fList.Add(fValue)
	}

	for i := 0; i < 2; i++ {
		fList.Add(sValue)
	}

	for i := 0; i < 5; i++ {
		fList.Add(fValue)
	}

	for i := 0; i < 2; i++ {
		fList.Add(sValue)
	}

	fList.RemoveAllByValue(fValue)

	if fList.len != int64(4) {
		t.Fatalf("Test \"RemoveAllByValueAndGetAllByValue\" failed")
	}

	ids, _ := fList.GetAllByValue(sValue)

	for i := 0; i < len(ids); i++ {
		if int64(i) != ids[i] {
			t.Fatalf("Test \"RemoveAllByValueAndGetAllByValue\" failed")
		}
	}

	fmt.Println("Test \"RemoveAllByValueAndGetAllByValue\" OK")
}

func TestGetAll(t *testing.T) {
	fList := NewList()

	var i int64
	for i = 0; i < 50; i++ {
		fList.Add(i)
	}

	values, _ := fList.GetAll()

	for i = 0; i < 50; i++ {
		if values[i] != i {
			t.Fatalf("Test \"GetAll\" failed")
		}
	}

	fmt.Println("Test \"GetAll\" OK")
}
