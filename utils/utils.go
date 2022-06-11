package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"github.com/google/uuid"
	"io/ioutil"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func ReadFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return ioutil.ReadAll(file)
}

func PhoneMasking(s string) string {
	if len(s) < 11 {
		return s
	}
	return s[:3] + "******" + s[9:]
}

func QQMasking(s string) string {
	if len(s) < 5 {
		return s
	}
	return s[:1] + "******" + s[len(s)-2:]
}

func QQEmailMasking(s string) string {
	// xxxxx@qq.com
	if len(s) < 12 {
		return s
	}
	return s[:1] + "******" + s[len(s)-9:]
}

func FormatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func ParseTime(s string) (time.Time, error) {
	return time.Parse(s, "2006-01-02 15:04:05")
}

func GenerateCaptcha() string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(rand.Intn(899999) + 100000)
}

func GenerateToken() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}

// IsExist 判断文件或文件夹是否存在
func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		return false
	}
	return true
}

func In(val interface{}, arr interface{}) bool {
	arrValue := reflect.ValueOf(arr)
	switch reflect.TypeOf(arr).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < arrValue.Len(); i++ {
			if arrValue.Index(i).Interface() == val {
				return true
			}
		}
	case reflect.Map:
		if arrValue.MapIndex(reflect.ValueOf(val)).IsValid() {
			return true
		}
	}
	return false
}

// SHA1 sha1
func SHA1(s string) string {
	o := sha1.New()
	o.Write([]byte(s))
	return hex.EncodeToString(o.Sum(nil))
}

func DetectContentType(header *multipart.FileHeader) string {
	file, _ := header.Open()
	defer file.Close()
	buffer := make([]byte, 512)
	if _, err := file.Read(buffer); err != nil {
		return "application/octet-stream"

	}
	return http.DetectContentType(buffer)
}

func HttpGet(url string) (bytes []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func HttpGetString(url string) (str string, err error) {
	bytes, err := HttpGet(url)
	if err != nil {
		return
	}

	return string(bytes), nil
}

func HttpGetJSON(url string, obj interface{}) (err error) {
	bytes, err := HttpGet(url)
	if err != nil {
		return
	}

	err = json.Unmarshal(bytes, obj)
	return
}
