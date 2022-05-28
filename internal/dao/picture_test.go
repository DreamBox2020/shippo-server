package dao

import (
	"fmt"
	"testing"
)

func TestPictureDao_FindByUri(t *testing.T) {
	d := newTest()

	r, err := d.Group.Picture.FindByUri("/pic/")
	if err != nil {
		panic(err)
	}

	fmt.Printf("TestPictureDao_FindByUri:%+v\n", r)
}
