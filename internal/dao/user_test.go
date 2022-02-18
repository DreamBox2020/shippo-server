package dao

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestDaoUserFindByPhone(t *testing.T) {
	d := New()

	rand.Seed(time.Now().UnixNano())
	phone := "12345" + strconv.Itoa(rand.Intn(899999)+100000)

	u, err := d.Group.User.UserCreate(phone)
	if err != nil {
		panic(err)
	}

	u, err = d.Group.User.UserFindByPhone(phone)
	if err != nil {
		panic(err)
	}

	fmt.Printf("TestDaoUserFindByPhone:%+v\n", u)

}

func TestDaoUserCreate(t *testing.T) {
	d := New()

	rand.Seed(time.Now().UnixNano())
	phone := "12345" + strconv.Itoa(rand.Intn(899999)+100000)

	u, err := d.Group.User.UserCreate(phone)
	if err != nil {
		panic(err)
	}

	fmt.Printf("TestDaoUserCreate:%+v\n", u)
}
