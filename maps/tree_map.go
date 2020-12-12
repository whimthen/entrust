package maps

type TreeMap struct {
	tree *RBTree
}

func NewStringTreeMap() *TreeMap {
	return NewWithComparator(StringComparator)
}

func NewDecimalTreeMap() *TreeMap {
	return NewWithComparator(DecimalComparator)
}

func NewWithComparator(comparator Comparator) *TreeMap {
	return &TreeMap{tree: &RBTree{compare: comparator}}
}

func (t *TreeMap) Size() int {
	return t.tree.size
}

func (t *TreeMap) Get(key interface{}) (value interface{}, found bool) {
	return
}

func (t *TreeMap) Put(key, value interface{}) {
	t.tree.put(key, value)
}

func (t *TreeMap) Remove(key interface{}) (ok bool, err error) {
	return
}

func (t *TreeMap) Empty() bool {
	return t.Size() <= 0
}

func (t *TreeMap) Clear() {
	return
}

func (t *TreeMap) Keys() []interface{} {
	return nil
}

func (t *TreeMap) Values() interface{} {
	return nil
}

func (t *TreeMap) ContainsKey(key interface{}) bool {
	return true
}

func (t *TreeMap) ContainsValue(value interface{}) bool {
	return true
}

func (t *TreeMap) Each(func(key interface{}, value interface{})) {
	return
}

func (t *TreeMap) PutIfAbsent(key, value interface{}) interface{} {
	return nil
}
