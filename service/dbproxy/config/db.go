package config

import "fmt"

var (
	MySQLSource = "root:root@tcp(192.168.88.131:3306)/fileserver?charset=utf8"
)

func UpdateDBHost(host string) {
	MySQLSource = fmt.Sprintf("root:root@tcp(%s)/fileserver?charset=utf8", host)
}
