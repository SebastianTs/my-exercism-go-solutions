package flatten

func Flatten(arr interface{}) []interface{} {
	out := make([]interface{}, 0)
	for _, list := range arr.([]interface{}) {
		switch list.(type) {
		case []interface{}:
			tmp := Flatten(list)
			out = append(out, tmp...)
		case int:
			out = append(out, list)
		}
	}
	return out
}
