package pkg

import "fmt"

func QueueMain() {
	myQueue := QueueConstructor(3)
	fmt.Printf("items: %+v, %d, %d\n", myQueue, myQueue.getRearIndex(), myQueue.Rear())
	fmt.Println(myQueue.EnQueue(1)) // return True
	fmt.Printf("EnQueue, items: %+v, %d, %d\n", myQueue, myQueue.getRearIndex(), myQueue.Rear())
	fmt.Println(myQueue.EnQueue(2)) // return True
	fmt.Printf("EnQueue, items: %+v, %d, %d\n", myQueue, myQueue.getRearIndex(), myQueue.Rear())
	fmt.Println(myQueue.EnQueue(3)) // return True
	fmt.Printf("EnQueue, items: %+v, %d, %d\n", myQueue, myQueue.getRearIndex(), myQueue.Rear())
	fmt.Println(myQueue.EnQueue(4)) // return False
	fmt.Printf("EnQueue, items: %+v, %d, %d\n", myQueue, myQueue.getRearIndex(), myQueue.Rear())
	fmt.Println(myQueue.Rear()) // return 3
	fmt.Printf("getRear, items: %+v, %d, %d\n", myQueue, myQueue.getRearIndex(), myQueue.Rear())
	fmt.Println(myQueue.IsFull()) // return True
	fmt.Printf("isFull, items: %+v, %d, %d\n", myQueue, myQueue.getRearIndex(), myQueue.Rear())
	fmt.Println(myQueue.DeQueue()) // return True
	fmt.Printf("DeQueue, items: %+v, %d, %d\n", myQueue, myQueue.getRearIndex(), myQueue.Rear())
	fmt.Println(myQueue.EnQueue(4)) // return True
	fmt.Printf("EnQueue, items: %+v, %d, %d\n", myQueue, myQueue.getRearIndex(), myQueue.Rear())
	fmt.Println(myQueue.Rear()) // return 4
	fmt.Printf("Rear, items: %+v, %d, %d\n", myQueue, myQueue.getRearIndex(), myQueue.Rear())

	d := QueueConstructor(5)
	fmt.Printf("%+v, %+v\n", d.DeQueue(), d)
	fmt.Printf("%+v, %+v\n", d.DeQueue(), d)
	fmt.Printf("%+v, %+v\n", d.EnQueue(103), d)
	fmt.Printf("%+v, %+v\n", d.EnQueue(102), d)
	fmt.Printf("%+v, %+v\n", d.EnQueue(101), d)
	fmt.Printf("%+v, %+v\n", d.EnQueue(100), d)
	fmt.Printf("%+v, %+v\n", d.EnQueue(99), d)
	fmt.Printf("%+v, %+v\n", d.DeQueue(), d)
	fmt.Printf("%+v, %+v\n", d.DeQueue(), d)
	fmt.Printf("%+v, %+v\n", d.DeQueue(), d)
	fmt.Printf("%+v, %+v\n", d.EnQueue(98), d)
	fmt.Printf("%+v, %+v\n", d.EnQueue(97), d)
	fmt.Printf("%+v, %+v\n", d.EnQueue(96), d)
	fmt.Printf("%+v, %+v\n", d.DeQueue(), d)
	fmt.Printf("%+v, %+v\n", d.EnQueue(98), d)
	fmt.Printf("%+v, %+v\n", d.DeQueue(), d)
	fmt.Printf("%+v, %+v\n", d.EnQueue(97), d)
	fmt.Printf("%+v, %+v\n", d.DeQueue(), d)
	fmt.Printf("%+v, %+v\n", d.EnQueue(96), d)
	fmt.Printf("%+v, %+v\n", d.DeQueue(), d)
	fmt.Printf("%+v, %+v\n", d.EnQueue(95), d)
	fmt.Printf("%+v, %+v\n", d.DeQueue(), d)
	fmt.Printf("%+v, %+v\n", d.EnQueue(94), d)
	fmt.Printf("%+v, %+v\n", d.DeQueue(), d)
	fmt.Printf("%+v, %+v\n", d.EnQueue(93), d)
	fmt.Printf("%+v, %+v\n", d.DeQueue(), d)
	fmt.Printf("%+v, %+v\n", d.EnQueue(92), d)

	q := QueueConstructor(2)
	fmt.Printf("%+v, %+v\n", q.EnQueue(1), q)
	fmt.Printf("%+v, %+v\n", q.EnQueue(2), q)
	fmt.Printf("%+v, %+v\n", q.DeQueue(), q)
	fmt.Printf("%+v, %+v\n", q.EnQueue(3), q)
	fmt.Printf("%+v, %+v\n", q.DeQueue(), q)
	fmt.Printf("%+v, %+v\n", q.EnQueue(3), q)
	fmt.Printf("%+v, %+v\n", q.DeQueue(), q)
	fmt.Printf("%+v, %+v\n", q.EnQueue(3), q)
	fmt.Printf("%+v, %+v\n", q.DeQueue(), q)
	fmt.Printf("%+v, %+v\n", q.Front(), q)
}

type MyCircularQueue struct {
	items []int
	front int
	size  int
}

func QueueConstructor(k int) MyCircularQueue {
	s := make([]int, k)
	for i := 0; i < k; i++ {
		s[i] = -1
	}
	return MyCircularQueue{items: s, front: 0, size: 0}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.items[this.front] = value
	this.front = (this.front + 1) % len(this.items)
	this.size++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.items[this.getRearIndex()] = -1
	this.size--
	return true
}

func (this *MyCircularQueue) getFrontIndex() int {
	return (this.front - 1 + len(this.items)) % len(this.items)
}

func (this *MyCircularQueue) getRearIndex() int {
	return (this.front - this.size + len(this.items)) % len(this.items)
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.items[this.getRearIndex()]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.items[this.getFrontIndex()]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.size == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.size == len(this.items)
}
