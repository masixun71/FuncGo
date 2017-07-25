package Func


func astMax(values ...int) {
	len := len(values)

	max := values[0]
	for i := 1; i < len; i++ {
		tmpvalue := values[i]

		if max < tmpvalue {
			max = tmpvalue
		}
	}
}