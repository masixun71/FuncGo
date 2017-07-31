package Func

type ts struct {

}

func (t *ts) CompareT(first, second T) int {
	if first > second {
		return 1
	}else if first < second{
		return -1
	} else {
		return 0
	}
}