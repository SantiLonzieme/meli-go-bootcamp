package calculadora

import "sort"

func Ordenamiento(numbers ...int) []int {

	sort.Ints(numbers)

	return numbers
}
