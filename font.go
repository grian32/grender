package grender

import (
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type Font struct {
	Atlas *FontAtlas
	Face  font.Face
}

func NewFont(path string, size float64) (*Font, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	fnt, err := opentype.Parse(bytes)
	if err != nil {
		return nil, err
	}

	face, err := opentype.NewFace(fnt, &opentype.FaceOptions{
		Size:    size,
		DPI:     72,
		Hinting: font.HintingNone,
	})

	return &Font{
		Atlas: NewFontAtlas(1024),
		Face:  face,
	}, nil
}
