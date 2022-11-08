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

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.items[this.front] = value
	this.front = (this.front + 1) % len(this.items)
	this.size++
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.items[this.getRearIndex()] = value
	this.size++
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.size--
	this.front = (this.front - 1 + len(this.items)) % len(this.items)
	this.items[this.front] = -1
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.size--
	this.items[this.getRearIndex()] = -1
	return true
}

func (this *MyCircularDeque) getFrontIndex() int {
	return (this.front - 1 + len(this.items)) % len(this.items)
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.items[this.getFrontIndex()]
}

func (this *MyCircularDeque) getRearIndex() int {
	return (this.front - this.size - 1 + len(this.items)) % len(this.items)
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.items[(this.getRearIndex()+1)%len(this.items)]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.size == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.size == len(this.items)
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
