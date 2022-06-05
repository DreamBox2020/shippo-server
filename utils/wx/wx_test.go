package wx

import (
	"fmt"
	"testing"
)

func TestNewArticle(t *testing.T) {

	html := `

        <meta property="og:title" content="文章标题" />
        <meta property="og:url" content="文章链接==&amp;tempkey=AAA" />
        <meta property="og:image" content="文章封面" />
        <meta property="og:description" content="文章简介" />
        <meta property="og:site_name" content="公众平台" />
        <meta property="og:type" content="article" />
        <meta property="og:article:author" content="文章作者" />
        var is_follow = "";
        var nickname = "公众号昵称";
        var ct = "1234567";
        var user_name = "公众号用户名";
        var version = "";
        var cdn_url_1_1 = "1:1比例的封面图";
        var cdn_url_235_1 = "1:1比例的封面图";
        <meta name="description" content="文章简介" />
        <meta name="author" content="文章作者" />

`

	a := &Article{html}

	fmt.Printf("TestNewArticle->Title:%+v\n", a.Title())
	fmt.Printf("TestNewArticle->URL:%+v\n", a.URL())
	fmt.Printf("TestNewArticle->Description:%+v\n", a.Description())
	fmt.Printf("TestNewArticle->Author:%+v\n", a.Author())
	fmt.Printf("TestNewArticle->IsTempURL:%+v\n", a.IsTempURL())
	fmt.Printf("TestNewArticle->Username:%+v\n", a.Username())
	fmt.Printf("TestNewArticle->Nickname:%+v\n", a.Nickname())
	fmt.Printf("TestNewArticle->Image1:%+v\n", a.Image1())
	fmt.Printf("TestNewArticle->Image2:%+v\n", a.Image2())

}
