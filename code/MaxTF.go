package code

func MaxPointerPointerInt(pvalues ...**int) (int, error) {
	values := make([]int, 0)
	for _,value := range pvalues { values = append(values, **value)}


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

func MaxPointerPointerInt16(pvalues ...**int16) (int16, error) {
	values := make([]int16, 0)
	for _,value := range pvalues { values = append(values, **value)}


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

func MaxPointerPointerInt32(pvalues ...**int32) (int32, error) {
	values := make([]int32, 0)
	for _,value := range pvalues { values = append(values, **value)}


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

func MaxPointerPointerUint(pvalues ...**uint) (uint, error) {
	values := make([]uint, 0)
	for _,value := range pvalues { values = append(values, **value)}


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

func MaxPointerPointerUint64(pvalues ...**uint64) (uint64, error) {
	values := make([]uint64, 0)
	for _,value := range pvalues { values = append(values, **value)}


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

func MaxPointerPointerFloat64(pvalues ...**float64) (float64, error) {
	values := make([]float64, 0)
	for _,value := range pvalues { values = append(values, **value)}


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

func MaxPointerPointerUint16(pvalues ...**uint16) (uint16, error) {
	values := make([]uint16, 0)
	for _,value := range pvalues { values = append(values, **value)}


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

func MaxPointerPointerUint8(pvalues ...**uint8) (uint8, error) {
	values := make([]uint8, 0)
	for _,value := range pvalues { values = append(values, **value)}


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

func MaxPointerPointerString(pvalues ...**string) (string, error) {
	values := make([]string, 0)
	for _,value := range pvalues { values = append(values, **value)}


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

func MaxPointerPointerInt8(pvalues ...**int8) (int8, error) {
	values := make([]int8, 0)
	for _,value := range pvalues { values = append(values, **value)}


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

func MaxPointerPointerInt64(pvalues ...**int64) (int64, error) {
	values := make([]int64, 0)
	for _,value := range pvalues { values = append(values, **value)}


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

func MaxPointerPointerUint32(pvalues ...**uint32) (uint32, error) {
	values := make([]uint32, 0)
	for _,value := range pvalues { values = append(values, **value)}


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

func MaxPointerPointerFloat32(pvalues ...**float32) (float32, error) {
	values := make([]float32, 0)
	for _,value := range pvalues { values = append(values, **value)}


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


func MaxInt8(values ...int8) (int8, error) {

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

func MaxUint16(values ...uint16) (uint16, error) {

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

func MaxUint32(values ...uint32) (uint32, error) {

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

func MaxFloat64(values ...float64) (float64, error) {

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

func MaxInt(values ...int) (int, error) {

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

func MaxInt32(values ...int32) (int32, error) {

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

func MaxUint(values ...uint) (uint, error) {

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

func MaxUint8(values ...uint8) (uint8, error) {

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

func MaxUint64(values ...uint64) (uint64, error) {

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

func MaxString(values ...string) (string, error) {

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

func MaxInt16(values ...int16) (int16, error) {

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

func MaxInt64(values ...int64) (int64, error) {

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

func MaxFloat32(values ...float32) (float32, error) {

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

func switchMax(values interface{}, second interface{})  (int, error) {
	switch getvalue := values.(type) {
	case int:
		return get(getvalue, second)
	}

panic("can't find type")
}


func get(first, second int) (int, error) {
	return 1, nil
}
