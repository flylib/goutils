package texture

import (
	"image"
	"image/draw"
	"image/png"
	"os"
)

func MixImages(bgImg, frontImg string, mixImgName string, offsetX, offsetY int) error {
	//背景图
	bgImageFile, err := os.Open(bgImg)
	if err != nil {
		return err
	}
	defer bgImageFile.Close()

	bgImage, err := png.Decode(bgImageFile)
	if err != nil {
		return err
	}
	bgImageBound := bgImage.Bounds()

	//前景图
	frontImageFile, err := os.Open(frontImg)
	if err != nil {
		return err
	}
	defer frontImageFile.Close()

	frontImage, err := png.Decode(frontImageFile)
	if err != nil {
		return err
	}
	preImageBound := frontImage.Bounds()

	offset := image.Pt((bgImageBound.Max.X-preImageBound.Max.X)/2+offsetX, (bgImageBound.Max.Y-preImageBound.Max.Y)/2+offsetY)
	bgImageBounds := bgImage.Bounds()
	RGBA := image.NewRGBA(bgImageBounds)
	//绘制入RGBA
	draw.Draw(RGBA, bgImageBounds, bgImage, image.Point{}, draw.Src)
	draw.Draw(RGBA, frontImage.Bounds().Add(offset), frontImage, image.Point{}, draw.Over)
	//混合后的图片文件
	mixImageFile, err := os.Create(mixImgName)
	if err != nil {
		return err
	}
	defer mixImageFile.Close()
	//渲染文件
	err = png.Encode(mixImageFile, RGBA)
	if err != nil {
		return err
	}
	return nil
}
