package html

import (
	"fmt"
	"regexp"
	"strings"
)

type Attr struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Node struct {
	NodeName      string   `json:"nodeName"`
	NodeValue     string   `json:"nodeValue"`
	ChildNodes    []*Node  `json:"childNodes"`
	parentElement *Element `json:"-"`
	parentNode    *Node    `json:"-"`
}

type Element struct {
	*Node
	TagName    string     `json:"tagName"`
	Attributes []*Attr    `json:"attributes"`
	Children   []*Element `json:"children"`
}

func Slice(str string, start int, end int) string {
	return str[start:end]
}

func CharAt(str string, pos int) string {
	return str[pos : pos+1]
}

func IndexOf(str, searchValue string, fromIndex int) int {
	str = Slice(str, fromIndex, len(str))
	return fromIndex + strings.Index(str, searchValue)
}

func replace(str string, repl string) (string, error) {
	reg, err := regexp.Compile("[\t\n\f\r]")
	if err != nil {
		return "", err
	}
	return reg.ReplaceAllString(str, repl), nil
}

func SpaceIndex(str string) (int, error) {
	reg, err := regexp.Compile("[\t\n\f\r ]")
	if err != nil {
		return -1, err
	}
	match := reg.FindStringIndex(str)
	if match == nil {
		return -1, nil
	} else {
		return match[0], nil
	}
}

func GetTagName(html string) (string, error) {
	i, err := SpaceIndex(html)
	if err != nil {
		return "", err
	}
	tagName := ""
	if i == -1 {
		tagName = Slice(html, 1, len(html)-1)
	} else {
		tagName = Slice(html, 1, i+1)
	}
	fmt.Printf("GetTagName:%+v\n", tagName)
	tagName = strings.ToUpper(strings.TrimSpace(tagName))
	if Slice(tagName, 0, 1) == "/" {
		tagName = Slice(tagName, 1, len(tagName))
	}
	if Slice(tagName, len(tagName)-1, len(tagName)) == "/" {
		tagName = Slice(tagName, 0, len(tagName)-1)
	}
	return tagName, nil
}

func IsClosing(html string) bool {
	return Slice(html, 0, 2) == "</"
}

func GetAttrs(html string) (string, bool, error) {
	i, err := SpaceIndex(html)
	if err != nil {
		return "", false, err
	}
	if i == -1 {
		return "", Slice(html, len(html)-2, len(html)-1) == "/", nil
	}

	html = strings.TrimSpace(Slice(html, i+1, len(html)-1))
	isClosing := Slice(html, len(html)-1, len(html)) == "/"
	if isClosing {
		html = strings.TrimSpace(Slice(html, 0, len(html)-1))
	}
	return html, isClosing, nil
}

func ParseFragment(html string) (element *Element, err error) {

	stack := make([]*Element, 0)

	chatList := strings.Split(html, "")
	tagStart := -1
	quoteStart := ""
	lastPos := 0
	rethtml := ""

Loop:
	for currentPos, c := range chatList {
		fmt.Printf("==============================>>> tagStart:%+v,lastPos:%+v,currentPos:%+v,c:%+v\n", tagStart, lastPos, currentPos, c)
		if tagStart == -1 {
			if c == "<" {
				// 记录标签开始位置
				tagStart = currentPos
				fmt.Printf("<<<============================== tagStart:%+v,lastPos:%+v\n", tagStart, lastPos)
				continue
			}
		} else {
			if quoteStart == "" {
				if c == "<" {
					if lastPos != 0 && currentPos != 0 {
						currentText := Slice(html, lastPos, currentPos)
						rethtml += currentText
						fmt.Printf("escapeHtml:%+v,lastPos:%+v,currentPos:%+v\n", currentText, lastPos, currentPos)

						// 创建节点
						if len(stack) > 0 {
							lastElement := stack[len(stack)-1]
							node := &Node{
								NodeName:      "#text",
								NodeValue:     currentText,
								ChildNodes:    nil,
								parentElement: nil,
								parentNode:    nil,
							}
							lastElement.ChildNodes = append(lastElement.ChildNodes, node)
						}

					}

					tagStart = currentPos
					lastPos = currentPos
					fmt.Printf("<<<============================== tagStart:%+v,lastPos:%+v\n", tagStart, lastPos)
					continue
				}

				if c == ">" {
					currentText := Slice(html, lastPos, tagStart)
					rethtml += currentText
					fmt.Printf("escapeHtml：%+v,lastPos:%+v,tagStart:%+v\n", currentText, lastPos, tagStart)

					// 创建节点
					if len(stack) > 0 {
						lastElement := stack[len(stack)-1]
						node := &Node{
							NodeName:      "#text",
							NodeValue:     currentText,
							ChildNodes:    nil,
							parentElement: nil,
							parentNode:    nil,
						}
						lastElement.ChildNodes = append(lastElement.ChildNodes, node)
					}

					currentHtml := Slice(html, tagStart, currentPos+1)
					fmt.Printf("onTag：%+v\n", currentHtml)

					currentTagName, _ := GetTagName(currentHtml)
					fmt.Printf("currentTagName:%+v\n", currentTagName)
					//rethtml += onTag(
					//	tagStart,
					//	rethtml.length,
					//	currentTagName,
					//	currentHtml,
					//	isClosing(currentHtml)
					//)
					rethtml += currentHtml

					// 不是结尾
					if !IsClosing(currentHtml) {
						// 创建元素
						ele := &Element{
							TagName:    currentTagName,
							Attributes: make([]*Attr, 0),
							Children:   make([]*Element, 0),
							Node: &Node{
								NodeName:      currentTagName,
								NodeValue:     "",
								ChildNodes:    make([]*Node, 0),
								parentElement: nil,
								parentNode:    nil,
							},
						}

						// 如果是root
						if element == nil {
							element = ele
						} else {
							if len(stack) > 0 {
								lastElement := stack[len(stack)-1]
								lastElement.Children = append(lastElement.Children, ele)
								lastElement.ChildNodes = append(lastElement.ChildNodes, ele.Node)
							}

						}

						attr, closing, _ := GetAttrs(currentHtml)
						fmt.Printf("GetAttrs:%+v,closing:%+v\n", attr, closing)

						parseAttr(attr)
						// 如果不是自闭合标签，就入栈
						if !closing && currentTagName != "IMG" {
							stack = append(stack, ele)
						}

					} else {
						// 如果遇到闭合标签，自动出栈
						if len(stack) > 0 {
							stack = stack[0 : len(stack)-1]
						}
					}

					lastPos = currentPos + 1
					tagStart = -1
					fmt.Printf("<<<============================== tagStart:%+v,lastPos:%+v\n", tagStart, lastPos)
					continue
				}

				if c == "\"" || c == "'" {
					i := 1
					// 当前字符的前一位
					var ic = CharAt(html, currentPos-i)

					for strings.TrimSpace(ic) == "" || ic == "=" {
						// 如果是等号，那么就是当前字符 就是引号的开始，并记录引号类型
						if ic == "=" {
							quoteStart = c
							fmt.Printf("<<<============================== tagStart:%+v,lastPos:%+v\n", tagStart, lastPos)
							continue Loop
						}
						// 否则，ic继续 向前截取
						ic = CharAt(html, currentPos-(i+1))
					}
				}

			} else {
				if c == quoteStart {
					quoteStart = ""
					fmt.Printf("<<<============================== tagStart:%+v,lastPos:%+v\n", tagStart, lastPos)
					continue
				}
			}
		}
		fmt.Printf("<<<============================== tagStart:%+v,lastPos:%+v\n", tagStart, lastPos)
	}

	return
}

