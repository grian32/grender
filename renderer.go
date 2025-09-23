package grender

import (
	"grender/util"
	"log"
	"unsafe"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
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
layout (location = 5) in vec2 size;

uniform mat4 projection;

out vec2 fragUV;

void main() {
	fragUV = atlasOffset + UV * atlasScale;
	
	vec2 scaledPos = pos.xy * size + worldPos;

	gl_Position = projection * vec4(scaledPos, pos.z, 1.0);
}
`

const fragShader = `#version 330 core
out vec4 FragColor;

in vec2 fragUV;

uniform sampler2D atlas;

void main() {
	FragColor = texture(atlas, fragUV);
}
`

func NewRenderer(atlas *Atlas) *Renderer {
	return &Renderer{
		Atlas:      atlas,
		atlasRects: make([]float32, 0),
	}
}

// InitGL meant to be called once after initializing the Renderer via NewRenderer
func (r *Renderer) InitGL() error {
	vertices := []float32{
		// pos, tex coord
		1.0, 1.0, 0.0, 1.0, 1.0,
		1.0, 0.0, 0.0, 1.0, 0.0,
		0.0, 0.0, 0.0, 0.0, 0.0,
		0.0, 1.0, 0.0, 0.0, 1.0,
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
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(r.atlasRects), nil, gl.DYNAMIC_DRAW)

	gl.VertexAttribPointer(2, 2, gl.FLOAT, false, 8*4, unsafe.Pointer(uintptr(0)))
	gl.EnableVertexAttribArray(2)
	gl.VertexAttribDivisor(2, 1)

	gl.VertexAttribPointer(3, 2, gl.FLOAT, false, 8*4, unsafe.Pointer(uintptr(2*4)))
	gl.EnableVertexAttribArray(3)
	gl.VertexAttribDivisor(3, 1)

	gl.VertexAttribPointer(4, 2, gl.FLOAT, false, 8*4, unsafe.Pointer(uintptr(4*4)))
	gl.EnableVertexAttribArray(4)
	gl.VertexAttribDivisor(4, 1)

	gl.VertexAttribPointer(5, 2, gl.FLOAT, false, 8*4, unsafe.Pointer(uintptr(6*4)))
	gl.EnableVertexAttribArray(5)
	gl.VertexAttribDivisor(5, 1)

	shaderId, err := util.CompileGLShader(vertShader, fragShader)
	if err != nil {
		return err
	}

	r.shader = shaderId

	return nil
}

// SetScreenSize must be called after InitGL
func (r *Renderer) SetScreenSize(w, h uint32) {
	projection := mgl32.Ortho(0, float32(w), float32(h), 0, -1, 1)

	gl.UseProgram(r.shader)
	projLoc := gl.GetUniformLocation(r.shader, gl.Str("projection\x00"))
	gl.UniformMatrix4fv(projLoc, 1, false, &projection[0])
}

func (r *Renderer) DrawTexture(t *Texture, x, y uint32) {
	atlasOffsetX := float32(t.AtlasPos.X) / float32(r.Atlas.Size)
	atlasOffsetY := float32(t.AtlasPos.Y) / float32(r.Atlas.Size)
	atlasScaleX := float32(t.Size.X) / float32(r.Atlas.Size)
	atlasScaleY := float32(t.Size.Y) / float32(r.Atlas.Size)

	log.Printf("Atlas Offset=(%.8f, %.8f) scale=(%.8f, %.8f)", atlasOffsetX, atlasOffsetY, atlasScaleX, atlasScaleY)

	r.atlasRects = append(r.atlasRects,
		atlasOffsetX, atlasOffsetY,
		atlasScaleX, atlasScaleY,
		float32(x), float32(y),
		float32(t.Size.X), float32(t.Size.Y),
	)
}

func (r *Renderer) Begin() {
	gl.UseProgram(r.shader)
	gl.BindVertexArray(r.vao)

	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, r.Atlas.texId)

	loc := gl.GetUniformLocation(r.shader, gl.Str("atlas\x00"))
	gl.Uniform1i(loc, 0)

	r.atlasRects = r.atlasRects[:0]
}

func (r *Renderer) End() {
	if len(r.atlasRects) == 0 {
		log.Printf("returning as no atlas rects")
		return
	}

	gl.BindBuffer(gl.ARRAY_BUFFER, r.atlasVBO)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(r.atlasRects), gl.Ptr(r.atlasRects), gl.DYNAMIC_DRAW)

	instanceCount := int32(len(r.atlasRects) / 8)
	gl.DrawElementsInstanced(gl.TRIANGLES, 6, gl.UNSIGNED_INT, nil, instanceCount)

	gl.BindVertexArray(0)
	gl.UseProgram(0)
}
