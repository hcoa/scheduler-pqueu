package pqueue

import "math"

// Item - a separate element. Can be changed. Only priority field is used for data structure
type Item struct {
	id       int
	action   string
	priority int
}

// PQueue - container for Item
type PQueue struct {
	MaxSize int
	Items   []Item
}

// NewPQueue - return pointer to PQueue instance
func NewPQueue(size int) *PQueue {
	return &PQueue{
		MaxSize: size * 2,
		Items:   []Item{},
	}
}

func (p *PQueue) parent(i int) int {
	return (i - 1) / 2
}

func (p *PQueue) leftChild(i int) int {
	return 2*i + 1
}

func (p *PQueue) rightChild(i int) int {
	return 2*i + 2
}

func (p *PQueue) siftUp(i int) {
	for i > 0 && p.Items[p.parent(i)].priority < p.Items[i].priority {
		p.Items[p.parent(i)], p.Items[i] = p.Items[i], p.Items[p.parent(i)]
		i = p.parent(i)
	}
}

func (p *PQueue) siftDown(i int) {
	maxInd := i
	l := p.leftChild(i)
	if l < p.Size() && p.Items[l].priority > p.Items[maxInd].priority {
		maxInd = l
	}
	r := p.rightChild(i)
	if r < p.Size() && p.Items[r].priority > p.Items[maxInd].priority {
		maxInd = r
	}
	if i != maxInd {
		p.Items[i], p.Items[maxInd] = p.Items[maxInd], p.Items[i]
		p.siftDown(maxInd)
	}
}

// Insert - add new Item to the priority queue
func (p *PQueue) Insert(item Item) {
	p.Items = append(p.Items, item)
	p.siftUp(p.Size() - 1)
}

// ExctractMax - pop element with the biggest priority
func (p *PQueue) ExctractMax() Item {
	r := p.Items[0]
	x := p.Size() - 1
	p.Items[0] = p.Items[x]

	p.Items = p.Items[:x]
	p.siftDown(0)
	return r
}

// Remove - remove element by index
func (p *PQueue) Remove(i int) {
	p.Items[i] = Item{
		id:       0,
		action:   "",
		priority: math.MaxInt32,
	}
	p.siftUp(i)
	p.ExctractMax()
}

// ChangePriority - change elements priority
func (p *PQueue) ChangePriority(i int, priority int) {
	oldP := p.Items[i]
	p.Items[i].priority = priority
	if priority > oldP.priority {
		p.siftUp(i)
	} else {
		p.siftDown(i)
	}
}

// Size - return the len of the slice
func (p *PQueue) Size() int {
	return len(p.Items)
}

// func (p *PQueue) resize(n int) -- need for shrink the cap of the slice
