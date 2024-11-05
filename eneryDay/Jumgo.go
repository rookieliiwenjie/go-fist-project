package eneryDay

import (
	"math"
)

type Jumgo struct{}

func (j Jumgo) GetSmallestString(c int) bool {
	for i := 0; i*i <= c; i++ {
		b := math.Sqrt(float64(c - i*i))
		if b == float64(int(b)) {
			return true
		}
	}
	return false
}
