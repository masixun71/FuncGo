package Func

func CompareTF(first, second T) int {
	if first > second {
		return 1
	} else if first < second {
		return -1
	} else {
		return 0
	}
}

func MaxTF(values ...T) (T, error) {

	//pvalues := make([]T, 0)
	//for _,value := range values {
	//	pvalues = append(pvalues, *value)
	//}

	len := len(values)

	max := values[0]
	for i := 1; i < len; i++ {
		tmpvalue := values[i]

		if max < tmpvalue {
			max = tmpvalue
		}
	}

	return max, nil
}

func StringTFInterface(first TInterface) string {
	res, ok := first.(string)
	if ok {
		return string(res)
	}

	return ""
}

func Switch(first, second T) int {

	switch first.(type) {
	case int:
		return CompareTF(first, second)

	}

	panic("can't find type")
}

func SwitchB(values ...T) (T, error) {

	switch values[0].(type) {
	case int:
		return MaxTF(values...)

	}

	panic("can't find type")
}
