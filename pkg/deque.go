package pkg

import "fmt"

func CircularDequeMain() {
	myCircularDeque := Constructor(3)
	fmt.Printf("items: %+v, %d, %d\n", myCircularDeque, myCircularDeque.getRearIndex(), myCircularDeque.GetRear())
	fmt.Println(myCircularDeque.InsertLast(1)) // return True
	fmt.Printf("insertLast, items: %+v, %d, %d\n", myCircularDeque, myCircularDeque.getRearIndex(), myCircularDeque.GetRear())
	fmt.Println(myCircularDeque.InsertLast(2)) // return True
	fmt.Printf("insertLast, items: %+v, %d, %d\n", myCircularDeque, myCircularDeque.getRearIndex(), myCircularDeque.GetRear())
	fmt.Println(myCircularDeque.InsertFront(3)) // return True
	fmt.Printf("insertFront, items: %+v, %d, %d\n", myCircularDeque, myCircularDeque.getRearIndex(), myCircularDeque.GetRear())
	fmt.Println(myCircularDeque.InsertFront(4)) // return False, the queue is full.
	fmt.Printf("insertFront, items: %+v, %d, %d\n", myCircularDeque, myCircularDeque.getRearIndex(), myCircularDeque.GetRear())
	fmt.Println(myCircularDeque.GetRear()) // return 2
	fmt.Printf("getRear, items: %+v, %d, %d\n", myCircularDeque, myCircularDeque.getRearIndex(), myCircularDeque.GetRear())
	fmt.Println(myCircularDeque.IsFull()) // return True
	fmt.Printf("isFull, items: %+v, %d, %d\n", myCircularDeque, myCircularDeque.getRearIndex(), myCircularDeque.GetRear())
	fmt.Println(myCircularDeque.DeleteLast()) // return True
	fmt.Printf("deleteLast, items: %+v, %d, %d\n", myCircularDeque, myCircularDeque.getRearIndex(), myCircularDeque.GetRear())
	fmt.Println(myCircularDeque.InsertFront(4)) // return True
	fmt.Printf("insertFront, items: %+v, %d, %d\n", myCircularDeque, myCircularDeque.getRearIndex(), myCircularDeque.GetRear())
	fmt.Println(myCircularDeque.GetFront()) // return 4
	fmt.Printf("getFront, items: %+v, %d, %d\n", myCircularDeque, myCircularDeque.getRearIndex(), myCircularDeque.GetRear())

	d := Constructor(5)
	fmt.Printf("%+v, %+v\n", d.InsertFront(18), d)
	fmt.Printf("%+v, %+v\n", d.InsertFront(19), d)
	fmt.Printf("%+v, %+v\n", d.InsertFront(20), d)
	fmt.Printf("%+v, %+v\n", d.InsertFront(21), d)
	fmt.Printf("%+v, %+v\n", d.DeleteFront(), d)
	fmt.Printf("%+v, %+v\n", d.DeleteFront(), d)
	fmt.Printf("%+v, %+v\n", d.InsertLast(99), d)
	fmt.Printf("%+v, %+v\n", d.InsertFront(22), d)
	fmt.Printf("%+v, %+v\n", d.InsertFront(23), d)
	fmt.Printf("%+v, %+v\n", d.InsertFront(24), d)
	fmt.Printf("%+v, %+v\n", d.InsertFront(25), d)
	fmt.Printf("%+v, %+v\n", d.DeleteFront(), d)
	fmt.Printf("%+v, %+v\n", d.DeleteFront(), d)
	fmt.Printf("%+v, %+v\n", d.DeleteFront(), d)
	fmt.Printf("%+v, %+v\n", d.InsertFront(26), d)
	fmt.Printf("%+v, %+v\n", d.InsertFront(27), d)
	fmt.Printf("%+v, %+v\n", d.InsertFront(28), d)
	fmt.Printf("%+v, %+v\n", d.InsertFront(29), d)
	fmt.Printf("%+v, %+v\n", d.DeleteLast(), d)
	fmt.Printf("%+v, %+v\n", d.DeleteLast(), d)
	fmt.Printf("%+v, %+v\n", d.DeleteLast(), d)
	fmt.Printf("%+v, %+v\n", d.InsertLast(99), d)
	fmt.Printf("%+v, %+v\n", d.DeleteLast(), d)
	fmt.Printf("%+v, %+v\n", d.InsertFront(30), d)
	fmt.Printf("%+v, %+v\n", d.InsertFront(31), d)
	fmt.Printf("%+v, %+v\n", d.DeleteLast(), d)
	fmt.Printf("%+v, %+v\n", d.InsertFront(32), d)
	fmt.Printf("%+v, %+v\n", d.DeleteLast(), d)
	fmt.Printf("%+v, %+v\n", d.InsertFront(33), d)
	fmt.Printf("%+v, %+v\n", d.DeleteLast(), d)
	fmt.Printf("%+v, %+v\n", d.InsertFront(34), d)
	fmt.Printf("%+v, %+v\n", d.DeleteLast(), d)
	fmt.Printf("%+v, %+v\n", d.InsertFront(35), d)
	fmt.Printf("%+v, %+v\n", d.DeleteLast(), d)
	fmt.Printf("%+v, %+v\n", d.InsertFront(36), d)
	fmt.Printf("%+v, %+v\n", d.DeleteLast(), d)
	fmt.Printf("%+v, %+v\n", d.DeleteFront(), d)
	fmt.Printf("%+v, %+v\n", d.InsertLast(98), d)
	fmt.Printf("%+v, %+v\n", d.DeleteFront(), d)
	fmt.Printf("%+v, %+v\n", d.InsertLast(97), d)
	fmt.Printf("%+v, %+v\n", d.DeleteFront(), d)
	fmt.Printf("%+v, %+v\n", d.InsertLast(96), d)
	fmt.Printf("%+v, %+v\n", d.DeleteFront(), d)
	fmt.Printf("%+v, %+v\n", d.InsertLast(95), d)
	fmt.Printf("%+v, %+v\n", d.DeleteFront(), d)
	fmt.Printf("%+v, %+v\n", d.InsertLast(94), d)
	fmt.Printf("%+v, %+v\n", d.DeleteFront(), d)
	fmt.Printf("%+v, %+v\n", d.InsertLast(93), d)
	fmt.Printf("%+v, %+v\n", d.DeleteFront(), d)
	fmt.Printf("%+v, %+v\n", d.InsertLast(92), d)
}

