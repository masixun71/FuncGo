package lib



type Comparer interface {

	compareInt() int
	compareInt8() int
	compareInt16() int
	compareInt32() int
	compareInt64() int
	compareUint() int
	compareUint8() int
	compareUint16() int
	compareUint32() int
	compareUint64() int
	compareFloat32() int
	compareFloat64() int
	compareString() int
}

type Compare struct {
	first interface{}
	second interface{}
}

func (c *Compare) compareInt() int {
	tFirst,_ := c.first.(int)
	tSecond,_ := c.first.(int)

	if tFirst > tSecond {
		return -1
	}else if tFirst < tSecond{
		return 1
	} else {
		return 0
	}
}

func (c *Compare) compareInt8() int {

	tFirst,_ := c.first.(int8)
	tSecond,_ := c.first.(int8)

	if tFirst > tSecond {
		return -1
	}else if tFirst < tSecond{
		return 1
	} else {
		return 0
	}
}

func (c *Compare) compareInt16() int {

	tFirst,_ := c.first.(int16)
	tSecond,_ := c.first.(int16)

	if tFirst > tSecond {
		return -1
	}else if tFirst < tSecond{
		return 1
	} else {
		return 0
	}
}

func (c *Compare) compareInt32() int {

	tFirst,_ := c.first.(int32)
	tSecond,_ := c.first.(int32)

	if tFirst > tSecond {
		return -1
	}else if tFirst < tSecond{
		return 1
	} else {
		return 0
	}
}

func (c *Compare) compareInt64() int {

	tFirst,_ := c.first.(int64)
	tSecond,_ := c.first.(int64)

	if tFirst > tSecond {
		return -1
	}else if tFirst < tSecond{
		return 1
	} else {
		return 0
	}
}

func (c *Compare) compareUint() int {
	tFirst,_ := c.first.(uint)
	tSecond,_ := c.first.(uint)

	if tFirst > tSecond {
		return -1
	}else if tFirst < tSecond{
		return 1
	} else {
		return 0
	}
}

func (c *Compare) compareUint8() int {

	tFirst,_ := c.first.(uint8)
	tSecond,_ := c.first.(uint8)

	if tFirst > tSecond {
		return -1
	}else if tFirst < tSecond{
		return 1
	} else {
		return 0
	}
}

func (c *Compare) compareUint16() int {

	tFirst,_ := c.first.(uint16)
	tSecond,_ := c.first.(uint16)

	if tFirst > tSecond {
		return -1
	}else if tFirst < tSecond{
		return 1
	} else {
		return 0
	}
}

func (c *Compare) compareUint32() int {
	tFirst,_ := c.first.(uint32)
	tSecond,_ := c.first.(uint32)

	if tFirst > tSecond {
		return -1
	}else if tFirst < tSecond{
		return 1
	} else {
		return 0
	}
}

func (c *Compare) compareUint64() int {
	tFirst,_ := c.first.(uint64)
	tSecond,_ := c.first.(uint64)

	if tFirst > tSecond {
		return -1
	}else if tFirst < tSecond{
		return 1
	} else {
		return 0
	}
}

func (c *Compare) compareFloat32() int {
	tFirst,_ := c.first.(float32)
	tSecond,_ := c.first.(float32)

	if tFirst > tSecond {
		return -1
	}else if tFirst < tSecond{
		return 1
	} else {
		return 0
	}
}

func (c *Compare) compareFloat64() int {
	tFirst,_ := c.first.(float64)
	tSecond,_ := c.first.(float64)

	if tFirst > tSecond {
		return -1
	}else if tFirst < tSecond{
		return 1
	} else {
		return 0
	}
}

func (c *Compare) compareString() int {
	tFirst,_ := c.first.(string)
	tSecond,_ := c.first.(string)

	if tFirst > tSecond {
		return -1
	}else if tFirst < tSecond{
		return 1
	} else {
		return 0
	}
}

func NewCompare(first, second interface{}) Comparer{
	return &Compare{first:first, second:second}
}



