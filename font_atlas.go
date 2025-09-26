package grender

import (
	"fmt"
	"image"
	"image/color"
	"math"

	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

type FontAtlas struct {
	Positions map[rune]RectP
	Atlas     *Atlas
	Size      uint32
}

func NewFontAtlas(size uint32) *FontAtlas {
	return &FontAtlas{
		Positions: make(map[rune]RectP),
		Atlas:     NewAtlas(AtlasOption(size)),
		Size:      size,
	}
}

func (fa *FontAtlas) AddAsciiGlyphs(face font.Face) error {
	for i := 32; i < 127; i++ {
		ch := rune(i)

		bounds, _, ok := face.GlyphBounds(ch)
		if !ok {
			return fmt.Errorf("could not find glyph bounds for %c", ch)
		}

		score1 := math.MaxInt32
		score2 := math.MaxInt32

		// we pad by 1px due to GlyphBounds presumably not accounting for AA pixels in measurement

		width := ((bounds.Max.X - bounds.Min.X) >> 6) + 4
		height := ((bounds.Max.Y - bounds.Min.Y) >> 6) + 4

		newNode := fa.Atlas.findPositionForNewNode(int32(width), int32(height), int32(score1), int32(score2))
		fa.Atlas.placeRect(newNode)

		dotX := fixed.Int26_6(int(newNode.X+2)-bounds.Min.X.Floor()) << 6
		dotY := fixed.Int26_6(int(newNode.Y+2)-bounds.Min.Y.Floor()) << 6

		d := &font.Drawer{
			Dst:  fa.Atlas.Img,
			Src:  image.NewUniform(color.White),
			Face: face,
			Dot: fixed.Point26_6{
				X: dotX,
				Y: dotY,
			},
		}

		d.DrawString(string(ch))

		fa.Positions[ch] = RectP{
			X: newNode.X,
			Y: newNode.Y,
			W: newNode.W,
			H: newNode.H,
		}
	}

	return nil
}
