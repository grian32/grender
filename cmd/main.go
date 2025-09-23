package main

import (
	"grender"
	"log"
	"runtime"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	if err := glfw.Init(); err != nil {
		log.Fatalln(err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)

	window, err := glfw.CreateWindow(640, 480, "GRender Test", nil, nil)
	if err != nil {
		log.Fatalln(err)
	}
	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		log.Fatalln(err)
	}

	tex1, err := grender.NewTexture("../testdata/tex1.jxl")
	if err != nil {
		log.Fatalln(err)
	}
	tex2, err := grender.NewTexture("../testdata/tex2.jxl")
	if err != nil {
		log.Fatalln(err)
	}
	tex3, err := grender.NewTexture("../testdata/tex4.jxl")
	if err != nil {
		log.Fatalln(err)
	}

	atlas := grender.NewAtlas(4096)

	atlas.AddTexture(tex1)
	atlas.AddTexture(tex2)
	atlas.AddTexture(tex3)

	atlas.Upload()

	r := grender.NewRenderer(atlas)

	err = r.InitGL()
	if err != nil {
		log.Fatalln(err)
	}

	r.SetScreenSize(640, 480)

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		r.Begin()

		r.DrawTexture(tex1, 0, 0)
		r.DrawTexture(tex2, uint32(tex1.Size.X), 0)
		r.DrawTexture(tex3, uint32(tex1.Size.X+tex2.Size.X), 0)

		r.End()

		window.SwapBuffers()
		glfw.PollEvents()
	}
}
