package main

import (
	"grender"
	"log"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	err := grender.CreateWindow(640, 480, "GRender Test")
	if err != nil {
		log.Fatalln(err)
	}
	defer grender.CloseWindow()

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

	r, err := grender.NewRenderer(atlas)
	if err != nil {
		log.Fatalln(err)
	}

	defer r.Cleanup()

	for grender.WindowShouldNotClose() {
		// loop logic here
		r.Begin()

		r.DrawTexture(tex1, 0, 0)
		r.DrawTexture(tex2, uint32(tex1.Size.X), 0)
		r.DrawTexture(tex3, uint32(tex1.Size.X+tex2.Size.X), 0)

		r.End()
	}
}
