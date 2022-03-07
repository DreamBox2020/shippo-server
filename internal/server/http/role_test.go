package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"shippo-server/internal/model"
	"shippo-server/utils"
	"shippo-server/utils/box"
	"strings"
	"testing"
)

func TestRoleServer_RoleFind(t *testing.T) {
	if err := utils.ReadConfigFromFile("configs/server.json", &serverConf); err != nil {
		panic(err)
	}

	url := "http://127.0.0.1" + serverConf.Addr + "/role/find"

	data := new(struct {
		Id uint `json:"id"`
	})
	data.Id = 1

	resource, _ := json.Marshal(data)

	request, _ := json.Marshal(box.Request{
		Passport: "",
		Session:  "",
		Resource: string(resource),
		Sign:     "",
		Other:    nil,
	})

	fmt.Printf("request\n%+v\n", string(request))

	resp, err := http.Post(url, "application/json", strings.NewReader(string(request)))
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("response\n%+v\n", string(body))

	r := box.Response{}
	json.Unmarshal(body, &r)

	fmt.Printf("r\n%+v\n", r)

	rr := make([]model.PermissionAccess, 10)
	json.Unmarshal([]byte(r.Resource), &rr)

	fmt.Printf("rr\n%+v\n", rr)

	for i, v := range rr {
		fmt.Printf("arr[%+v]--->\n%+v\n", i, v)
	}

}
