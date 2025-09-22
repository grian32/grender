package grender

import (
	"fmt"
	"image"
	"image/draw"
	"os"

	"github.com/gen2brain/jpegxl"
)

type Texture struct {
	Image    *image.NRGBA
	Size     Vec2I
	AtlasPos RectP
}

func NewTexture(jxlPath string) (*Texture, error) {
	tex := &Texture{
		Size:     Vec2I{},
		AtlasPos: RectP{},
	}

	f, err := os.Open(jxlPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	jxl, err := jpegxl.Decode(f)
	if err != nil {
		return nil, err
	}

	switch img := jxl.(type) {
	case *image.NRGBA:
		tex.Image = img
	case *image.NRGBA64:
		nrgba := image.NewNRGBA(img.Bounds())
		draw.Draw(nrgba, nrgba.Bounds(), img, img.Bounds().Min, draw.Src)
		tex.Image = nrgba
	default:
		return nil, fmt.Errorf("unexpected img %T", img)
	}

	tex.Size.X = int32(jxl.Bounds().Dx())
	tex.Size.Y = int32(jxl.Bounds().Dy())

	return tex, nil
}