func findNextEqual(str string, i int) int {
	for ; i < len(str); i++ {
		c := CharAt(str, i)
		if c == " " {
			continue
		}
		if c == "=" {
			return i
		}
		return -1
	}
	return -1
}

func findBeforeEqual(str string, i int) int {
	for ; i < len(str); i-- {
		c := CharAt(str, i)
		if c == " " {
			continue
		}
		if c == "=" {
			return i
		}
		return -1
	}
	return -1
}

func isQuoteWrapString(text string) bool {
	if CharAt(text, 0) == "\"" && CharAt(text, len(text)-1) == "\"" {
		return true
	}
	if CharAt(text, 0) == "'" && CharAt(text, len(text)-1) == "'" {
		return true
	}
	return false
}

func stripQuoteWrap(text string) string {
	if isQuoteWrapString(text) {
		return Slice(text, 1, len(text)-1)
	} else {
		return text
	}
}

func parseAttr(html string) {
	html, _ = replace(html, " ")

	chatList := strings.Split(html, "")
	lastPos := 0
	tmpName := ""

	var addAttr = func(name string, value string) {
		name = strings.TrimSpace(name)
		//name = name.replace(REGEXP_ILLEGAL_ATTR_NAME, '').toLowerCase()
		if len(name) < 1 {
			return
		}
		//var ret = onAttr(name, value || '')
		//if (ret) retAttrs.push(ret)
		fmt.Printf("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ name:%+v,value:%+v\n", name, value)

	}

	for i, c := range chatList {
		fmt.Printf("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ i:%+v,c:%+v\n", i, c)

		var v string
		var j int
		if tmpName == "" && c == "=" {
			tmpName = Slice(html, lastPos, i)
			fmt.Printf("^^^^^^^^^^^^^^^^ tmpName:%+v,lastPos:%+v,i:%+v\n", tmpName, lastPos, i)
			lastPos = i + 1
			continue
		}

		if tmpName != "" {
			if i == lastPos && (c == "\"" || c == "'") && CharAt(html, i-1) == "=" {
				j = IndexOf(html, c, i+1)
				fmt.Printf("^^^^^^^^^^^^^^^^ html:%+v,c:%+v,c:%+v,j:%+v\n", html, c, i+1, j)
				if j == -1 {
					break
				} else {
					fmt.Printf("^^^^^^^^^^^^^^^^ lastPos:%+v,j:%+v\n", lastPos+1, j)
					v = strings.TrimSpace(Slice(html, lastPos+1, j))
					addAttr(tmpName, v)
					tmpName = ""
					i = j
					lastPos = i + 1
					continue
				}
			}
		}

		if c == " " {
			if tmpName == "" {
				fmt.Printf("^^^^^^^^^^^^^^^^ 执行了11111")
				j = findNextEqual(html, i)
				if j == -1 {
					v = strings.TrimSpace(Slice(html, lastPos, i))
					addAttr(v, "")
					tmpName = ""
					lastPos = i + 1
					continue
				} else {
					i = j - 1
					continue
				}
			} else {
				fmt.Printf("^^^^^^^^^^^^^^^^ 执行了22222")
				j = findBeforeEqual(html, i-1)
				if j == -1 {
					v = strings.TrimSpace(Slice(html, lastPos, i))
					v = stripQuoteWrap(v)
					addAttr(tmpName, v)
					tmpName = ""
					lastPos = i + 1
					continue
				} else {
					continue
				}
			}
		}
	}

	return
}

func Serialize(element *Element) string {
	return ""
}
