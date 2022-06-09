package dao

import (
	"fmt"
	"testing"
)

func TestWxOffiaccountDao_FindAll(t *testing.T) {
	d := newTest()
	r, _ := d.Group.WxOffiaccount.FindAll()
	fmt.Printf("TestWxOffiaccountDao_FindAll:%+v\n", r)

}

func TestWxOffiaccountDao_Find(t *testing.T) {
	d := newTest()
	r, _ := d.Group.WxOffiaccount.Find(1)
	fmt.Printf("TestWxOffiaccountDao_Find:%+v\n", r)

}

func TestWxOffiaccountDao_FindByUsername(t *testing.T) {
	d := newTest()
	r, _ := d.Group.WxOffiaccount.Find(1)
	r, _ = d.Group.WxOffiaccount.FindByUsername(r.Username)
	fmt.Printf("TestWxOffiaccountDao_FindByUsername:%+v\n", r)

}
