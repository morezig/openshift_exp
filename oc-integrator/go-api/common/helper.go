package common

import (
	"crypto/md5"
	"encoding/hex"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-ini/ini"
)

var configPath string
var DefaultConfPath = os.Getenv("APP_CONFIG") + "/app.conf"

func init() {
	SetConfigPath(DefaultConfPath)
}

func SetConfigPath(path string) {
	configPath = path
}

func GetConfig(section string, key string) *ini.Key {
	cfg, _ := ini.InsensitiveLoad(configPath)

	v, _ := cfg.Section(section).GetKey(key)
	return v
}

func GetServerPort() string {
	port := GetConfig("system", "httpport").String()
	return port
}

func Md5(source string) string {
	md5h := md5.New()
	md5h.Write([]byte(source))
	return hex.EncodeToString(md5h.Sum(nil))
}

func ResponseJsonString(c *gin.Context, j []byte) {
	w := c.Writer
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = []string{"application/json; charset=utf-8"}
	}
	w.Write(j)
}

// func bindForm2Model(form map[string]interface{}, model *models.User) {
// 	for k, v := range form {
// 		fmt.Printf("key[%s] value[%s]\n", k, v)
// 		switch v.(type) {
// 		case int:
// 			fmt.Println("int")
// 		case []interface{}:
// 			fmt.Println("[]interface {}")
// 		default:
// 			fmt.Println(v)
// 		}
// 	}
// }
