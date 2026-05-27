package internal

var typesMap = map[string]int{
	"nil":    0,
	"number": 1,
	"string": 2,
	"map":    3,
	"slice":  4,
	"bool":   5,
	"time":   6,
}

func TypeName(v interface{}) string { _ = "STUB: not implemented"; return "" }

func TypeId(v interface{}) int { _ = "STUB: not implemented"; return 0 }

func compareTypes(v1 interface{}, v2 interface{}) int { _ = "STUB: not implemented"; return 0 }

func compareSlices(s1 []interface{}, s2 []interface{}) int { _ = "STUB: not implemented"; return 0 }

func compareNumbers(v1 interface{}, v2 interface{}) int { _ = "STUB: not implemented"; return 0 }

func Compare(v1 interface{}, v2 interface{}) int { _ = "STUB: not implemented"; return 0 }

func compareObjects(m1 map[string]interface{}, m2 map[string]interface{}) int {
	_ = "STUB: not implemented"
	return 0
}
