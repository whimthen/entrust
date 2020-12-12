package maps

import (
	"github.com/shopspring/decimal"
	"strings"
)

type Comparator func(v1, v2 interface{}) int

func StringComparator(v1, v2 interface{}) int {
	return strings.Compare(v1.(string), v2.(string))
}

func DecimalComparator(v1, v2 interface{}) int {
	return v1.(decimal.Decimal).Cmp(v2.(decimal.Decimal))
}
