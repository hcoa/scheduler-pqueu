// idea: wrapper or interface to work with PQueue

// interface{} two words  [type;data|pointer]

```
	# for 1-based
	Parent(i) return i/2
	LeftChild(i) return 2*i
	RightChild(i) return 2*i + 1

	# for 0-based
	Parent(i) return (i-1)/2
	LeftChild(i) return 2*i + 1
	RightChild(i) return 2*i + 2


	SiftUp(i)
		while i > 1 and H[Parent(i)] < H[i]:
			swap H[Parent(i)] and H[i]
			i <- Parent(i)


	SiftDown(i)
		maxIndex = i
		l <- LeftChild(i)
		if l <= size and H[l] > H[maxIndex]:
			maxIndex = l
		r <- RightChild(i)
		if r <= size and H[r] > H[maxIndex]:
			maxIndex = r
		if i != maxIndex:
			swap H[i] and H[maxIndex]
			SiftDown(maxIndex)


	Insert(p)
		if size == maxSize: return ERROR
		size = size + 1
		H[size] = p
		SiftUp(size)

	ExtractMax()
		result = H[1]
		H[1] = H[size]
		size -= 1
		SiftDown(1)
		return result

	Remove(i)
		H[i] = infinity
		SiftUp(i)
		ExtractMax()

	ChangePriority(i, p)
		oldp = H[i]
		H[i] = p
		if p > oldp:
			SiftUp(i)
		else:
			SiftDown(i)
```