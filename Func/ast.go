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


func StringTInterface(first TInterface) string {
	res, ok := first.(string)
	if ok {
		return string(res)
	}

	return ""
}
