package maps

import (
	"github.com/shopspring/decimal"
	"testing"
)

func TestTreeMap_Put(t *testing.T) {
	treeMap := NewDecimalTreeMap()
	treeMap.Put(decimal.RequireFromString("11"), "11")
	treeMap.Put(decimal.RequireFromString("22"), "22")
	treeMap.Put(decimal.RequireFromString("33"), "33")
	treeMap.Put(decimal.RequireFromString("44"), "44")
	treeMap.Put(decimal.RequireFromString("55"), "55")
	treeMap.Put(decimal.RequireFromString("66"), "66")
	treeMap.Put(decimal.RequireFromString("77"), "77")
	treeMap.Put(decimal.RequireFromString("88"), "88")
	treeMap.Put(decimal.RequireFromString("99"), "99")
	treeMap.Put(decimal.RequireFromString("1"), "1")
	treeMap.Put(decimal.RequireFromString("2"), "2")
	treeMap.Put(decimal.RequireFromString("3"), "3")
	treeMap.Put(decimal.RequireFromString("4"), "4")
	treeMap.Put(decimal.RequireFromString("5"), "5")
	treeMap.Put(decimal.RequireFromString("6"), "6")
	treeMap.Put(decimal.RequireFromString("6"), "6dup")
	treeMap.Put(decimal.RequireFromString("7"), "7")
	treeMap.Put(decimal.RequireFromString("8"), "8")
	treeMap.Put(decimal.RequireFromString("9"), "9")
	treeMap.Put(decimal.RequireFromString("111"), "111")
	treeMap.Put(decimal.RequireFromString("222"), "222")
	treeMap.Put(decimal.RequireFromString("333"), "333")
	treeMap.Put(decimal.RequireFromString("444"), "444")

	t.Logf("%#v", treeMap)
}
