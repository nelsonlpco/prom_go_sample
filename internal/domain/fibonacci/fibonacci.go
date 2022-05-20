package fibonacci

type Fibonacci struct {
	cache map[uint64]uint64
}

func New() *Fibonacci {
	return &Fibonacci{cache: make(map[uint64]uint64)}
}

func (f *Fibonacci) Calc(n uint64) uint64 {
	if n < 2 {
		return n
	}

	if v, ok := f.cache[n]; ok {
		return v
	}

	f.cache[n] = f.Calc(n-2) + f.Calc(n-1)

	return f.cache[n]
}
