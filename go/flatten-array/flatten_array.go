package flatten

func Flatten(nested interface{}) []interface{} {
	output := []interface{}{}
	switch i := nested.(type) {
	case []interface{}:
		{
			for _, val := range i {
				output = append(output, Flatten(val)...)
			}
		}
	case interface{}:
		{
			output = append(output, i)
		}
	}
	return output
}
