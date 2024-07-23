package algs

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

func (d *MyCircularQueue) EnQueue(value int) bool {
	if d.IsFull() {
		return false
	}
	d.items[d.front] = value
	d.front = (d.front + 1) % len(d.items)
	d.size++
	return true
}

func (d *MyCircularQueue) DeQueue() bool {
	if d.IsEmpty() {
		return false
	}
	d.items[d.getRearIndex()] = -1
	d.size--
	return true
}

func (d *MyCircularQueue) getFrontIndex() int {
	return (d.front - 1 + len(d.items)) % len(d.items)
}

func (d *MyCircularQueue) getRearIndex() int {
	return (d.front - d.size + len(d.items)) % len(d.items)
}

func (d *MyCircularQueue) Front() int {
	if d.IsEmpty() {
		return -1
	}
	return d.items[d.getRearIndex()]
}

func (d *MyCircularQueue) Rear() int {
	if d.IsEmpty() {
		return -1
	}
	return d.items[d.getFrontIndex()]
}

func (d *MyCircularQueue) IsEmpty() bool {
	return d.size == 0
}

func (d *MyCircularQueue) IsFull() bool {
	return d.size == len(d.items)
}
