package html

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	str := "1234567890"
	fmt.Printf("TestSlice：%+v\n", Slice(str, 2, 4))
	fmt.Printf("TestSlice：%+v\n", Slice(str, 2, len(str)))
}

func TestCharAt(t *testing.T) {
	fmt.Printf("TestCharAt：%+v\n", CharAt("1234567890", 2))
}

func TestIndexOf(t *testing.T) {
	fmt.Printf("TestIndexOf：%+v\n", IndexOf("1234567890", "89", 0))
	fmt.Printf("TestIndexOf：%+v\n", IndexOf("1234567890", "89", 2))
}

func TestParseFragment(t *testing.T) {
	ele, _ := ParseFragment(`<div class="aaa" id="box" demo="<>">
        aaa
        <h1 class="title">666</h1>
        bbb
        <img id="img" src="0.jpg" />
      </div>`)

	res, _ := json.Marshal(ele)
	fmt.Printf("TestParseFragment：%+v\n", string(res))

}
