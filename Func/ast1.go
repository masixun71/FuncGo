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
