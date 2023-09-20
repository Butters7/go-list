package list

import (
	"fmt"
)

type List struct {
	len       int64
	firstNode *node
}

// NewList создает новый список
func NewList() *List {
	return &List{
		len:       0,
		firstNode: nil,
	}
}

// Len возвращает длину списка
func (l *List) Len() int64 {
	return l.len
}

// Add добавляет элемент в список и возвращает его индекс
func (l *List) Add(value int64) int64 {
	if l.firstNode == nil {
		l.len = 1
		l.firstNode = &node{
			index: 0,
			value: value,
			next:  nil,
		}
		return l.firstNode.index
	}

	var index int64 = 0
	curNode := l.firstNode

	for curNode.next != nil {
		curNode = curNode.next
		index++
	}

	l.len++
	curNode.next = &node{
		index: index + 1,
		value: value,
		next:  nil,
	}

	return curNode.next.index
}

// RemoveByIndex удаляет элемент из списка по индексу
func (l *List) RemoveByIndex(id int64) {
	if id < 0 || id >= l.len {
		return
	}

	curNode := l.firstNode
	prevNode := curNode

	for curNode.index != id {
		prevNode = curNode
		curNode = curNode.next
	}

	var index int64

	if prevNode == curNode {
		l.firstNode = curNode.next
		curNode = l.firstNode
		index = 0

	} else {
		prevNode.next = curNode.next
		curNode = prevNode.next
		index = prevNode.index + 1
	}

	for curNode != nil {
		curNode.index = index
		index++
		curNode = curNode.next
	}

	l.len--
}

// RemoveByValue удаляет элемент из списка по значению
func (l *List) RemoveByValue(value int64) {
	curNode := l.firstNode
	prevNode := curNode

	for curNode != nil && curNode.value != value {
		prevNode = curNode
		curNode = curNode.next
	}

	if curNode.value != value {
		return
	}

	var index int64

	if prevNode == curNode {
		l.firstNode = curNode.next
		index = 0
		curNode = l.firstNode

	} else {
		prevNode.next = curNode.next
		curNode = prevNode.next
		index = prevNode.index + 1
	}

	for curNode != nil {
		curNode.index = index
		index++
		curNode = curNode.next
	}

	l.len--
}

// RemoveAllByValue удаляет все элементы из списка по значению
func (l *List) RemoveAllByValue(value int64) {
	curNode := l.firstNode
	prevNode := curNode

	var wasDeleteElement bool

	for curNode != nil {
		if curNode.value == value {
			if curNode == prevNode {
				l.firstNode = curNode.next
				prevNode = l.firstNode
				curNode = l.firstNode
			} else {
				prevNode = curNode.next
				curNode = curNode.next
			}
			l.len--
			wasDeleteElement = true
			continue
		}
		prevNode = curNode
		curNode = curNode.next
	}

	if wasDeleteElement {
		curNode = l.firstNode
		var idx int64 = 0

		for curNode != nil {
			curNode.index = idx
			idx++
			curNode = curNode.next
		}
	}
}

// GetByIndex возвращает значение элемента по индексу
//
// Если элемента с таким индексом нет, то возвращается 0 и false
func (l *List) GetByIndex(id int64) (value int64, ok bool) {
	if id < 0 || id >= l.len {
		return 0, false
	}

	curNode := l.firstNode
	for curNode.index != id {
		curNode = curNode.next
	}

	return curNode.value, true
}

// GetByValue возвращает индекс первого найденного элемента по значению.
//
// Если элемента с таким значением нет, то возвращается 0 и false.
func (l *List) GetByValue(value int64) (index int64, ok bool) {
	curNode := l.firstNode

	for curNode != nil && curNode.value != value {
		curNode = curNode.next
	}

	if curNode == nil || curNode.value != value {
		return 0, false
	} else {
		return curNode.index, true
	}
}

// GetAllByValue возвращает индексы всех найденных элементов по значению
//
// Если элементов с таким значением нет, то возвращается nil и false.
func (l *List) GetAllByValue(value int64) (ids []int64, ok bool) {
	curNode := l.firstNode

	for curNode != nil {
		if curNode.value == value {
			ids = append(ids, curNode.index)
		}
		curNode = curNode.next
	}

	if len(ids) == 0 {
		return nil, false
	} else {
		return ids, false
	}
}

// GetAll возвращает все элементы списка
//
// Если список пуст, то возвращается nil и false.
func (l *List) GetAll() (values []int64, ok bool) {
	if l.len == 0 {
		return nil, false
	}

	curNode := l.firstNode
	for curNode != nil {
		values = append(values, curNode.value)
		curNode = curNode.next
	}

	return values, true
}

// Clear очищает список
// Функция в 2 строки, так как в Golang есть сборщик мусора, который высвобождает неиспольщующиеся участки памяти
func (l *List) Clear() {
	l.len = 0
	l.firstNode = nil
}

// Print выводит список в консоль
func (l *List) Print() {
	if l.len == 0 {
		return
	}

	fmt.Print("List: ")
	curNode := l.firstNode
	for curNode != nil {
		fmt.Printf("%d:%d, ", curNode.index, curNode.value)
		curNode = curNode.next
	}

	fmt.Println()
}
