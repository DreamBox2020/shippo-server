package dao

import (
	"fmt"
	"math/rand"
	"shippo-server/internal/model"
	"strconv"
	"testing"
	"time"
)

func TestUserDao_UserFindByPhone(t *testing.T) {
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

func TestUserDao_UserCreate(t *testing.T) {
	d := New()

	rand.Seed(time.Now().UnixNano())
	phone := "12345" + strconv.Itoa(rand.Intn(899999)+100000)

	u, err := d.Group.User.UserCreate(phone)
	if err != nil {
		panic(err)
	}

	fmt.Printf("TestDaoUserCreate:%+v\n", u)
}

func TestUserDao_FindAll(t *testing.T) {
	d := New()
	res, err := d.Group.User.FindAll(model.UserFindAllReq{
		Pagination: model.Pagination{
			Current:  2,
			PageSize: 100,
		},
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("TestUserDao_FindAll:%+v\n", res.Total)
	for _, v := range res.Items {
		fmt.Printf("%+v\n", v)
	}
}
