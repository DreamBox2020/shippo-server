package wx

import (
	"regexp"
	"shippo-server/utils"
	"strings"
)

type Article struct {
	html string
}

func NewArticle(url string) (*Article, error) {
	bytes, err := utils.HttpGet(url)
	if err != nil {
		return nil, err
	}
	return &Article{html: string(bytes)}, nil
}

func (t *Article) find(reg *regexp.Regexp) string {
	arr := reg.FindStringSubmatch(t.html)
	if arr != nil && len(arr) > 1 {
		return arr[1]
	}
	return ""
}

func (t *Article) Title() string {
	return t.find(regexp.MustCompile(`<meta property="og:title" content="(.*?)"`))
}

func (t *Article) URL() string {
	return t.find(regexp.MustCompile(`<meta property="og:url" content="(.*?)"`))
}

func (t *Article) IsTempURL() bool {
	return strings.Contains(t.URL(), "tempKey")
}

func (t *Article) Description() string {
	return t.find(regexp.MustCompile(`<meta property="og:description" content="(.*?)"`))
}

func (t *Article) Author() string {
	return t.find(regexp.MustCompile(`<meta property="og:article:author" content="(.*?)"`))
}

func (t *Article) Username() string {
	return t.find(regexp.MustCompile(`var user_name = "(.*?)"`))
}

func (t *Article) Nickname() string {
	return t.find(regexp.MustCompile(`var nickname = "(.*?)"`))
}

// Image1 获取封面 1:1
func (t *Article) Image1() string {
	return t.find(regexp.MustCompile(`var cdn_url_1_1 = "(.*?)"`))
}

// Image2 获取封面 2.35:1
func (t *Article) Image2() string {
	return t.find(regexp.MustCompile(`var cdn_url_235_1 = "(.*?)"`))
}
