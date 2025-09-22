package util

import (
	"log"

	"github.com/go-gl/gl/v3.3-core/gl"
)

func GetGlIntP(pname uint32) int32 {
	var value int32
	gl.GetIntegerv(pname, &value)
	if err := gl.GetError(); err != gl.NO_ERROR {
		log.Fatalf("OpenGL getIntegerv(0x%x) error: 0x%x\n", pname, err)
	}

	return value
}
