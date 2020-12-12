package maps

type imap interface {
	Size() int
	Get(key interface{}) (value interface{}, found bool)
	Put(key, value interface{})
	Remove(key interface{}) (ok bool, err error)
	Empty() bool
	Clear()
	Keys() []interface{}
	Values() interface{}
	ContainsKey(key interface{}) bool
	ContainsValue(value interface{}) bool
	Each(func(key interface{}, value interface{}))
	PutIfAbsent(key, value interface{}) interface{}
}
