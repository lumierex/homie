package qrcode

import (
	"image"
	"image/draw"
	// "image/gif"
	"image/jpeg"
	// "image/png"
	"log"
	"net/url"
	"os"
	"path/filepath"
	// "strings"

	"github.com/nfnt/resize"
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

func QRCodeGenerateWithAvatar(url string) error {
	qrCode, err := qrcode.New(url, qrcode.Medium)
	if err != nil {
		log.Fatal(err)
		return err
	}
	bgImage := qrCode.Image(256)

	avatarFile, err := os.Open("./fs.png")
	
	if err != nil {
		return err
	}
	avatar, err := jpeg.Decode(avatarFile)
	if err != nil {
		return err
	}
	//修改图片的大小
	avatar = resize.Resize(40, 40, avatar, resize.Lanczos3)

	//得到背景图的大小
	b := bgImage.Bounds()
	//居中设置icon到二维码图片
	offset := image.Pt((b.Max.X-avatar.Bounds().Max.X)/2, (b.Max.Y-avatar.Bounds().Max.Y)/2)
	m := image.NewRGBA(b)
	draw.Draw(m, b, bgImage, image.Point{X: 0, Y: 0}, draw.Src)
	draw.Draw(m, avatar.Bounds().Add(offset), avatar, image.Point{X: 0, Y: 0}, draw.Over)

	err = SaveImage("./new.png", m)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil

}

func SaveImage(p string, src image.Image) error {
	f, err := os.OpenFile(p, os.O_SYNC|os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	// f.Write()//直接写入
    // defer f.Close()
    // return nil
	// defer f.Close()
	ext := filepath.Ext(p)
	log.Println(ext)
	err = jpeg.Encode(f, src, &jpeg.Options{Quality: 80})
	// log.Println("ext: ", ext)
	// err = png.Encode(f, src)

	// if strings.EqualFold(ext, ".jpg") || strings.EqualFold(ext, ".jpeg") {
	// 	err = jpeg.Encode(f, src, &jpeg.Options{Quality: 80})
	// } else if strings.EqualFold(ext, ".png") {
	// 	err = png.Encode(f, src)
	// } else if strings.EqualFold(ext, ".gif") {
	// 	err = gif.Encode(f, src, &gif.Options{NumColors: 256})
	// }
	return err
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
