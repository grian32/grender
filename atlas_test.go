package grender

import (
	"bytes"
	"os"
	"testing"

	"github.com/gen2brain/jpegxl"
)

func TestAtlas(t *testing.T) {
	tex1, err := NewTexture("./testdata/tex1.jxl")
	if err != nil {
		t.Errorf("read tex1.jxl: %v", err)
	}
	tex2, err := NewTexture("./testdata/tex2.jxl")
	if err != nil {
		t.Errorf("read tex2.jxl: %v", err)
	}
	tex3, err := NewTexture("./testdata/tex3.jxl")
	if err != nil {
		t.Errorf("read tex3.jxl: %v", err)
	}
	tex4, err := NewTexture("./testdata/tex4.jxl")
	if err != nil {
		t.Errorf("read tex4.jxl: %v", err)
	}

	atlas := NewAtlas(8092)

	atlas.AddTexture(tex1)
	atlas.AddTexture(tex2)
	atlas.AddTexture(tex3)
	atlas.AddTexture(tex4)

	enc := &bytes.Buffer{}
	err = jpegxl.Encode(enc, atlas.Img, jpegxl.Options{
		Quality: 100,
		Effort:  2,
	})
	if err != nil {
		t.Errorf("jxl encode: %v", err)
	}

	expectedBytes, err := os.ReadFile("./testdata/out.jxl")
	if err != nil {
		t.Errorf("read out.jxl: %v", err)
	}

	if !bytes.Equal(expectedBytes, enc.Bytes()) {
		t.Errorf("EncodeAtlas=%v, want match for %v", enc.Bytes(), expectedBytes)
	}
}
