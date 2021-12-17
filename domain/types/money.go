package types

type Money int

func (m Money) ToInt() int {
	return int(m)
}

func (m Money) ToFloat() float64 {
	return float64(m) / 100
}
