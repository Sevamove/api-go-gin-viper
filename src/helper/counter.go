package helper

type Counter struct {
	counter uint64
}

func (cnt *Counter) Increment() uint64 {
	cnt.counter++
	return cnt.counter
}

func (cnt *Counter) Decrease() uint64 {
	cnt.counter--
	return cnt.counter
}

func (cnt *Counter) Current() uint64 {
	return cnt.counter
}
