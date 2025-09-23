package grender

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"os"

	"github.com/gen2brain/jpegxl"
)

type Texture struct {
	Image    *image.NRGBA
	Size     Vec2I
	AtlasPos RectP
}

// NewTexture accepts a path to a .jxl file and creates a texture
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

// NewTextureB accepts a []byte of a jxl file and creates a texture
func NewTextureB(b []byte) (*Texture, error) {
	tex := &Texture{
		Size:     Vec2I{},
		AtlasPos: RectP{},
	}

	r := bytes.NewReader(b)

	jxl, err := jpegxl.Decode(r)
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

func NewTextureColor(r, g, b uint8) *Texture {
	img := image.NewNRGBA(image.Rect(0, 0, 1, 1))
	img.Set(0, 0, color.NRGBA{r, g, b, 255})

	return &Texture{
		Image:    img,
		Size:     Vec2I{X: 1, Y: 1},
		AtlasPos: RectP{},
	}
}
