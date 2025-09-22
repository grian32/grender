package grender

import (
	"unsafe"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type Renderer struct {
	Atlas      *Atlas
	vao        uint32
	vbo        uint32
	ebo        uint32
	atlasVBO   uint32
	shader     uint32
	atlasRects []float32
}

const vertShader = `#version 330 core
layout (location = 0) in vec3 pos;
layout (location = 1) in vec2 UV;
layout (location = 2) in vec2 atlasOffset;
layout (location = 3) in vec2 atlasScale;
layout (location = 4) in vec2 worldPos;

out vec2 fragUV;

void main() {
	fragUV = atlasOffset + UV * atlasScale;
	
	gl_Position = vec4(pos.xy + worldPos, pos.z, 1.0);
}
`

const fragShader = `#version 330 core

in vec2 fragUV;
out vec4 outColor;

uniform sampler2D atlas;

void main() {
	outColor = texture(atlas, fragUV);
}
`

func NewRenderer(atlas *Atlas) *Renderer {
	return &Renderer{
		Atlas:      atlas,
		atlasRects: make([]float32, 0),
	}
}

func (r *Renderer) InitGL() {
	vertices := []float32{
		// pos, tex coord
		0.5, 0.5, 0.0, 1.0, 1.0,
		0.5, -0.5, 0.0, 1.0, 0.0,
		-0.5, -0.5, 0.0, 0.0, 0.0,
		-0.5, 0.5, 0.0, 0.0, 1.0,
	}

	indices := []uint32{
		0, 1, 3,
		1, 2, 3,
	}

	gl.GenVertexArrays(1, &r.vao)
	gl.BindVertexArray(r.vao)

	gl.GenBuffers(1, &r.vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, r.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(vertices), gl.Ptr(vertices), gl.STATIC_DRAW)

	gl.GenBuffers(1, &r.ebo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, r.ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, 4*len(indices), gl.Ptr(indices), gl.STATIC_DRAW)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 5*4, unsafe.Pointer(uintptr(0)))
	gl.EnableVertexAttribArray(0)

	gl.VertexAttribPointer(1, 2, gl.FLOAT, false, 5*4, unsafe.Pointer(uintptr(3*4)))
	gl.EnableVertexAttribArray(1)

	gl.GenBuffers(1, &r.atlasVBO)
	gl.BindBuffer(gl.ARRAY_BUFFER, r.atlasVBO)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(r.atlasRects), gl.Ptr(r.atlasRects), gl.DYNAMIC_DRAW)

	gl.VertexAttribPointer(2, 2, gl.FLOAT, false, 6*4, unsafe.Pointer(uintptr(0)))
	gl.EnableVertexAttribArray(2)
	gl.VertexAttribDivisor(2, 1)

	gl.VertexAttribPointer(3, 2, gl.FLOAT, false, 6*4, unsafe.Pointer(uintptr(2*4)))
	gl.EnableVertexAttribArray(3)
	gl.VertexAttribDivisor(3, 1)

	gl.VertexAttribPointer(4, 2, gl.FLOAT, false, 6*4, unsafe.Pointer(uintptr(4*4)))
	gl.EnableVertexAttribArray(4)
	gl.VertexAttribDivisor(4, 1)
}
