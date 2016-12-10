package maths

func SafeSum(x, y float64) (sum float64, ok bool) {
	sum = x + y
	ok = !(sum < x || sum < y)
	return
}

func (v *Var) SafeSum() (sum float64, ok bool) {
	for _, x := range v.Data {
		if sum, ok = SafeSum(sum, x); !ok {
			return
		}
	}
	return
}

func (v *Var) Sum() float64 {
	var sum float64
	for _, x := range v.Data {
		sum += x
	}
	return sum
}
