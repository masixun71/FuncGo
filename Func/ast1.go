package Func


func CompareTF1(first , second T) int {

	if first > second {
		return 1
	}else if first < second{
		return -1
	} else {
		return 0
	}
}

//func CompareTFP(pfirst , psecond *T) int {
//	first := *pfirst
//	second := *psecond
//
//	if first > second {
//		return 1
//	}else if first < second{
//		return -1
//	} else {
//		return 0
//	}
//}
//
func MaxTFP(values ...T) (T, error) {
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
