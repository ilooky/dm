package test

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
	"xorm.io/core"
)

func TestParse(t *testing.T) {
	dataSourceName := fmt.Sprintf("dm://%s:%s@%s:%s?characterEncoding=utf-8&useSSL=false&schema=%s",
		"Username",
		"Password",
		"Host",
		"Port",
		"Username",
	)
	dsnPattern := regexp.MustCompile(
		`^(?:(?P<user>.*?)(?::(?P<passwd>.*))?@)?` + // [user[:password]@]
			`(?:(?P<net>[^\(]*)(?:\((?P<addr>[^\)]*)\))?)?` + // [net[(addr)]]
			`\/(?P<dbname>.*?)` + // /dbname
			`(?:\?(?P<params>[^\?]*))?$`) // [?param1=value1&paramN=valueN]
	matches := dsnPattern.FindStringSubmatch(dataSourceName)
	// tlsConfigRegister := make(map[string]*tls.Config)
	names := dsnPattern.SubexpNames()

	uri := &core.Uri{DbType: core.MYSQL}

	for i, match := range matches {
		switch names[i] {
		case "dbname":
			fmt.Println(match)
			uri.DbName = match
		case "params":
			if len(match) > 0 {
				kvs := strings.Split(match, "&")
				for _, kv := range kvs {
					splits := strings.Split(kv, "=")
					if len(splits) == 2 {
						switch splits[0] {
						case "charset":
							uri.Charset = splits[1]
						}
					}
				}
			}

		}
	}
}
