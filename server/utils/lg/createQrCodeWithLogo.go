package lg

import (
	"bytes"
	"github.com/nfnt/resize"
	"github.com/skip2/go-qrcode"
	"image"
	"image/draw"
	"image/png"
	"os"
)

func CreateQrCodeBs64WithLogo(content, logoPath, outPath string, size int) (err error) {
	code, err := qrcode.New(content, qrcode.High)
	if err != nil {
		return
	}
	//code.DisableBorder = true
	// 设置文件大小并创建画板
	qrcodeImg := code.Image(size)
	outImg := image.NewRGBA(qrcodeImg.Bounds())

	// 读取logo文件
	logoFile, err := os.Open(logoPath)
	if err != nil {
		return
	}
	logoImg, _, err := image.Decode(logoFile)
	logoImg = resize.Resize(uint(size/5), 0, logoImg, resize.Lanczos3)

	// 添加边框
	// 图片到边框距离
	pic2FramePadding := logoImg.Bounds().Dx() / 10

	// 新建一个边框图层
	transparentImg := image.NewRGBA(image.Rect(0, 0, logoImg.Bounds().Dx()+pic2FramePadding, logoImg.Bounds().Dy()+pic2FramePadding))
	// 图层颜色设为白色
	draw.Draw(transparentImg, transparentImg.Bounds(), image.White, image.Point{}, draw.Over)
	// 将缩略图放到透明图层上
	draw.Draw(transparentImg,
		image.Rect(pic2FramePadding/2, pic2FramePadding/2, logoImg.Bounds().Dx(), logoImg.Bounds().Dy()),
		logoImg,
		image.Point{},
		draw.Over)

	// logo和二维码拼接
	draw.Draw(outImg, qrcodeImg.Bounds(), qrcodeImg, image.Pt(0, 0), draw.Over)
	offset := image.Pt((outImg.Bounds().Max.X-transparentImg.Bounds().Max.X)/2, (outImg.Bounds().Max.Y-transparentImg.Bounds().Max.Y)/2)
	draw.Draw(outImg, transparentImg.Bounds().Add(offset), transparentImg, image.Pt(0, 0), draw.Over)

	buf := new(bytes.Buffer)
	_ = png.Encode(buf, outImg)

	// 写入文件
	f, _ := os.Create(outPath)
	_ = png.Encode(f, outImg)

	return nil
}
