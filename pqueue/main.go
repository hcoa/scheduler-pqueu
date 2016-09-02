package pqueue

import "math"

type Item struct {
	id       int
	action   string
	priority int
}

// implementation with binary max heap
type PQueue struct {
	MaxSize int
	Items   []Item
}

func NewPQueue(size int) *PQueue {
	return &PQueue{
		MaxSize: size * 2,
		Items:   []Item{},
	}
}

func (p *PQueue) Parent(i int) int {
	return (i - 1) / 2
}

func (p *PQueue) LeftChild(i int) int {
	return 2*i + 1
}

func (p *PQueue) RightChild(i int) int {
	return 2*i + 2
}

func (p *PQueue) SiftUp(i int) {
	for i > 0 && p.Items[p.Parent(i)].priority < p.Items[i].priority {
		p.Items[p.Parent(i)], p.Items[i] = p.Items[i], p.Items[p.Parent(i)]
		i = p.Parent(i)
	}
}

func (p *PQueue) SiftDown(i int) {
	maxInd := i
	l := p.LeftChild(i)
	if l < p.Size() && p.Items[l].priority > p.Items[maxInd].priority {
		maxInd = l
	}
	r := p.RightChild(i)
	if r < p.Size() && p.Items[r].priority > p.Items[maxInd].priority {
		maxInd = r
	}
	if i != maxInd {
		p.Items[i], p.Items[maxInd] = p.Items[maxInd], p.Items[i]
		p.SiftDown(maxInd)
	}
}

func (p *PQueue) Insert(item Item) {
	p.Items = append(p.Items, item)
	p.SiftUp(p.Size() - 1)
}

func (p *PQueue) ExctractMax() Item {
	r := p.Items[0]
	x := p.Size() - 1
	p.Items[0] = p.Items[x]

	p.Items = p.Items[:x]
	p.SiftDown(0)
	return r
}

func (p *PQueue) Remove(i int) {
	p.Items[i] = Item{
		id:       0,
		action:   "",
		priority: math.MaxInt32,
	}
	p.SiftUp(i)
	p.ExctractMax()
}

func (p *PQueue) ChangePriority(i int, priority int) {
	oldP := p.Items[i]
	p.Items[i].priority = priority
	if priority > oldP.priority {
		p.SiftUp(i)
	} else {
		p.SiftDown(i)
	}
}

func (p *PQueue) Size() int {
	return len(p.Items)
}

// func (p *PQueue) resize(n int) -- need for shrink the cap of the slice
