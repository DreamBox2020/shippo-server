package dao

import (
	"fmt"
	"testing"
)

func TestDaoGetPassport(t *testing.T) {
	d := New()
	p, err := d.GetPassport("ddd")
	if err != nil {
		panic(err)
	}
	fmt.Printf("TestDaoGetPassport:%+v\n", p)
}
