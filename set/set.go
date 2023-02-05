package set

type Set struct {
	container map[int]struct{}
}

func NewSet() Set {
	return Set{
		container: make(map[int]struct{}),
	}
}

func (set Set) Add(x int) {
	set.container[x] = struct{}{}
}

func (set Set) Remove(x int) {
	delete(set.container, x)
}

func (set Set) Len() int {
	return len(set.container)
}

func (set Set) Contains(x int) bool {
	_, found := set.container[x]
	return found
}
