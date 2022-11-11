package qrcode

import (
	"bufio"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"net/url"
	"os"

	"github.com/llgcode/draw2d/draw2dimg"
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
	// qrCode.DisableBorder = false
	bgImage := qrCode.Image(500)

	avatarFile, err := os.Open("./log.png")

	if err != nil {
		return err
	}

	// 此处的decode、要跟保存图片的encode一致
	avatar, err := png.Decode(avatarFile)

	if err != nil {
		return err
	}
	//修改图片的大小
	avatar = resize.Resize(80, 80, avatar, resize.Lanczos3)

	pic2FramePadding := 20
	// transparentAvatar := image.NewRGBA(image.Rect(0, 0, avatar.Bounds().Dx()+pic2FramePadding, avatar.Bounds().Dy()+pic2FramePadding))
	transparentAvatar := image.NewRGBA(image.Rect(0, 0, avatar.Bounds().Dx()+pic2FramePadding, avatar.Bounds().Dy()+pic2FramePadding))
	white := color.RGBA{
        R: 255,
        G: 255,
        B: 255,
        A: 255,
    }
	for i := 0; i < transparentAvatar.Bounds().Size().X; i++ {
        for j := 0; j < transparentAvatar.Bounds().Size().Y; j++ {
            transparentAvatar.SetRGBA(i, j, white)
        }
    }

	lineToPic(transparentAvatar)

	//得到背景图的大小
	b := bgImage.Bounds()
	//居中设置icon到二维码图片
	offset1 := image.Pt((b.Max.X- transparentAvatar.Bounds().Max.X)/2, (b.Max.Y-transparentAvatar.Bounds().Max.Y)/2)
	offset2 := image.Pt((b.Max.X-transparentAvatar.Bounds().Max.X)/2+pic2FramePadding/2, (b.Max.Y-transparentAvatar.Bounds().Max.Y)/2+pic2FramePadding/2)
	m := image.NewRGBA(b)
	draw.Draw(m, b, bgImage, image.Point{X: 0, Y: 0}, draw.Src)
	draw.Draw(m, transparentAvatar.Bounds().Add(offset1), transparentAvatar, image.Point{X: 0, Y: 0}, draw.Over)
	draw.Draw(m, avatar.Bounds().Add(offset2), avatar, image.Point{X: 0, Y: 0}, draw.Over)

	err = SaveImage("./new.png", m)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil

}

func lineToPic(transparentImg *image.RGBA) {
	gc := draw2dimg.NewGraphicContext(transparentImg)
	gc.SetStrokeColor(color.RGBA{ // 线框颜色
		R: uint8(2),
		G: uint8(119),
		B: uint8(255),
		A: 0xff})
	gc.SetFillColor(color.RGBA{})
	gc.SetLineWidth(5) // 线框宽度
	gc.BeginPath()
	gc.MoveTo(0, 0)
	gc.LineTo(float64(transparentImg.Bounds().Dx()), 0)
	gc.LineTo(float64(transparentImg.Bounds().Dx()), float64(transparentImg.Bounds().Dy()))
	gc.LineTo(0, float64(transparentImg.Bounds().Dy()))
	gc.LineTo(0, 0)
	gc.Close()
	gc.FillStroke()
}

func SaveImage(p string, src image.Image) error {

	outFile, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	b := bufio.NewWriter(outFile)
	err = png.Encode(b, src)
	if err != nil {
		panic(err)
	}
	err = b.Flush()
	if err != nil {
		panic(err)
	}
	// f, err := os.OpenFile(p, os.O_SYNC|os.O_RDWR|os.O_CREATE, 0666)
	// if err != nil {
	// 	return err
	// }
	// // f.Write()//直接写入
	// // defer f.Close()
	// // return nil
	// // defer f.Close()
	// ext := filepath.Ext(p)
	// log.Println(ext)
	// err = png.Encode(f, src)
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
