package set

type IntSet struct {
	container map[int]struct{}
}

func NewIntSet() IntSet {
	return IntSet{
		container: make(map[int]struct{}),
	}
}

func (set IntSet) Add(x int) {
	set.container[x] = struct{}{}
}

func (set IntSet) Remove(x int) {
	delete(set.container, x)
}

func (set IntSet) Len() int {
	return len(set.container)
}

func (set IntSet) Contains(x int) bool {
	_, found := set.container[x]
	return found
}
