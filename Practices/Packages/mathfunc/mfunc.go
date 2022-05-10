package mathfunc

import (
	"fmt"
	"math"
)

func Sqroot(n int) (sqroot float64) {
	fmt.Printf("SQUARE ROOT OF %d \n", n)
	return math.Sqrt(float64(n))
}
