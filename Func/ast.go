package Func


func CompareT(first, second T) int {
	if first > second {
		return 1
	}else if first < second{
		return -1
	} else {
		return 0
	}
}

func MaxTF(values ...T) (T, error) {

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




func StringTInterface(first TInterface) string {
	res, ok := first.(string)
	if ok {
		return string(res)
	}

	return ""
}
