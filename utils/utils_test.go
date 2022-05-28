package utils

import (
	"fmt"
	"testing"
)

func TestIn(t *testing.T) {
	var arr = []int{2, 4, 6, 8}
	fmt.Printf("TestIn%+v\n", In(4, arr))
	fmt.Printf("TestIn%+v\n", In(5, arr))
}