type MyCircularDeque struct {
	items []int
	front int
	size  int
}

func Constructor(k int) MyCircularDeque {
	s := make([]int, k)
	for i := 0; i < k; i++ {
		s[i] = -1
	}
	return MyCircularDeque{items: s, front: 0, size: 0}
}

func (d *MyCircularDeque) InsertFront(value int) bool {
	if d.IsFull() {
		return false
	}
	d.items[d.front] = value
	d.front = (d.front + 1) % len(d.items)
	d.size++
	return true
}

func (d *MyCircularDeque) InsertLast(value int) bool {
	if d.IsFull() {
		return false
	}
	d.items[d.getRearIndex()] = value
	d.size++
	return true
}

func (d *MyCircularDeque) DeleteFront() bool {
	if d.IsEmpty() {
		return false
	}
	d.size--
	d.front = (d.front - 1 + len(d.items)) % len(d.items)
	d.items[d.front] = -1
	return true
}

func (d *MyCircularDeque) DeleteLast() bool {
	if d.IsEmpty() {
		return false
	}
	d.size--
	d.items[d.getRearIndex()] = -1
	return true
}

func (d *MyCircularDeque) getFrontIndex() int {
	return (d.front - 1 + len(d.items)) % len(d.items)
}

func (d *MyCircularDeque) GetFront() int {
	if d.IsEmpty() {
		return -1
	}
	return d.items[d.getFrontIndex()]
}

func (d *MyCircularDeque) getRearIndex() int {
	return (d.front - d.size - 1 + len(d.items)) % len(d.items)
}

func (d *MyCircularDeque) GetRear() int {
	if d.IsEmpty() {
		return -1
	}
	return d.items[(d.getRearIndex()+1)%len(d.items)]
}

func (d *MyCircularDeque) IsEmpty() bool {
	return d.size == 0
}

func (d *MyCircularDeque) IsFull() bool {
	return d.size == len(d.items)
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
