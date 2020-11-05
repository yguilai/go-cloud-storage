package main

import (
	"fmt"
	"github.com/yguilai/go-cloud-storage/store/kodo"
)

func main() {
	// Mysql分库分表.md
	//d, _ := ioutil.ReadFile("./tmp/Mysql分库分表.md")
	//ret, err := kodo.UploadToQiniu("Mysql分库分表.md", bytes.NewReader(d), int64(len(d)))
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//
	//fmt.Println(ret.Key, ret.Hash, len(ret.Hash))
	//fmt.Println(kodo.DownloadUrl(ret.Key))

	fmt.Println(kodo.GetEtag("./tmp/Mysql分库分表.md"))
}
