package maps

type RBTree struct {
	Root    *Entry
	compare Comparator
	size    int
}

type color bool

const (
	black, red color = true, false
)

type Entry struct {
	K      interface{}
	V      interface{}
	left   *Entry
	right  *Entry
	parent *Entry
	color  color
}

func (t *RBTree) put(k, v interface{}) interface{} {
	if k == nil {
		return nil
	}

	root := t.Root
	// empty map
	if root == nil {
		t.compare(k, k)
		t.Root = &Entry{K: k, V: v, color: black}
		t.size++
		return nil
	}

	//e.fixAfterInsertion(root)
	t.size++
	return nil
}

func (e *Entry) fixAfterInsertion(root *Entry) {
	e.color = red
	for e.parent != nil && e != root && e.parent.color == red {
		if e.parent == e.doubleParentLeft() {
			right := e.doubleParentRight()
			if right.color == red {

			}
		} else {

		}
	}
	root.color = black
}

func (e *Entry) doubleParent() *Entry {
	if e.parent == nil {
		return nil
	}
	if e.parent.parent == nil {
		return nil
	}
	return e.parent.parent
}

func (e *Entry) doubleParentLeft() *Entry {
	if p := e.doubleParent(); p != nil {
		return p.left
	}
	return nil
}

func (e *Entry) doubleParentRight() *Entry {
	if p := e.doubleParent(); p != nil {
		return p.right
	}
	return nil
}

// set a new value, return old value
func (e *Entry) setValue(v interface{}) interface{} {
	oldValue := e.V
	e.V = v
	return oldValue
}
