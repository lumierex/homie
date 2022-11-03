package qrcode

import (
	"log"
	"net/url"
	"os"

	"github.com/skip2/go-qrcode"
)

func QRCodeGeneate(content, name string, size int, path string) {
	// 默认
	fileName := name
	u, err := url.Parse(content)
	if err != nil {
		log.Panic(err)
	}
	// 默认值不存在
	if len(fileName) == 0 {
		fileName = "./" + u.Hostname()
	}

	// 格式化后统一加后缀
	fileName = fileName + ".png"

	// 如果有目录
	fileName = path + fileName

	// TODO 目录是否存在，如果不存在，创建新目录
	_, exist := PathExist(path)
	if !exist {
		os.MkdirAll(path, os.ModePerm)
	}

	log.Println("generate qrcode start...", content, fileName)
	qrcode.WriteFile(content, qrcode.Medium, size, fileName)
	log.Println("generate qrcode finish...")

}

type QRCodeMetaInfo struct {
	name string
	url  string
}

// 批量
func QRCodeBatchGenerate(contents []QRCodeMetaInfo, size int, path string) {
	for _, v := range contents {
		log.Println(v.url, v.name)
		QRCodeGeneate(v.url, v.name, 500, path)
	}
}

// 判断路径是否存在
func PathExist(path string) (os.FileInfo, bool) {
	f, err := os.Stat(path)
	return f, err == nil || os.IsExist(err)
}
