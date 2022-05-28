package dao

import (
	"fmt"
	"testing"
)

func TestDaoRoleFindPermissionPolicy(t *testing.T) {
	d := newTest()

	p, err := d.Group.Role.RoleFindPermissionPolicy(1)
	if err != nil {
		panic(err)
	}

	fmt.Printf("RoleFindPermissionPolicy:%+v\n", p)
}

func TestDaoRoleFindPermissionAccess(t *testing.T) {
	d := newTest()

	p, err := d.Group.Role.RoleFindPermissionAccess(1)
	if err != nil {
		panic(err)
	}

	fmt.Printf("RoleFindPermissionAccess:%+v\n", p)
}
