package grender

import (
	"bytes"
	"log"
	"os"
	"testing"

	"github.com/gen2brain/jpegxl"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

func TestFontAtlas(t *testing.T) {
	fa := NewFontAtlas(1024)

	fBytes, err := os.ReadFile("./testdata/arial.ttf")
	if err != nil {
		t.Errorf("reading ttf file: %v", err)
	}
	fnt, err := opentype.Parse(fBytes)
	if err != nil {
		t.Errorf("parsing ttf file: %v", err)
	}

	face, err := opentype.NewFace(fnt, &opentype.FaceOptions{
		Size:    24,
		DPI:     72,
		Hinting: font.HintingNone,
	})

	err = fa.AddAsciiGlyphs(face)
	if err != nil {
		log.Fatalln(err)
	}

	//f, _ := os.Create("./testdata/font.out.jxl")
	//_ = jpegxl.Encode(f, fa.Atlas.Img)

	enc := &bytes.Buffer{}
	err = jpegxl.Encode(enc, fa.Atlas.Img)
	if err != nil {
		t.Errorf("enconding jxl to buf: %v", err)
	}

	expectedBytes, err := os.ReadFile("./testdata/font.out.jxl")
	if err != nil {
		t.Errorf("reading expected jxl: %v", err)
	}

	if !bytes.Equal(enc.Bytes(), expectedBytes) {
		t.Errorf("got %v, want %v", enc.Bytes(), expectedBytes)
	}
}
