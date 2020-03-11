package exemplos

import (
	"fmt"
	"math"
)

func constante() {
	const Pi = 3.14

	// Pi = 4.93 -> cannot assign to Pi because is Const

	fmt.Println(Pi)

	fmt.Println(math.Max(1, 5))
}
