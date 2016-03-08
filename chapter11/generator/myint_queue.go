package main

type MyIntQueue struct {
	q []MyInt
}

func NewMyIntQueue() *MyIntQueue {
	return &MyIntQueue{
		q: []MyInt{},
	}
}

func (o *MyIntQueue) Insert(v MyInt) {
	o.q = append(o.q, v)
}

func (o *MyIntQueue) Remove() MyInt {
	if len(o.q) == 0 {
		panic("Oops.")
	}
	first := o.q[0]
	o.q = o.q[1:]
	return first
}
