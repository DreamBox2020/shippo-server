package dao

import (
	"fmt"
	"testing"
)

func TestPermissionPolicyDao_PermissionPolicyFindAll(t *testing.T) {
	d := New()
	res, _ := d.Group.PermissionPolicy.PermissionPolicyFindAll()
	for _, v := range res {
		fmt.Printf("PermissionPolicyFindAll:%+v\n", v)
	}
}
