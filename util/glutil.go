package util

import (
	"fmt"
	"log"
	"strings"

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

func CompileGLShader(v, f string) (uint32, error) {
	var success int32
	infoLog := gl.Str(strings.Repeat("\x00", 512))

	vGlSrc, vFreeFnc := gl.Strs(v + "\x00")
	defer vFreeFnc()

	fGlSrc, fFreeFnc := gl.Strs(f + "\x00")
	defer fFreeFnc()

	vertex := gl.CreateShader(gl.VERTEX_SHADER)

	gl.ShaderSource(vertex, 1, vGlSrc, nil)
	gl.CompileShader(vertex)

	gl.GetShaderiv(vertex, gl.COMPILE_STATUS, &success)

	if success != gl.TRUE {
		gl.GetShaderInfoLog(vertex, 512, nil, infoLog)
		return 0, fmt.Errorf("ERROR::SHADER::VERTEX::COMPILE:\n%v", infoLog)
	}

	fragment := gl.CreateShader(gl.FRAGMENT_SHADER)

	gl.ShaderSource(fragment, 1, fGlSrc, nil)
	gl.CompileShader(fragment)

	gl.GetShaderiv(fragment, gl.COMPILE_STATUS, &success)

	if success != gl.TRUE {
		gl.GetShaderInfoLog(fragment, 512, nil, infoLog)
		return 0, fmt.Errorf("ERROR::SHADER::FRAGMENT::COMPILE:\n%v", infoLog)
	}

	id := gl.CreateProgram()
	gl.AttachShader(id, vertex)
	gl.AttachShader(id, fragment)
	gl.LinkProgram(id)

	gl.GetProgramiv(id, gl.LINK_STATUS, &success)

	if success != gl.TRUE {
		gl.GetProgramInfoLog(id, 512, nil, infoLog)
		return 0, fmt.Errorf("ERROR::SHADER::LINK:\n%v", infoLog)
	}

	gl.DeleteShader(vertex)
	gl.DeleteShader(fragment)

	return id, nil
}
